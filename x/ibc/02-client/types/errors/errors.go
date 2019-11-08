package errors

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// client error codes
const (
	DefaultCodespace sdk.CodespaceType = "client"

	CodeClientExists           sdk.CodeType = 101
	CodeClientNotFound         sdk.CodeType = 102
	CodeClientFrozen           sdk.CodeType = 103
	CodeConsensusStateNotFound sdk.CodeType = 104
	CodeInvalidConsensusState  sdk.CodeType = 105
	CodeClientTypeNotFound     sdk.CodeType = 106
	CodeInvalidClientType      sdk.CodeType = 107
	CodeRootNotFound           sdk.CodeType = 108
	CodeInvalidHeader          sdk.CodeType = 109
	CodeInvalidEvidence        sdk.CodeType = 110
)

// ErrClientExists implements sdk.Error
func ErrClientExists(codespace sdk.CodespaceType, clientID string) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeClientExists),
		fmt.Sprintf("client with ID %s already exists", clientID),
	)
}

// ErrClientNotFound implements sdk.Error
func ErrClientNotFound(codespace sdk.CodespaceType, clientID string) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeClientNotFound),
		fmt.Sprintf("client with ID %s not found", clientID),
	)
}

// ErrClientFrozen implements sdk.Error
func ErrClientFrozen(codespace sdk.CodespaceType, clientID string) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeClientFrozen),
		fmt.Sprintf("client with ID %s is frozen due to misbehaviour", clientID),
	)
}

// ErrConsensusStateNotFound implements sdk.Error
func ErrConsensusStateNotFound(codespace sdk.CodespaceType) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeConsensusStateNotFound),
		"consensus state not found",
	)
}

// ErrInvalidConsensus implements sdk.Error
func ErrInvalidConsensus(codespace sdk.CodespaceType) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeInvalidConsensusState),
		"invalid consensus state",
	)
}

// ErrClientTypeNotFound implements sdk.Error
func ErrClientTypeNotFound(codespace sdk.CodespaceType) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeClientTypeNotFound),
		"client type not found",
	)
}

// ErrInvalidClientType implements sdk.Error
func ErrInvalidClientType(codespace sdk.CodespaceType, msg string) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeInvalidClientType),
		msg,
	)
}

// ErrRootNotFound implements sdk.Error
func ErrRootNotFound(codespace sdk.CodespaceType) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeRootNotFound),
		"commitment root not found",
	)
}

// ErrInvalidHeader implements sdk.Error
func ErrInvalidHeader(codespace sdk.CodespaceType) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeInvalidHeader),
		"invalid header",
	)
}

// ErrInvalidEvidence implements sdk.Error
func ErrInvalidEvidence(codespace sdk.CodespaceType, msg string) error {
	return sdkerrors.Register(
		string(codespace),
		uint32(CodeInvalidEvidence),
		fmt.Sprintf("invalid evidence: %s", msg),
	)
}
