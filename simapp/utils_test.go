package simapp

import (
	"encoding/binary"
	"fmt"
	"time"

	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/tendermint/crypto/ed25519"
	cmn "github.com/tendermint/tendermint/libs/common"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/staking"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	delPk1    = ed25519.GenPrivKey().PubKey()
	delAddr1  = sdk.AccAddress(delPk1.Address())
	valAddr1  = sdk.ValAddress(delPk1.Address())
	consAddr1 = sdk.ConsAddress(delPk1.Address().Bytes())
)

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	authtypes.RegisterCodec(cdc)
	distr.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	return
}

func TestDecodeAccountStore(t *testing.T) {
	cdc := makeTestCodec()
	acc := authtypes.NewBaseAccountWithAddress(delAddr1)
	bz := cdc.MustMarshalBinaryBare(acc)

	require.Equal(t, fmt.Sprintf("%v\n%v", acc, acc), decodeAccountStore(cdc, cdc, bz, bz))
}

func TestDecodeMintStore(t *testing.T) {
	cdc := makeTestCodec()
	minter := minttypes.NewMinter(sdk.OneDec(), sdk.NewDec(15))
	bz := cdc.MustMarshalBinaryLengthPrefixed(minter)
	require.Equal(t, fmt.Sprintf("%v\n%v", minter, minter), decodeMintStore(cdc, cdc, bz, bz))
}

func TestDecodeDistributionStore(t *testing.T) {
	cdc := makeTestCodec()

	decCoins := sdk.DecCoins{sdk.NewDecCoinFromDec(sdk.DefaultBondDenom, sdk.OneDec())}
	feePool := distr.InitialFeePool()
	feePool.CommunityPool = decCoins
	info := distr.NewDelegatorStartingInfo(2, sdk.OneDec(), 200)
	outstanding := distr.ValidatorOutstandingRewards{decCoins[0]}
	commission := distr.ValidatorAccumulatedCommission{decCoins[0]}
	historicalRewards := distr.NewValidatorHistoricalRewards(decCoins, 100)
	currentRewards := distr.NewValidatorCurrentRewards(decCoins, 5)
	slashEvent := distr.NewValidatorSlashEvent(10, sdk.OneDec())

	kvPairs := cmn.KVPairs{
		cmn.KVPair{Key: distr.FeePoolKey, Value: cdc.MustMarshalBinaryLengthPrefixed(feePool)},
		cmn.KVPair{Key: distr.ProposerKey, Value: consAddr1.Bytes()},
		cmn.KVPair{Key: distr.GetValidatorOutstandingRewardsKey(valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(outstanding)},
		cmn.KVPair{Key: distr.GetDelegatorWithdrawAddrKey(delAddr1), Value: delAddr1.Bytes()},
		cmn.KVPair{Key: distr.GetDelegatorStartingInfoKey(valAddr1, delAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(info)},
		cmn.KVPair{Key: distr.GetValidatorHistoricalRewardsKey(valAddr1, 100), Value: cdc.MustMarshalBinaryLengthPrefixed(historicalRewards)},
		cmn.KVPair{Key: distr.GetValidatorCurrentRewardsKey(valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(currentRewards)},
		cmn.KVPair{Key: distr.GetValidatorAccumulatedCommissionKey(valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(commission)},
		cmn.KVPair{Key: distr.GetValidatorSlashEventKey(valAddr1, 13), Value: cdc.MustMarshalBinaryLengthPrefixed(slashEvent)},
		cmn.KVPair{Key: []byte{0x99}, Value: []byte{0x99}},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"FeePool", fmt.Sprintf("%v\n%v", feePool, feePool)},
		{"Proposer", fmt.Sprintf("%v\n%v", consAddr1, consAddr1)},
		{"ValidatorOutstandingRewards", fmt.Sprintf("%v\n%v", outstanding, outstanding)},
		{"DelegatorWithdrawAddr", fmt.Sprintf("%v\n%v", delAddr1, delAddr1)},
		{"DelegatorStartingInfo", fmt.Sprintf("%v\n%v", info, info)},
		{"ValidatorHistoricalRewards", fmt.Sprintf("%v\n%v", historicalRewards, historicalRewards)},
		{"ValidatorCurrentRewards", fmt.Sprintf("%v\n%v", currentRewards, currentRewards)},
		{"ValidatorAccumulatedCommission", fmt.Sprintf("%v\n%v", commission, commission)},
		{"ValidatorSlashEvent", fmt.Sprintf("%v\n%v", slashEvent, slashEvent)},
		{"other", ""},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { decodeDistributionStore(cdc, cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, decodeDistributionStore(cdc, cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}

func TestDecodeStakingStore(t *testing.T) {
	cdc := makeTestCodec()

	bondTime := time.Now().UTC()

	pool := staking.InitialPool()
	val := staking.NewValidator(valAddr1, delPk1, staking.NewDescription("test", "test", "test", "test"))
	del := staking.NewDelegation(delAddr1, valAddr1, sdk.OneDec())
	ubd := staking.NewUnbondingDelegation(delAddr1, valAddr1, 15, bondTime, sdk.OneInt())
	red := staking.NewRedelegation(delAddr1, valAddr1, valAddr1, 12, bondTime, sdk.OneInt(), sdk.OneDec())

	kvPairs := cmn.KVPairs{
		cmn.KVPair{Key: staking.PoolKey, Value: cdc.MustMarshalBinaryLengthPrefixed(pool)},
		cmn.KVPair{Key: staking.LastTotalPowerKey, Value: cdc.MustMarshalBinaryLengthPrefixed(sdk.OneInt())},
		cmn.KVPair{Key: staking.GetValidatorKey(valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(val)},
		cmn.KVPair{Key: staking.LastValidatorPowerKey, Value: valAddr1.Bytes()},
		cmn.KVPair{Key: staking.GetDelegationKey(delAddr1, valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(del)},
		cmn.KVPair{Key: staking.GetUBDKey(delAddr1, valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(ubd)},
		cmn.KVPair{Key: staking.GetREDKey(delAddr1, valAddr1, valAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(red)},
		cmn.KVPair{Key: []byte{0x99}, Value: []byte{0x99}},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Pool", fmt.Sprintf("%v\n%v", pool, pool)},
		{"LastTotalPower", fmt.Sprintf("%v\n%v", sdk.OneInt(), sdk.OneInt())},
		{"Validator", fmt.Sprintf("%v\n%v", val, val)},
		{"LastValidatorPower/ValidatorsByConsAddr/ValidatorsByPowerIndex", fmt.Sprintf("%v\n%v", valAddr1, valAddr1)},
		{"Delegation", fmt.Sprintf("%v\n%v", del, del)},
		{"UnbondingDelegation", fmt.Sprintf("%v\n%v", ubd, ubd)},
		{"Redelegation", fmt.Sprintf("%v\n%v", red, red)},
		{"other", ""},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { decodeStakingStore(cdc, cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, decodeStakingStore(cdc, cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}

func TestDecodeGovStore(t *testing.T) {
	cdc := makeTestCodec()

	endTime := time.Now().UTC()

	content := gov.ContentFromProposalType("test", "test", gov.ProposalTypeText)
	proposal := gov.NewProposal(content, 1, endTime, endTime.Add(24*time.Hour))
	proposalIDBz := make([]byte, 8)
	binary.LittleEndian.PutUint64(proposalIDBz, 1)
	deposit := gov.NewDeposit(1, delAddr1, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.OneInt())))
	vote := gov.NewVote(1, delAddr1, gov.OptionYes)

	kvPairs := cmn.KVPairs{
		cmn.KVPair{Key: gov.ProposalKey(1), Value: cdc.MustMarshalBinaryLengthPrefixed(proposal)},
		cmn.KVPair{Key: gov.InactiveProposalQueueKey(1, endTime), Value: proposalIDBz},
		cmn.KVPair{Key: gov.DepositKey(1, delAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(deposit)},
		cmn.KVPair{Key: gov.VoteKey(1, delAddr1), Value: cdc.MustMarshalBinaryLengthPrefixed(vote)},
		cmn.KVPair{Key: []byte{0x99}, Value: []byte{0x99}},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"proposals", fmt.Sprintf("%v\n%v", proposal, proposal)},
		{"proposal IDs", "proposalIDA: 1\nProposalIDB: 1"},
		{"deposits", fmt.Sprintf("%v\n%v", deposit, deposit)},
		{"votes", fmt.Sprintf("%v\n%v", vote, vote)},
		{"other", ""},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { decodeGovStore(cdc, cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, decodeGovStore(cdc, cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
