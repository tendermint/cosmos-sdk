package commitment

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ContextKeyRemoteKVStore struct{}

func WithStore(ctx sdk.Context, store Store) sdk.Context {
	return ctx.WithValue(ContextKeyRemoteKVStore{}, store)
}
