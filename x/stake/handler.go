package stake

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

//nolint
const (
	GasDeclareCandidacy int64 = 20
	GasEditCandidacy    int64 = 20
	GasDelegate         int64 = 20
	GasUnbond           int64 = 20
)

//XXX fix initstater
// separated for testing
//func InitState(ctx sdk.Context, k Keeper, key, value string) sdk.Error {

//params := k.GetParams(ctx)
//switch key {
//case "allowed_bond_denom":
//params.BondDenom = value
//case "max_vals", "gas_bond", "gas_unbond":

//i, err := strconv.Atoi(value)
//if err != nil {
//return sdk.ErrUnknownRequest(fmt.Sprintf("input must be integer, Error: %v", err.Error()))
//}

//switch key {
//case "max_vals":
//if i < 0 {
//return sdk.ErrUnknownRequest("cannot designate negative max validators")
//}
//params.MaxValidators = uint16(i)
//case "gas_bond":
//GasDelegate = int64(i)
//case "gas_unbound":
//GasUnbond = int64(i)
//}
//default:
//return sdk.ErrUnknownRequest(key)
//}

//k.setParams(params)
//return nil
//}

//_______________________________________________________________________

func NewHandler(k Keeper, ck bank.CoinKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case MsgDeclareCandidacy:
			return handleMsgDeclareCandidacy(ctx, msg, k)
		case MsgEditCandidacy:
			return handleMsgEditCandidacy(ctx, msg, k)
		case MsgDelegate:
			return handleMsgDelegate(ctx, msg, k)
		case MsgUnbond:
			return handleMsgUnbond(ctx, msg, k)
		default:
			return sdk.ErrTxDecode("invalid message parse in staking module").Result()
		}
	}
}

//_____________________________________________________________________

// XXX should be send in the msg (init in CLI)
//func getSender() sdk.Address {
//signers := msg.GetSigners()
//if len(signers) != 1 {
//return sdk.ErrUnauthorized("there can only be one signer for staking transaction").Result()
//}
//sender := signers[0]
//}

//_____________________________________________________________________

// These functions assume everything has been authenticated,
// now we just perform action and save

func handleMsgDeclareCandidacy(ctx sdk.Context, msg MsgDeclareCandidacy, k Keeper) sdk.Result {

	// check to see if the pubkey or sender has been registered before
	_, found := k.GetCandidate(ctx, msg.CandidateAddr)
	if found {
		return ErrCandidateExistsAddr().Result()
	}
	if msg.Bond.Denom != k.GetParams(ctx).BondDenom {
		return ErrBadBondingDenom().Result()
	}
	if ctx.IsCheckTx() {
		return sdk.Result{
			GasUsed: GasDeclareCandidacy,
		}
	}

	candidate := NewCandidate(msg.CandidateAddr, msg.PubKey, msg.Description)
	k.setCandidate(ctx, candidate)

	// move coins from the msg.Address account to a (self-bond) delegator account
	// the candidate account and global shares are updated within here
	return delegateWithCandidate(ctx, k, msg.CandidateAddr, msg.Bond, candidate).Result()
}

func handleMsgEditCandidacy(ctx sdk.Context, msg MsgEditCandidacy, k Keeper) sdk.Result {

	// candidate must already be registered
	candidate, found := k.GetCandidate(ctx, msg.CandidateAddr)
	if !found {
		return ErrBadCandidateAddr().Result()
	}
	if ctx.IsCheckTx() {
		return sdk.Result{
			GasUsed: GasEditCandidacy,
		}
	}
	if candidate.Status == Unbonded { //candidate has been withdrawn
		return ErrBondNotNominated().Result()
	}

	// XXX move to types
	//check and edit any of the editable terms
	if msg.Description.Moniker != "" {
		candidate.Description.Moniker = msg.Description.Moniker
	}
	if msg.Description.Identity != "" {
		candidate.Description.Identity = msg.Description.Identity
	}
	if msg.Description.Website != "" {
		candidate.Description.Website = msg.Description.Website
	}
	if msg.Description.Details != "" {
		candidate.Description.Details = msg.Description.Details
	}

	k.setCandidate(ctx, candidate)
	return sdk.Result{}
}

func handleMsgDelegate(ctx sdk.Context, msg MsgDelegate, k Keeper) sdk.Result {

	candidate, found := k.GetCandidate(ctx, msg.CandidateAddr)
	if !found {
		return ErrBadCandidateAddr().Result()
	}
	if msg.Bond.Denom != k.GetParams(ctx).BondDenom {
		return ErrBadBondingDenom().Result()
	}
	if ctx.IsCheckTx() {
		return sdk.Result{
			GasUsed: GasDelegate,
		}
	}
	return delegateWithCandidate(ctx, k, msg.DelegatorAddr, msg.Bond, candidate).Result()
}

func delegateWithCandidate(ctx sdk.Context, k Keeper, delegatorAddr sdk.Address,
	bondAmt sdk.Coin, candidate Candidate) sdk.Error {

	if candidate.Status == Revoked { //candidate has been withdrawn
		return ErrBondNotNominated()
	}

	// Get or create the delegator bond
	existingBond, found := k.getDelegatorBond(ctx, delegatorAddr, candidate.Address)
	if !found {
		existingBond = DelegatorBond{
			DelegatorAddr: delegatorAddr,
			CandidateAddr: candidate.Address,
			Shares:        sdk.ZeroRat,
		}
	}

	// Account new shares, save
	err := BondCoins(ctx, k, existingBond, candidate, bondAmt)
	if err != nil {
		return err
	}
	k.setDelegatorBond(ctx, existingBond)
	k.setCandidate(ctx, candidate)
	return nil
}

// Perform all the actions required to bond tokens to a delegator bond from their account
func BondCoins(ctx sdk.Context, k Keeper, bond DelegatorBond, candidate Candidate, amount sdk.Coin) sdk.Error {

	_, err := k.coinKeeper.SubtractCoins(ctx, bond.DelegatorAddr, sdk.Coins{amount})
	if err != nil {
		return err
	}
	newShares := k.candidateAddTokens(ctx, candidate, amount.Amount)
	bond.Shares = bond.Shares.Add(newShares)
	k.setDelegatorBond(ctx, bond)
	return nil
}

func handleMsgUnbond(ctx sdk.Context, msg MsgUnbond, k Keeper) sdk.Result {

	// check if bond has any shares in it unbond
	bond, found := k.getDelegatorBond(ctx, msg.DelegatorAddr, msg.CandidateAddr)
	if !found {
		return ErrNoDelegatorForAddress().Result()
	}
	if !bond.Shares.GT(sdk.ZeroRat) { // bond shares < msg shares
		return ErrInsufficientFunds().Result()
	}

	// test getting rational number from decimal provided
	shares, err := sdk.NewRatFromDecimal(msg.Shares)
	if err != nil {
		return err.Result()
	}

	// test that there are enough shares to unbond
	if msg.Shares == "MAX" {
		if !bond.Shares.GT(sdk.ZeroRat) {
			return ErrNotEnoughBondShares(msg.Shares).Result()
		}
	} else {
		if !bond.Shares.GT(shares) {
			return ErrNotEnoughBondShares(msg.Shares).Result()
		}
	}

	// get candidate
	candidate, found := k.GetCandidate(ctx, msg.CandidateAddr)
	if !found {
		return ErrNoCandidateForAddress().Result()
	}

	if ctx.IsCheckTx() {
		return sdk.Result{
			GasUsed: GasUnbond,
		}
	}

	// retrieve the amount of bonds to remove (TODO remove redundancy already serialized)
	var shares sdk.Rat
	var err sdk.Error
	if msg.Shares == "MAX" {
		shares = bond.Shares
	} else {
		shares, err = sdk.NewRatFromDecimal(msg.Shares)
		if err != nil {
			return err.Result()
		}
	}

	// subtract bond tokens from delegator bond
	bond.Shares = bond.Shares.Sub(shares)

	// remove the bond
	revokeCandidacy := false
	if bond.Shares.IsZero() {

		// if the bond is the owner of the candidate then
		// trigger a revoke candidacy
		if bytes.Equal(bond.DelegatorAddr, candidate.Address) &&
			candidate.Status != Revoked {
			revokeCandidacy = true
		}

		k.removeDelegatorBond(ctx, bond)
	} else {
		k.setDelegatorBond(ctx, bond)
	}

	// Add the coins
	returnAmount := k.candidateRemoveShares(ctx, candidate, shares)
	returnCoins := sdk.Coins{{k.GetParams(ctx).BondDenom, returnAmount}}
	k.coinKeeper.AddCoins(ctx, bond.DelegatorAddr, returnCoins)

	// revoke candidate if necessary
	if revokeCandidacy {

		// change the share types to unbonded if they were not already
		if candidate.Status == Bonded {
			k.bondedToUnbondedPool(ctx, candidate)
		}

		// lastly update the status
		candidate.Status = Revoked
	}

	// deduct shares from the candidate
	if candidate.Liabilities.IsZero() {
		k.removeCandidate(ctx, candidate.Address)
	} else {
		k.setCandidate(ctx, candidate)
	}
	return sdk.Result{}
}

// XXX where this used
// Perform all the actions required to bond tokens to a delegator bond from their account
func UnbondCoins(ctx sdk.Context, k Keeper, bond DelegatorBond, candidate Candidate, shares sdk.Rat) sdk.Error {

	// subtract bond tokens from delegator bond
	if bond.Shares.LT(shares) {
		return sdk.ErrInsufficientFunds("") //XXX variables inside
	}
	bond.Shares = bond.Shares.Sub(shares)

	returnAmount := k.candidateRemoveShares(ctx, candidate, shares)
	returnCoins := sdk.Coins{{k.GetParams(ctx).BondDenom, returnAmount}}

	_, err := k.coinKeeper.AddCoins(ctx, candidate.Address, returnCoins)
	if err != nil {
		return err
	}
	return nil
}
