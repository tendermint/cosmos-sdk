package types

import (
	"fmt"
	"strconv"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

const (
	// SubModuleName defines the IBC client name
	SubModuleName string = "client"

	// RouterKey is the message route for IBC client
	RouterKey string = SubModuleName

	// QuerierRoute is the querier route for IBC client
	QuerierRoute string = SubModuleName

	// KeyNextClientSequence is the key used to store the next client sequence in
	// the keeper.
	KeyNextClientSequence = "nextClientSequence"
)

// FormatClientIdentifier returns the client identifier with the sequence appended.
func FormatClientIdentifier(clientType string, sequence uint64) string {
	return fmt.Sprintf("%s-%d", clientType, sequence)
}

// IsValidClientID return true if the client identifier is valid.
func IsValidClientID(clientID string) bool {
	_, _, err := ParseClientIdentifier(clientID)
	return err == nil
}

// ParseClientIdentifier parses the client type and sequence from the client identifier.
func ParseClientIdentifier(clientID string) (string, uint64, error) {
	splitStr := strings.Split(clientID, "-")
	if len(splitStr) != 2 {
		return "", 0, sdkerrors.Wrap(host.ErrInvalidID, "client identifier must be in format: `{client-type}-{N}`")
	}

	clientType := splitStr[0]
	if strings.TrimSpace(clientType) == "" {
		return "", 0, sdkerrors.Wrap(host.ErrInvalidID, "client identifier must be in format: `{client-type}-{N}` and client type cannot be blank")
	}

	sequence, err := strconv.ParseUint(splitStr[1], 10, 64)
	if err != nil {
		return "", 0, sdkerrors.Wrap(err, "failed to parse client identifier sequence")
	}

	return clientType, sequence, nil
}
