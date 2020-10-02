package tx

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// ConvertTxToStdTx converts a transaction to the legacy StdTx format
func ConvertTxToStdTx(codec *codec.LegacyAmino, tx signing.Tx) (legacytx.StdTx, error) {
	if stdTx, ok := tx.(legacytx.StdTx); ok {
		return stdTx, nil
	}

	aminoTxConfig := legacytx.StdTxConfig{Cdc: codec}
	builder := aminoTxConfig.NewTxBuilder()

	err := CopyTx(tx, builder)
	if err != nil {

		return legacytx.StdTx{}, err
	}

	stdTx, ok := builder.GetTx().(legacytx.StdTx)
	if !ok {
		return legacytx.StdTx{}, fmt.Errorf("expected %T, got %+v", legacytx.StdTx{}, builder.GetTx())
	}

	return stdTx, nil
}

// CopyTx copies a Tx to a new TxBuilder, allowing conversion between
// different transaction formats.
func CopyTx(tx signing.Tx, builder client.TxBuilder) error {
	err := builder.SetMsgs(tx.GetMsgs()...)
	if err != nil {
		return err
	}

	sigs, err := tx.GetSignaturesV2()
	if err != nil {
		return err
	}

	err = builder.SetSignatures(sigs...)
	if err != nil {
		return err
	}

	builder.SetMemo(tx.GetMemo())
	builder.SetFeeAmount(tx.GetFee())
	builder.SetGasLimit(tx.GetGas())

	return nil
}

func ConvertAndEncodeStdTx(txConfig client.TxConfig, stdTx legacytx.StdTx) ([]byte, error) {
	builder := txConfig.NewTxBuilder()

	var theTx sdk.Tx

	// check if we need a StdTx anyway, in that case don't copy
	if _, ok := builder.GetTx().(legacytx.StdTx); ok {
		theTx = stdTx
	} else {
		err := CopyTx(stdTx, builder)
		if err != nil {
			return nil, err
		}

		theTx = builder.GetTx()
	}

	return txConfig.TxEncoder()(theTx)
}
