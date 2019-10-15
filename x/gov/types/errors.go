package types

// DONTCOVER

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Codes for governance errors
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeUnknownProposal          sdk.CodeType = 1
	CodeInactiveProposal         sdk.CodeType = 2
	CodeAlreadyActiveProposal    sdk.CodeType = 3
	CodeAlreadyFinishedProposal  sdk.CodeType = 4
	CodeAddressNotStaked         sdk.CodeType = 5
	CodeInvalidContent           sdk.CodeType = 6
	CodeInvalidProposalType      sdk.CodeType = 7
	CodeInvalidVote              sdk.CodeType = 8
	CodeInvalidGenesis           sdk.CodeType = 9
	CodeInvalidProposalStatus    sdk.CodeType = 10
	CodeProposalHandlerNotExists sdk.CodeType = 11
)

// ErrUnknownProposal error for unknown proposals
func ErrUnknownProposal(codespace sdk.CodespaceType, proposalID uint64) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownProposal, fmt.Sprintf("unknown proposal with id %d", proposalID))
}

// ErrInactiveProposal error for inactive (i.e finalized) proposals
func ErrInactiveProposal(codespace sdk.CodespaceType, proposalID uint64) sdk.Error {
	return sdk.NewError(codespace, CodeInactiveProposal, fmt.Sprintf("inactive proposal with id %d", proposalID))
}

// ErrAlreadyActiveProposal error for proposals that are already active
func ErrAlreadyActiveProposal(codespace sdk.CodespaceType, proposalID uint64) sdk.Error {
	return sdk.NewError(codespace, CodeAlreadyActiveProposal, fmt.Sprintf("proposal %d has been already active", proposalID))
}

// ErrInvalidProposalContent error for invalid proposal title or description
func ErrInvalidProposalContent(cs sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(cs, CodeInvalidContent, fmt.Sprintf("invalid proposal content: %s", msg))
}

// ErrInvalidProposalType error for non registered proposal types
func ErrInvalidProposalType(codespace sdk.CodespaceType, proposalType string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidProposalType, fmt.Sprintf("proposal type '%s' is not valid", proposalType))
}

// ErrInvalidVote error for an invalid vote option
func ErrInvalidVote(codespace sdk.CodespaceType, voteOption fmt.Stringer) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidVote, fmt.Sprintf("'%v' is not a valid voting option", voteOption.String()))
}

// ErrInvalidGenesis error for an invalid governance GenesisState
func ErrInvalidGenesis(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidVote, msg)
}

// ErrNoProposalHandlerExists error when proposal handler is not defined
func ErrNoProposalHandlerExists(codespace sdk.CodespaceType, content interface{}) sdk.Error {
	return sdk.NewError(codespace, CodeProposalHandlerNotExists, fmt.Sprintf("'%T' does not have a corresponding handler", content))
}
