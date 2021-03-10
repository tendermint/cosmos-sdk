package multisig

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// tmMultisig implements a K of N threshold multisig. It is used for
// Amino marshaling of LegacyAminoPubKey (see below for details).
//
// This struct is copy-pasted from:
// https://github.com/tendermint/tendermint/blob/v0.33.9/crypto/multisig/threshold_pubkey.go
//
// This struct was used in the SDK <=0.39. In 0.40 and the switch to protobuf,
// it has been converted to LegacyAminoPubKey. However, there's one difference:
// the threshold field was an `uint` before, and an `uint32` after. This caused
// amino marshaling to be breaking: amino marshals `uint32` as a JSON number,
// and `uint` as a JSON string.
//
// In this file, we're overriding LegacyAminoPubKey's default Amino marshaling
// by using this struct.
//
// ref: https://github.com/cosmos/cosmos-sdk/issues/8776
type tmMultisig struct {
	K       uint                 `json:"threshold"`
	PubKeys []cryptotypes.PubKey `json:"pubkeys"`
}

// protoToTm converts a LegacyAminoPubKey into a tmMultisig.
func protoToTm(protoPk *LegacyAminoPubKey) (tmMultisig, error) {
	var ok bool
	pks := make([]cryptotypes.PubKey, len(protoPk.PubKeys))
	for i, pk := range protoPk.PubKeys {
		pks[i], ok = pk.GetCachedValue().(cryptotypes.PubKey)
		if !ok {
			return tmMultisig{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expected %T, got %T", (cryptotypes.PubKey)(nil), pk.GetCachedValue())
		}
	}

	return tmMultisig{
		K:       uint(protoPk.Threshold),
		PubKeys: pks,
	}, nil
}

// tmToProto converts a tmMultisig into a LegacyAminoPubKey.
func tmToProto(tmPk tmMultisig) (*LegacyAminoPubKey, error) {
	var err error
	pks := make([]*types.Any, len(tmPk.PubKeys))
	for i, pk := range tmPk.PubKeys {
		pks[i], err = types.NewAnyWithValue(pk)
		if err != nil {
			return nil, err
		}
	}

	return &LegacyAminoPubKey{
		Threshold: uint32(tmPk.K),
		PubKeys:   pks,
	}, nil
}

// MarshalAmino overrides amino binary unmarshaling.
func (m LegacyAminoPubKey) MarshalAmino() (tmMultisig, error) {
	return protoToTm(&m)
}

// UnmarshalAmino overrides amino binary unmarshaling.
func (m *LegacyAminoPubKey) UnmarshalAmino(tmPk tmMultisig) error {
	protoPk, err := tmToProto(tmPk)
	if err != nil {
		return err
	}

	*m = *protoPk

	return nil
}

// MarshalAminoJSON overrides amino JSON unmarshaling.
func (m LegacyAminoPubKey) MarshalAminoJSON() (tmMultisig, error) {
	return protoToTm(&m)
}

// UnmarshalAminoJSON overrides amino JSON unmarshaling.
func (m *LegacyAminoPubKey) UnmarshalAminoJSON(tmPk tmMultisig) error {
	protoPk, err := tmToProto(tmPk)
	if err != nil {
		return err
	}

	*m = *protoPk

	return nil
}
