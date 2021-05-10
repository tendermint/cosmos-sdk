package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

// RandomGenesisDefaultSendParam computes randomized allow all send transfers param for the bank module
func RandomGenesisDefaultSendParam(r *rand.Rand) bool {
	// 90% chance of transfers being enable or P(a) = 0.9 for success
	return r.Int63n(101) <= 90
}

// RandomGenesisSendParams randomized Parameters for the bank module
func RandomGenesisSendParams(r *rand.Rand) types.SendEnabledParams {
	params := types.DefaultParams()
	// 90% chance of transfers being DefaultSendEnabled=true or P(a) = 0.9 for success
	// 50% of the time add an additional denom specific record (P(b) = 0.475 = 0.5 * 0.95)
	if r.Int63n(101) <= 50 {
		// set send enabled 95% of the time
		bondEnabled := r.Int63n(101) <= 95
		params = params.SetSendEnabledParam(
			sdk.DefaultBondDenom,
			bondEnabled)
	}

	// overall probability of enabled for bond denom is 94.75% (P(a)+P(b) - P(a)*P(b))
	return params.SendEnabled
}

// RandomGenesisBalances returns a slice of account balances. Each account has
// a balance of simState.InitialStake for sdk.DefaultBondDenom.
func RandomGenesisBalances(simState *module.SimulationState) []types.Balance {
	genesisBalances := []types.Balance{}

	for _, acc := range simState.Accounts {
		genesisBalances = append(genesisBalances, types.Balance{
			Address: acc.Address.String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(simState.InitialStake))),
		})
	}

	return genesisBalances
}

func randomBondedAccountsGenesisBalances(randomGenesisBalance []types.Balance, numBonded int64, initialStake int64) []types.Balance {
	genesisBalances := randomGenesisBalance
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := simtypes.RandomAccounts(r, int(numBonded))
	for _, acc := range accounts {
		genesisBalances = append(genesisBalances, types.Balance{
			Address: acc.Address.String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(initialStake))),
		})
	}
	return genesisBalances
}

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(simState *module.SimulationState) {
	var sendEnabledParams types.SendEnabledParams
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.KeySendEnabled), &sendEnabledParams, simState.Rand,
		func(r *rand.Rand) { sendEnabledParams = RandomGenesisSendParams(r) },
	)

	var defaultSendEnabledParam bool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.KeyDefaultSendEnabled), &defaultSendEnabledParam, simState.Rand,
		func(r *rand.Rand) { defaultSendEnabledParam = RandomGenesisDefaultSendParam(r) },
	)

	numAccs := int64(len(simState.Accounts))
	totalSupply := sdk.NewInt(simState.InitialStake * (numAccs + simState.NumBonded))
	supply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, totalSupply))

	// genesisBalances := []types.Balance{}
	// s := rand.NewSource(1)
	// r := rand.New(s)
	// accounts := simtypes.RandomAccounts(r, int(simState.NumBonded))
	// fmt.Println("===================")
	// fmt.Println(len(accounts))
	// fmt.Println(len(simState.Accounts))
	// for _, acc := range accounts {
	// 	genesisBalances = append(genesisBalances, types.Balance{
	// 		Address: acc.Address.String(),
	// 		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(simState.InitialStake))),
	// 	})
	// }
	// simStateBalances := RandomGenesisBalances(simState)
	// for i := 0; i < len(simStateBalances); i++ {
	// 	genesisBalances = append(genesisBalances, simStateBalances...)
	// }

	genesisBalance := RandomGenesisBalances(simState)
	totalBalance := randomBondedAccountsGenesisBalances(genesisBalance, simState.NumBonded, simState.InitialStake)
	bankGenesis := types.GenesisState{
		Params: types.Params{
			SendEnabled:        sendEnabledParams,
			DefaultSendEnabled: defaultSendEnabledParam,
		},
		Balances: totalBalance,
		// Balances: RandomGenesisBalances(simState),
		Supply: supply,
	}

	paramsBytes, err := json.MarshalIndent(&bankGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated bank parameters:\n%s\n", paramsBytes)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bankGenesis)
}
