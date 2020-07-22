// +build !test_amino

package params

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() EncodingConfig {
	cdc := codec.New()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewHybridCodec(cdc, interfaceRegistry)
	txGen := tx.NewTxConfig(marshaler, std.DefaultPublicKeyCodec{}, tx.DefaultSignModeHandler())

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txGen,
		Amino:             cdc,
	}
}
