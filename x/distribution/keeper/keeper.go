package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/tendermint/tendermint/libs/log"
)

// keeper of the staking store
type Keeper struct {
	storeKey            sdk.StoreKey
	cdc                 *codec.Codec
	paramSpace          params.Subspace
	bankKeeper          types.BankKeeper
	stakingKeeper       types.StakingKeeper
	feeCollectionKeeper types.FeeCollectionKeeper
	supplyKeeper        SupplyKeeper

	// codespace
	codespace sdk.CodespaceType
}

// create a new keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace, ck types.BankKeeper,
	sk types.StakingKeeper, fck types.FeeCollectionKeeper, supplyKeeper SupplyKeeper, codespace sdk.CodespaceType) Keeper {
	keeper := Keeper{
		storeKey:            key,
		cdc:                 cdc,
		paramSpace:          paramSpace.WithKeyTable(ParamKeyTable()),
		bankKeeper:          ck,
		stakingKeeper:       sk,
		feeCollectionKeeper: fck,
		supplyKeeper:        supplyKeeper,
		codespace:           codespace,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger { return ctx.Logger().With("module", "x/distr") }

// set withdraw address
func (k Keeper) SetWithdrawAddr(ctx sdk.Context, delegatorAddr sdk.AccAddress, withdrawAddr sdk.AccAddress) sdk.Error {
	if !k.GetWithdrawAddrEnabled(ctx) {
		return types.ErrSetWithdrawAddrDisabled(k.codespace)
	}

	k.SetDelegatorWithdrawAddr(ctx, delegatorAddr, withdrawAddr)

	return nil
}

// withdraw rewards from a delegation
func (k Keeper) WithdrawDelegationRewards(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) sdk.Error {
	val := k.stakingKeeper.Validator(ctx, valAddr)
	if val == nil {
		return types.ErrNoValidatorDistInfo(k.codespace)
	}

	del := k.stakingKeeper.Delegation(ctx, delAddr, valAddr)
	if del == nil {
		return types.ErrNoDelegationDistInfo(k.codespace)
	}

	// withdraw rewards
	if err := k.withdrawDelegationRewards(ctx, val, del); err != nil {
		return err
	}

	// reinitialize the delegation
	k.initializeDelegation(ctx, valAddr, delAddr)
	return nil
}

// withdraw validator commission
func (k Keeper) WithdrawValidatorCommission(ctx sdk.Context, valAddr sdk.ValAddress) sdk.Error {

	// fetch validator accumulated commission
	commission := k.GetValidatorAccumulatedCommission(ctx, valAddr)
	if commission.IsZero() {
		return types.ErrNoValidatorCommission(k.codespace)
	}

	coins, remainder := commission.TruncateDecimal()
	k.SetValidatorAccumulatedCommission(ctx, valAddr, remainder) // leave remainder to withdraw later

	// update outstanding
	outstanding := k.GetValidatorOutstandingRewards(ctx, valAddr)
	k.SetValidatorOutstandingRewards(ctx, valAddr, outstanding.Sub(sdk.NewDecCoins(coins)))

	if !coins.IsZero() {
		accAddr := sdk.AccAddress(valAddr)
		withdrawAddr := k.GetDelegatorWithdrawAddr(ctx, accAddr)

		if _, err := k.bankKeeper.AddCoins(ctx, withdrawAddr, coins); err != nil {
			return err
		}

		k.supplyKeeper.InflateSupply(ctx, supply.TypeCirculating, coins)
	}

	return nil
}

// GetTotalRewards returns the total amount of fee distribution rewards held in the store
func (k Keeper) GetTotalRewards(ctx sdk.Context) (totalRewards sdk.DecCoins) {
	k.IterateValidatorOutstandingRewards(ctx,
		func(valAddr sdk.ValAddress, rewards types.ValidatorOutstandingRewards) (stop bool) {
			totalRewards = totalRewards.Add(rewards)
			return false
		},
	)
	return totalRewards
}
