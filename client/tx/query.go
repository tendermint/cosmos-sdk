package tx

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// QueryTxCmd implements the default command for a tx query.
func QueryTxCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx [hash]",
		Short: "Matches this txhash over all committed blocks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// find the key to look up the account
			hashHexStr := args[0]

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			output, err := queryTx(cdc, cliCtx, hashHexStr)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().StringP(client.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")
	cmd.Flags().Bool(client.FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
	cmd.Flags().String(client.FlagChainID, "", "Chain ID of Tendermint node")
	return cmd
}

func queryTx(cdc *wire.Codec, cliCtx context.CLIContext, hashHexStr string) ([]byte, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	res, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return nil, err
	}

	info, err := formatTxResult(cdc, cliCtx, res)
	if err != nil {
		return nil, err
	}

	return wire.MarshalJSONIndent(cdc, info)
}

func formatTxResult(cdc *wire.Codec, cliCtx context.CLIContext, res *ctypes.ResultTx) (Info, error) {
	if !cliCtx.TrustNode {
		check, err := cliCtx.Certify(res.Height)
		if err != nil {
			return Info{}, err
		}

		err = res.Proof.Validate(check.Header.DataHash)
		if err != nil {
			return Info{}, err
		}
	}

	tx, err := parseTx(cdc, res.Tx)
	if err != nil {
		return Info{}, err
	}

	return Info{
		Hash:   res.Hash,
		Height: res.Height,
		Tx:     tx,
		Result: res.TxResult,
	}, nil
}

// Info is used to prepare info to display
type Info struct {
	Hash   common.HexBytes        `json:"hash"`
	Height int64                  `json:"height"`
	Tx     sdk.Tx                 `json:"tx"`
	Result abci.ResponseDeliverTx `json:"result"`
}

func parseTx(cdc *wire.Codec, txBytes []byte) (sdk.Tx, error) {
	var tx auth.StdTx

	err := cdc.UnmarshalBinary(txBytes, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// REST

// transaction query REST handler
func QueryTxRequestHandlerFn(cdc *wire.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hashHexStr := vars["hash"]

		output, err := queryTx(cdc, cliCtx, hashHexStr)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}
