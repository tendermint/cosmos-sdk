package host

import (
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IsValidID defines regular expression to check if the string consist of
// characters in one of the following categories only:
// 	- Alphanumeric
// - `.`, `_`, `+`, `-`, `#`
// - `[`, `]`, `<`, `>`
var IsValidID = regexp.MustCompile(`[a-z\.\_\+\-\#\[\]\<\>]+$`).MatchString

// ICS 024 Identifier and Path Validation Implementation
//
// This file defines ValidateFn to validate identifier and path strings
// The spec for ICS 024 can be located here:
// https://github.com/cosmos/ics/tree/master/spec/ics-024-host-requirements

// ValidateFn function type to validate path and identifier bytestrings
type ValidateFn func(string) error

func defaultIdentifierValidator(id string, min, max int) error { //nolint:unparam
	// valid id MUST NOT contain "/" separator
	if strings.Contains(id, "/") {
		return sdkerrors.Wrapf(ErrInvalidID, "identifier %s cannot contain separator '/'", id)
	}
	// valid id must be between 9 and 20 characters
	if len(id) < min || len(id) > max {
		return sdkerrors.Wrapf(ErrInvalidID, "identifier %s has invalid length: %d, must be between %d-%d characters", id, len(id), min, max)
	}
	// valid id must contain only lower alphabetic characters
	if !IsValidID(id) {
		return sdkerrors.Wrapf(ErrInvalidID, "identifier %s must contain only lowercase alphabetic characters", id)
	}
	return nil
}

// ClientIdentifierValidator is the default validator function for Client identifiers.
// A valid Identifier must be between 9-20 characters and only contain lowercase
// alphabetic characters,
func ClientIdentifierValidator(id string) error {
	return defaultIdentifierValidator(id, 9, 20)
}

// ConnectionIdentifierValidator is the default validator function for Connection identifiers.
// A valid Identifier must be between 10-20 characters and only contain lowercase
// alphabetic characters,
func ConnectionIdentifierValidator(id string) error {
	return defaultIdentifierValidator(id, 10, 20)
}

// ChannelIdentifierValidator is the default validator function for Channel identifiers.
// A valid Identifier must be between 10-20 characters and only contain lowercase
// alphabetic characters,
func ChannelIdentifierValidator(id string) error {
	return defaultIdentifierValidator(id, 10, 20)
}

// PortIdentifierValidator is the default validator function for Port identifiers.
// A valid Identifier must be between 2-20 characters and only contain lowercase
// alphabetic characters,
func PortIdentifierValidator(id string) error {
	return defaultIdentifierValidator(id, 2, 20)
}

// NewPathValidator takes in a Identifier Validator function and returns
// a Path Validator function which requires path only has valid identifiers
// alphanumeric character strings, and "/" separators
func NewPathValidator(idValidator ValidateFn) ValidateFn {
	return func(path string) error {
		pathArr := strings.Split(path, "/")
		for _, p := range pathArr {
			// Each path element must either be valid identifier
			err := idValidator(p)
			if err != nil && !IsValidID(p) {
				return sdkerrors.Wrapf(ErrInvalidPath, "path %s contains invalid identifier or non-alphanumeric path element: %s", path, p)
			}
		}
		return nil
	}
}

// PathValidator takes in path string and validateswith def ault identifier rules.
// This is optimized by simply checking that all path elements are alphanumeric.
func PathValidator(path string) error {
	pathArr := strings.Split(path, "/")
	if pathArr[0] == path {
		return sdkerrors.Wrapf(ErrInvalidPath, "path %s doesn't contain any separator '/'", path)
	}

	for _, p := range pathArr {
		// Each path element must be alphanumeric and non-blank
		if strings.TrimSpace(p) == "" || !IsValidID(p) {
			return sdkerrors.Wrapf(ErrInvalidPath, "path %s contains an invalid non-alphanumeric character: '%s'", path, p)
		}
	}
	return nil
}
