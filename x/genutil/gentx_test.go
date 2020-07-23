package genutil_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	priv1 = secp256k1.GenPrivKey()
	priv2 = secp256k1.GenPrivKey()
	pk1   = priv1.PubKey()
	pk2   = priv2.PubKey()
	addr1 = sdk.AccAddress(pk1.Address())
	addr2 = sdk.AccAddress(pk2.Address())
	desc  = stakingtypes.NewDescription("testname", "", "", "", "")
	comm  = stakingtypes.CommissionRates{}
	msg1  = stakingtypes.NewMsgCreateValidator(sdk.ValAddress(pk1.Address()), pk1,
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 50), desc, comm, sdk.OneInt())
	msg2 = stakingtypes.NewMsgCreateValidator(sdk.ValAddress(pk2.Address()), pk1,
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 50), desc, comm, sdk.OneInt())
)

// GenTxTestSuite is a test suite to be used with gentx tests.
type GenTxTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	app            *simapp.SimApp
	encodingConfig simappparams.EncodingConfig
}

func (suite *GenTxTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)
	suite.ctx = app.BaseApp.NewContext(checkTx, abci.Header{})
	suite.app = app

	suite.encodingConfig = simapp.MakeEncodingConfig()
	// sdk.RegisterCodec(suite.encodingConfig.Amino)
}

func (suite *GenTxTestSuite) setAccountBalance(cdc *codec.Codec, addr sdk.AccAddress, amount int64) json.RawMessage {
	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr)
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	err := suite.app.BankKeeper.SetBalances(
		suite.ctx, addr, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 25)},
	)
	suite.Require().NoError(err)

	bankGenesisState := suite.app.BankKeeper.ExportGenesis(suite.ctx)
	bankGenesis, err := cdc.MarshalJSON(bankGenesisState)
	suite.Require().NoError(err)

	return bankGenesis
}

func (suite *GenTxTestSuite) TestSetGenTxsInAppGenesisState() {
	var (
		txBuilder = suite.encodingConfig.TxConfig.NewTxBuilder()
		cdc       *codec.Codec
		genTxs    []sdk.Tx
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"unregistered message",
			func() {
				msg := testdata.NewTestMsg(sdk.AccAddress("some-address"))
				err := txBuilder.SetMsgs(msg)
				suite.Require().NoError(err)
				tx := txBuilder.GetTx()
				genTxs = []sdk.Tx{tx}
			},
			false,
		},
		{
			"one genesis transaction",
			func() {
				err := txBuilder.SetMsgs(msg1)
				suite.Require().NoError(err)
				tx := txBuilder.GetTx()
				genTxs = []sdk.Tx{tx}
			},
			true,
		},
		{
			"two genesis transactions",
			func() {
				err := txBuilder.SetMsgs(msg1, msg2)
				suite.Require().NoError(err)
				tx := txBuilder.GetTx()
				genTxs = []sdk.Tx{tx}
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()
			cdc = suite.app.Codec()

			tc.malleate()
			appGenesisState, err := genutil.SetGenTxsInAppGenesisState(cdc, make(map[string]json.RawMessage), genTxs)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(appGenesisState[types.ModuleName])

				var genesisState types.GenesisState
				err := cdc.UnmarshalJSON(appGenesisState[types.ModuleName], &genesisState)
				suite.Require().NoError(err)
				suite.Require().NotNil(genesisState.GenTxs)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *GenTxTestSuite) TestValidateAccountInGenesis() {
	var (
		appGenesisState = make(map[string]json.RawMessage)
		cdc             *codec.Codec
		coins           sdk.Coins
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"no accounts",
			func() {
				coins = sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)}
			},
			false,
		},
		{
			"account without balance in the genesis state",
			func() {
				coins = sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)}
				appGenesisState[banktypes.ModuleName] = suite.setAccountBalance(cdc, addr2, 50)
			},
			false,
		},
		{
			"account without enough funds of default bond denom",
			func() {
				coins = sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 50)}
				appGenesisState[banktypes.ModuleName] = suite.setAccountBalance(cdc, addr1, 25)
			},
			false,
		},
		{
			"account with enough funds of default bond denom",
			func() {
				coins = sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)}
				appGenesisState[banktypes.ModuleName] = suite.setAccountBalance(cdc, addr1, 25)
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()
			cdc = suite.app.Codec()

			suite.app.StakingKeeper.SetParams(suite.ctx, stakingtypes.DefaultParams())
			stakingGenesisState := staking.ExportGenesis(suite.ctx, suite.app.StakingKeeper)
			suite.Require().Equal(stakingGenesisState.Params, stakingtypes.DefaultParams())
			stakingGenesis, err := cdc.MarshalJSON(stakingGenesisState)
			suite.Require().NoError(err)
			appGenesisState[stakingtypes.ModuleName] = stakingGenesis

			tc.malleate()
			err = genutil.ValidateAccountInGenesis(
				appGenesisState, banktypes.GenesisBalancesIterator{},
				addr1, coins, cdc,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *GenTxTestSuite) TestDeliverGenTxs() {
	var (
		genTxs    []json.RawMessage
		cdc       *codec.Codec
		txBuilder = suite.encodingConfig.TxConfig.NewTxBuilder()
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"no signature supplied",
			func() {
				err := txBuilder.SetMsgs(msg1)
				suite.Require().NoError(err)

				genTxs = make([]json.RawMessage, 1)
				tx, err := cdc.MarshalJSON(txBuilder.GetTx())
				suite.Require().NoError(err)
				genTxs[0] = tx
			},
			false,
		},
		{
			"unregistered message",
			func() {
				cdc.RegisterConcrete(testdata.TestMsg{}, "cosmos-sdk/Test", nil)
				msg := testdata.NewTestMsg(sdk.AccAddress("some-address"))

				err := txBuilder.SetMsgs(msg)
				suite.Require().NoError(err)

				genTxs = make([]json.RawMessage, 1)
				tx, err := cdc.MarshalJSON(txBuilder.GetTx())
				suite.Require().NoError(err)
				genTxs[0] = tx
			},
			false,
		},
		{
			"success",
			func() {
				_ = suite.setAccountBalance(cdc, addr1, 50)
				_ = suite.setAccountBalance(cdc, addr2, 0)

				msg := banktypes.NewMsgSend(addr1, addr2, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)})
				tx, err := helpers.GenTx(
					suite.encodingConfig.TxConfig,
					[]sdk.Msg{msg},
					sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)},
					helpers.DefaultGenTxGas,
					suite.ctx.ChainID(),
					[]uint64{0},
					[]uint64{0},
					priv1,
				)
				suite.Require().NoError(err)

				genTxs = make([]json.RawMessage, 1)
				genTx, err := cdc.MarshalJSON(tx)
				suite.Require().NoError(err)
				genTxs[0] = genTx
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()
			cdc = suite.app.Codec()

			tc.malleate()

			if tc.expPass {
				suite.Require().NotPanics(func() {
					genutil.DeliverGenTxs(
						suite.ctx, genTxs, suite.app.StakingKeeper, suite.app.BaseApp.DeliverTx,
						simapp.AminoJSONTxDecoder(suite.encodingConfig), suite.encodingConfig.TxConfig.TxEncoder(),
					)
				})
			} else {
				suite.Require().Panics(func() {
					genutil.DeliverGenTxs(
						suite.ctx, genTxs, suite.app.StakingKeeper, suite.app.BaseApp.DeliverTx,
						simapp.AminoJSONTxDecoder(suite.encodingConfig), suite.encodingConfig.TxConfig.TxEncoder(),
					)
				})
			}
		})
	}
}

func TestGenTxTestSuite(t *testing.T) {
	suite.Run(t, new(GenTxTestSuite))
}
