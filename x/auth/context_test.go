package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContextWithSigners(t *testing.T) {
	input := setupTestInput()

	_, _, addr1 := newTestKeyPubAddr()
	_, _, addr2 := newTestKeyPubAddr()
	acc1 := NewBaseAccountWithAddress(addr1)
	acc1.SetSequence(7132)
	acc2 := NewBaseAccountWithAddress(addr2)
	acc2.SetSequence(8821)

	// new ctx has no signers
	signers := GetSigners(input.ctx)
	require.Equal(t, 0, len(signers))

	ctx2 := WithSigners(input.ctx, []Account{&acc1, &acc2})

	// original context is unchanged
	signers = GetSigners(input.ctx)
	require.Equal(t, 0, len(signers))

	// new context has signers
	signers = GetSigners(ctx2)
	require.Equal(t, 2, len(signers))
	require.Equal(t, acc1, *(signers[0].(*BaseAccount)))
	require.Equal(t, acc2, *(signers[1].(*BaseAccount)))
}
