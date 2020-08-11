package proposal

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type Codec struct {
	codec.Marshaler

	// Keep reference to the amino codec to allow backwards compatibility along
	// with type, and interface registration.
	amino *codec.LegacyAmino
}

func NewCodec(amino *codec.LegacyAmino) *Codec {
	return &Codec{Marshaler: codec.NewAminoCodec(amino), amino: amino}
}

// ModuleCdc is the module codec.
var ModuleCdc *Codec

func init() {
	ModuleCdc = NewCodec(codec.New())

	RegisterCodec(ModuleCdc.amino)
	ModuleCdc.amino.Seal()
}

// RegisterCodec registers all necessary param module types with a given codec.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ParameterChangeProposal{},
	)
}
