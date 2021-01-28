package address

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MaxAddrLen is the maximum allowed length (in bytes) for an address.
const MaxAddrLen = 255

// LengthPrefixedStoreKey prefixes the address bytes with its length, this is used
// for variable-length components in store keys.
func LengthPrefixedStoreKey(bz []byte) ([]byte, error) {
	bzLen := len(bz)
	if bzLen == 0 {
		return bz, nil
	}

	if bzLen > MaxAddrLen {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address length should be max %d bytes, got %d", MaxAddrLen, bzLen)
	}

	return append([]byte{byte(bzLen)}, bz...), nil
}

// MustLengthPrefixedStoreKey is LengthPrefixedStoreKey with panic on error.
func MustLengthPrefixedStoreKey(bz []byte) []byte {
	res, err := LengthPrefixedStoreKey(bz)
	if err != nil {
		panic(err)
	}

	return res
}
