package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

func TestAccountRetriever(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	network := network.New(t, cfg)

	_, err := network.WaitForHeight(1)
	require.NoError(t, err)

	val := network.Validators[0]
	clientCtx := val.ClientCtx
	ar := types.AccountRetriever{}

	acc, err := ar.GetAccount(clientCtx, val.Address)
	require.NoError(t, err)
	require.NotNil(t, acc)

	require.NoError(t, ar.EnsureExists(clientCtx, val.Address))

	accNum, accSeq, err := ar.GetAccountNumberSequence(clientCtx, val.Address)
	require.NoError(t, err)
	require.Equal(t, accNum, 0)
	require.Equal(t, accSeq, 0)
}
