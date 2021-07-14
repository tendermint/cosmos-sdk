package keyring

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
)

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}

// RegisterLegacyAminoCodec registers concrete types and interfaces on the given codec.
// TODO how to remove Info entirely?
// rename to LegacyInfo
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*LegacyInfo)(nil), nil)
	cdc.RegisterConcrete(hd.BIP44Params{}, "crypto/keys/hd/BIP44Params", nil)
	cdc.RegisterConcrete(legacyLocalInfo{}, "crypto/keys/localInfo", nil)
	cdc.RegisterConcrete(legacyLedgerInfo{}, "crypto/keys/ledgerInfo", nil)
	cdc.RegisterConcrete(legacyOfflineInfo{}, "crypto/keys/offlineInfo", nil)
	cdc.RegisterConcrete(legacyMultiInfo{}, "crypto/keys/multiInfo", nil)
}
