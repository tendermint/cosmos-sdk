package cli

import (
	"github.com/spf13/cobra"

	"github.com/KiraCore/cosmos-sdk/x/ibc/07-tendermint/types"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "Tendermint client transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	txCmd.AddCommand(
		NewCreateClientCmd(),
		NewUpdateClientCmd(),
		NewSubmitMisbehaviourCmd(),
	)

	return txCmd
}
