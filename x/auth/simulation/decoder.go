package simulation

import (
	"bytes"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"
	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/KiraCore/cosmos-sdk/codec"
	"github.com/KiraCore/cosmos-sdk/x/auth/types"
)

type AuthUnmarshaler interface {
	UnmarshalAccount([]byte) (types.AccountI, error)
	GetCodec() codec.BinaryMarshaler
}

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding auth type.
func NewDecodeStore(ak AuthUnmarshaler) func(kvA, kvB tmkv.Pair) string {
	return func(kvA, kvB tmkv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.AddressStoreKeyPrefix):
			accA, err := ak.UnmarshalAccount(kvA.Value)
			if err != nil {
				panic(err)
			}

			accB, err := ak.UnmarshalAccount(kvB.Value)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", accA, accB)

		case bytes.Equal(kvA.Key, types.GlobalAccountNumberKey):
			var globalAccNumberA, globalAccNumberB gogotypes.UInt64Value
			ak.GetCodec().MustUnmarshalBinaryBare(kvA.Value, &globalAccNumberA)
			ak.GetCodec().MustUnmarshalBinaryBare(kvB.Value, &globalAccNumberB)

			return fmt.Sprintf("GlobalAccNumberA: %d\nGlobalAccNumberB: %d", globalAccNumberA, globalAccNumberB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
