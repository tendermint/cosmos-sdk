package host

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	msg     string
	id      string
	expPass bool
}

func TestDefaultIdentifierValidator(t *testing.T) {
	testCases := []testCase{
		{"valid lowercase", "lowercaseid", true},
		{"valid id special chars", "._+-#[]<>._+-#[]<>", true},
		{"valid id lower and special chars", "lower._+-#[]<>", true},
		{"numeric id", "1234567890", true},
		{"uppercase id", "NOTLOWERCASE", true},
		{"numeric id", "1234567890", true},
		{"blank id", "               ", false},
		{"id length out of range", "1", false},
		{"path-like id", "lower/case/id", false},
		{"invalid id", "(clientid)", false},
		{"empty string", "", false},
	}

	for _, tc := range testCases {

		err := ClientIdentifierValidator(tc.id)
		err1 := ConnectionIdentifierValidator(tc.id)
		err2 := ChannelIdentifierValidator(tc.id)
		err3 := PortIdentifierValidator(tc.id)
		if tc.expPass {
			require.NoError(t, err, tc.msg)
			require.NoError(t, err1, tc.msg)
			require.NoError(t, err2, tc.msg)
			require.NoError(t, err3, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
			require.Error(t, err1, tc.msg)
			require.Error(t, err2, tc.msg)
			require.Error(t, err3, tc.msg)
		}
	}
}

func TestPathValidator(t *testing.T) {
	testCases := []testCase{
		{"valid lowercase", "/lowercaseid", true},
		{"numeric path", "/239123", true},
		{"valid id special chars", "/._+-#[]<>._+-#[]<>", true},
		{"valid id lower and special chars", "lower/._+-#[]<>", true},
		{"id length out of range", "/l", true},
		{"uppercase id", "/NOTLOWERCASE", true},
		{"invalid path", "lowercaseid", false},
		{"blank id", "/               ", false},
		{"id length out of range", "/123456789012345678901", false},
		{"invalid id", "/(clientid)", false},
		{"empty string", "", false},
		{"separators only", "////", false},
		{"just separator", "/", false},
	}

	for _, tc := range testCases {
		err := PathValidator(tc.id)
		if tc.expPass {
			seps := strings.Count(tc.id, "/")
			require.Equal(t, 1, seps)
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func TestCustomPathValidator(t *testing.T) {
	validateFn := NewPathValidator(func(path string) error {
		if !strings.HasPrefix(path, "id_") {
			return fmt.Errorf("identifier %s must start with 'id_", path)
		}
		return nil
	})

	testCases := []testCase{
		{"valid custom path", "/id_client/id_one", true},
		{"invalid path", "client", false},
		{"invalid custom path", "/client", false},
		{"invalid identifier", "/id_client/id_123456789012345678901", false},
		{"separators only", "////", false},
		{"just separator", "/", false},
	}

	for _, tc := range testCases {
		err := validateFn(tc.id)
		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}
