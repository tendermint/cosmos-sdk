package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	eviexported "github.com/cosmos/cosmos-sdk/x/evidence/exported"
	"github.com/cosmos/cosmos-sdk/x/supply"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
)

var (
	_ auth.Codec     = (*Codec)(nil)
	_ supply.Codec   = (*Codec)(nil)
	_ evidence.Codec = (*Codec)(nil)
)

// Codec defines the application-level codec. This codec contains all the
// required module-specific codecs that are to be provided upon initialization.
type Codec struct {
	codec.Marshaler

	// Keep reference to the amino codec to allow backwards compatibility along
	// with type, and interface registration.
	amino *codec.Codec
}

func NewAppCodec(amino *codec.Codec) *Codec {
	return &Codec{Marshaler: codec.NewHybridCodec(amino), amino: amino}
}

// MarshalAccount marshals an Account interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (c *Codec) MarshalAccount(accI authexported.Account) ([]byte, error) {
	acc := &Account{}
	if err := acc.SetAccount(accI); err != nil {
		return nil, err
	}

	return c.Marshaler.MarshalBinaryLengthPrefixed(acc)
}

// UnmarshalAccount returns an Account interface from raw encoded account bytes
// of a Proto-based Account type. An error is returned upon decoding failure.
func (c *Codec) UnmarshalAccount(bz []byte) (authexported.Account, error) {
	acc := &Account{}
	if err := c.Marshaler.UnmarshalBinaryLengthPrefixed(bz, acc); err != nil {
		return nil, err
	}

	return acc.GetAccount(), nil
}

// MarshalAccountJSON JSON encodes an account object implementing the Account
// interface.
func (c *Codec) MarshalAccountJSON(acc authexported.Account) ([]byte, error) {
	return c.Marshaler.MarshalJSON(acc)
}

// UnmarshalAccountJSON returns an Account from JSON encoded bytes.
func (c *Codec) UnmarshalAccountJSON(bz []byte) (authexported.Account, error) {
	acc := &Account{}
	if err := c.Marshaler.UnmarshalJSON(bz, acc); err != nil {
		return nil, err
	}

	return acc.GetAccount(), nil
}

// MarshalSupply marshals a SupplyI interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (c *Codec) MarshalSupply(supplyI supplyexported.SupplyI) ([]byte, error) {
	supply := &Supply{}
	if err := supply.SetSupplyI(supplyI); err != nil {
		return nil, err
	}

	return c.Marshaler.MarshalBinaryLengthPrefixed(supply)
}

// UnmarshalSupply returns a SupplyI interface from raw encoded account bytes
// of a Proto-based SupplyI type. An error is returned upon decoding failure.
func (c *Codec) UnmarshalSupply(bz []byte) (supplyexported.SupplyI, error) {
	supply := &Supply{}
	if err := c.Marshaler.UnmarshalBinaryLengthPrefixed(bz, supply); err != nil {
		return nil, err
	}

	return supply.GetSupplyI(), nil
}

// MarshalSupplyJSON JSON encodes a supply object implementing the SupplyI
// interface.
func (c *Codec) MarshalSupplyJSON(supply supplyexported.SupplyI) ([]byte, error) {
	return c.Marshaler.MarshalJSON(supply)
}

// UnmarshalSupplyJSON returns a SupplyI from JSON encoded bytes.
func (c *Codec) UnmarshalSupplyJSON(bz []byte) (supplyexported.SupplyI, error) {
	supply := &Supply{}
	if err := c.Marshaler.UnmarshalJSON(bz, supply); err != nil {
		return nil, err
	}

	return supply.GetSupplyI(), nil
}

// MarshalEvidence marshals an EvidenceI interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (c *Codec) MarshalEvidence(evidenceI eviexported.EvidenceI) ([]byte, error) {
	evidence := &Evidence{}
	if err := evidence.SetEvidenceI(evidenceI); err != nil {
		return nil, err
	}

	return c.Marshaler.MarshalBinaryLengthPrefixed(evidence)
}

// UnmarshalEvidence returns an EvidenceI interface from raw encoded evidence
// bytes of a Proto-based EvidenceI type. An error is returned upon decoding
// failure.
func (c *Codec) UnmarshalEvidence(bz []byte) (eviexported.EvidenceI, error) {
	evidence := &Evidence{}
	if err := c.Marshaler.UnmarshalBinaryLengthPrefixed(bz, evidence); err != nil {
		return nil, err
	}

	return evidence.GetEvidenceI(), nil
}

// MarshalEvidenceJSON JSON encodes an evidence object implementing the EvidenceI
// interface.
func (c *Codec) MarshalEvidenceJSON(evidence eviexported.EvidenceI) ([]byte, error) {
	return c.Marshaler.MarshalJSON(evidence)
}

// UnmarshalEvidenceJSON returns an EvidenceI from JSON encoded bytes
func (c *Codec) UnmarshalEvidenceJSON(bz []byte) (eviexported.EvidenceI, error) {
	evidence := &Evidence{}
	if err := c.Marshaler.UnmarshalJSON(bz, evidence); err != nil {
		return nil, err
	}

	return evidence.GetEvidenceI(), nil
}

// ----------------------------------------------------------------------------
// necessary types and interfaces registered. This codec is provided to all the
// modules the application depends on.
//
// NOTE: This codec will be deprecated in favor of AppCodec once all modules are
// migrated.
func MakeCodec(bm module.BasicManager) *codec.Codec {
	cdc := codec.New()

	bm.RegisterCodec(cdc)
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}
