package keys

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

func listKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all keys",
		Long: `Return a list of all public keys stored by this key manager
along with their associated name and address.`,
		RunE: runListCmd,
	}
	cmd.Flags().Uint(client.FlagIndent, client.DefaultIndent, "Add indent to JSON response")
	return cmd
}

func runListCmd(cmd *cobra.Command, args []string) error {
	kb, err := NewKeyBaseFromHomeFlag()
	if err != nil {
		return err
	}

	infos, err := kb.List()
	if err == nil {
		printInfos(infos)
	}
	return err
}
