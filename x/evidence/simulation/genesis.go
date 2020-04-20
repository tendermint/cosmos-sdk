package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
	"github.com/cosmos/cosmos-sdk/x/evidence/types"
)

// Simulation parameter constants
const (
	Evidence = "evidence"
)

// GenEvidences randomized Evidence
func GenEvidences(r *rand.Rand, accs []simtypes.Account) []exported.Evidence {
	totalEv := r.Intn(20)
	evidence := make([]exported.Evidence, totalEv)

	for i := 0; i < totalEv; i++ {
		simacc, _ := simtypes.RandomAcc(r, accs)
		eq := types.Equivocation{
			Height:           r.Int63n(200),
			Time:             time.Unix(r.Int63(), 0),
			Power:            r.Int63n(1000),
			ConsensusAddress: sdk.ConsAddress(simacc.Address),
		}
		evidence = append(evidence, eq)
	}
	return evidence
}

// RandomizedGenState generates a random GenesisState for evidence
func RandomizedGenState(simState *module.SimulationState) {
	var evidence []exported.Evidence
	simState.AppParams.GetOrGenerate(
		simState.Cdc, Evidence, &evidence, simState.Rand,
		func(r *rand.Rand) { evidence = GenEvidences(r, simState.Accounts) },
	)

	evidenceGenesis := types.GenesisState{
		Evidence: evidence,
	}

	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, codec.MustMarshalJSONIndent(simState.Cdc, evidenceGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(evidenceGenesis)
}
