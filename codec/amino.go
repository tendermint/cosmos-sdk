package codec

import (
	"bytes"
	"encoding/json"
	"fmt"

	amino "github.com/tendermint/go-amino"
	cryptoamino "github.com/tendermint/tendermint/crypto/encoding/amino"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

// Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all Tendermint crypto and evidence types registered.
//
// TODO: Consider removing this global.
var Cdc *Codec

func init() {
	cdc := New()
	RegisterCrypto(cdc)
	RegisterEvidences(cdc)
	Cdc = cdc.Seal()
}

// Codec defines a wrapper for an Amino codec that properly handles protobuf
// types with Any's
type Codec struct {
	Amino *amino.Codec
}

var _ JSONMarshaler = &Codec{}

func (cdc *Codec) Seal() *Codec {
	return &Codec{cdc.Amino.Seal()}
}

func New() *Codec {
	return &Codec{amino.NewCodec()}
}

// RegisterCrypto registers all crypto dependency types with the provided Amino
// codec.
func RegisterCrypto(cdc *Codec) {
	cryptoamino.RegisterAmino(cdc.Amino)
}

// RegisterEvidences registers Tendermint evidence types with the provided Amino
// codec.
func RegisterEvidences(cdc *Codec) {
	tmtypes.RegisterEvidences(cdc.Amino)
}

// MarshalJSONIndent provides a utility for indented JSON encoding of an object
// via an Amino codec. It returns an error if it cannot serialize or indent as
// JSON.
func MarshalJSONIndent(m JSONMarshaler, obj interface{}) ([]byte, error) {
	bz, err := m.MarshalJSON(obj)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err = json.Indent(&out, bz, "", "  "); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// MustMarshalJSONIndent executes MarshalJSONIndent except it panics upon failure.
func MustMarshalJSONIndent(m JSONMarshaler, obj interface{}) []byte {
	bz, err := MarshalJSONIndent(m, obj)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal JSON: %s", err))
	}

	return bz
}

func (cdc *Codec) marshalAnys(o interface{}) error {
	return types.UnpackInterfaces(o, types.AminoPacker{Cdc: cdc.Amino})
}

func (cdc *Codec) unmarshalAnys(o interface{}) error {
	return types.UnpackInterfaces(o, types.AminoUnpacker{Cdc: cdc.Amino})
}

func (cdc *Codec) jsonMarshalAnys(o interface{}) error {
	return types.UnpackInterfaces(o, types.AminoJSONPacker{Cdc: cdc.Amino})
}

func (cdc *Codec) jsonUnmarshalAnys(o interface{}) error {
	return types.UnpackInterfaces(o, types.AminoJSONUnpacker{Cdc: cdc.Amino})
}

func (cdc *Codec) MarshalBinaryBare(o interface{}) ([]byte, error) {
	err := cdc.marshalAnys(o)
	if err != nil {
		return nil, err
	}
	return cdc.Amino.MarshalBinaryBare(o)
}

func (cdc *Codec) MustMarshalBinaryBare(o interface{}) []byte {
	err := cdc.marshalAnys(o)
	if err != nil {
		panic(err)
	}
	return cdc.Amino.MustMarshalBinaryBare(o)
}

func (cdc *Codec) MarshalBinaryLengthPrefixed(o interface{}) ([]byte, error) {
	err := cdc.marshalAnys(o)
	if err != nil {
		return nil, err
	}
	return cdc.Amino.MarshalBinaryLengthPrefixed(o)
}

func (cdc *Codec) MustMarshalBinaryLengthPrefixed(o interface{}) []byte {
	err := cdc.marshalAnys(o)
	if err != nil {
		panic(err)
	}
	return cdc.Amino.MustMarshalBinaryLengthPrefixed(o)
}

func (cdc *Codec) UnmarshalBinaryBare(bz []byte, ptr interface{}) error {
	err := cdc.Amino.UnmarshalBinaryBare(bz, ptr)
	if err != nil {
		return err
	}
	return cdc.unmarshalAnys(ptr)
}

func (cdc *Codec) MustUnmarshalBinaryBare(bz []byte, ptr interface{}) {
	cdc.Amino.MustUnmarshalBinaryBare(bz, ptr)
	err := cdc.unmarshalAnys(ptr)
	if err != nil {
		panic(err)
	}
}

func (cdc *Codec) UnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) error {
	err := cdc.Amino.UnmarshalBinaryLengthPrefixed(bz, ptr)
	if err != nil {
		return err
	}
	return cdc.unmarshalAnys(ptr)
}

func (cdc *Codec) MustUnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) {
	cdc.Amino.MustUnmarshalBinaryLengthPrefixed(bz, ptr)
	err := cdc.unmarshalAnys(ptr)
	if err != nil {
		panic(err)
	}
}

func (cdc *Codec) MarshalJSON(o interface{}) ([]byte, error) {
	err := cdc.jsonMarshalAnys(o)
	if err != nil {
		return nil, err
	}
	return cdc.Amino.MarshalJSON(o)
}

func (cdc *Codec) MustMarshalJSON(o interface{}) []byte {
	err := cdc.jsonMarshalAnys(o)
	if err != nil {
		panic(err)
	}
	return cdc.Amino.MustMarshalJSON(o)
}

func (cdc *Codec) UnmarshalJSON(bz []byte, ptr interface{}) error {
	err := cdc.Amino.UnmarshalJSON(bz, ptr)
	if err != nil {
		return err
	}
	return cdc.jsonUnmarshalAnys(ptr)
}

func (cdc *Codec) MustUnmarshalJSON(bz []byte, ptr interface{}) {
	cdc.Amino.MustUnmarshalJSON(bz, ptr)
	err := cdc.jsonUnmarshalAnys(ptr)
	if err != nil {
		panic(err)
	}
}

func (*Codec) UnpackAny(*types.Any, interface{}) error {
	return fmt.Errorf("AminoCodec can't handle unpack protobuf Any's")
}

func (cdc *Codec) RegisterInterface(ptr interface{}, iopts *amino.InterfaceOptions) {
	cdc.Amino.RegisterInterface(ptr, iopts)
}

func (cdc *Codec) RegisterConcrete(o interface{}, name string, copts *amino.ConcreteOptions) {
	cdc.Amino.RegisterConcrete(o, name, copts)
}

func (cdc *Codec) MarshalJSONIndent(o interface{}, prefix, indent string) ([]byte, error) {
	err := cdc.jsonMarshalAnys(o)
	if err != nil {
		panic(err)
	}
	return cdc.Amino.MarshalJSONIndent(o, prefix, indent)
}
