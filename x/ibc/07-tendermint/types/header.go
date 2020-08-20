package types

import (
	"bytes"
	"time"

	tmtypes "github.com/tendermint/tendermint/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	commitmenttypes "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/types"
)

var _ clientexported.Header = Header{}

// ClientType defines that the Header is a Tendermint consensus algorithm
func (h Header) ClientType() clientexported.ClientType {
	return clientexported.Tendermint
}

// ConsensusState returns the updated consensus state associated with the header
func (h Header) ConsensusState() *ConsensusState {
	return &ConsensusState{
		Height:             uint64(h.GetHeight()),
		Timestamp:          h.GetTime(),
		Root:               commitmenttypes.NewMerkleRoot(h.Header.AppHash),
		NextValidatorsHash: h.Header.NextValidatorsHash,
	}
}

// GetHeight returns the current height. It returns 0 if the tendermint
// header is nil.
//
// NOTE: also referred as `sequence`
func (h Header) GetHeight() uint64 {
	if h.Header == nil {
		return 0
	}

	return uint64(h.GetHeight())
}

// GetTime returns the current block timestamp. It returns a zero time if
// the tendermint header is nil.
func (h Header) GetTime() time.Time {
	if h.SignedHeader.Header == nil {
		return time.Time{}
	}
	return h.SignedHeader.Header.Time
}

// ValidateBasic calls the SignedHeader ValidateBasic function and checks
// that validatorsets are not nil.
// NOTE: TrustedHeight and TrustedValidators may be empty when creating client
// with MsgCreateClient
func (h Header) ValidateBasic() error {
	if h.Header == nil {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "tendermint header cannot be nil")
	}
	tmHeader, err := tmtypes.HeaderFromProto(h.Header)
	if err != nil {
		return sdkerrors.Wrap(err, "header is not a tendermint header")
	}
	if err := tmHeader.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "header failed basic validation")
	}

	// TrustedHeight is less than Header for updates
	// and less than or equal to Header for misbehaviour
	if h.TrustedHeight > uint64(h.GetHeight()) {
		return sdkerrors.Wrapf(ErrInvalidHeaderHeight, "TrustedHeight %d must be less than or equal to header height %d",
			h.TrustedHeight, h.GetHeight())
	}

	if h.ValidatorSet == nil {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "validator set is nil")
	}
	tmValset, err := tmtypes.ValidatorSetFromProto(h.ValidatorSet)
	if err != nil {
		return sdkerrors.Wrap(err, "validator set is not tendermint validator set")
	}
	if !bytes.Equal(h.Header.ValidatorsHash, tmValset.Hash()) {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "validator set does not match hash")
	}
	return nil
}
