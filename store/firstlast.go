package store

import (
	"bytes"

	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/store/types"
)

// Gets the first item.
func First(st types.KVStore, start, end []byte) (kv cmn.KVPair, ok bool) {
	iter := st.Iterator(start, end)
	if !iter.Valid() {
		return kv, false
	}
	defer iter.Close()

	return cmn.KVPair{Key: iter.Key(), Value: iter.Value()}, true
}

// Gets the last item.  `end` is exclusive.
func Last(st types.KVStore, start, end []byte) (kv cmn.KVPair, ok bool) {
	iter := st.ReverseIterator(end, start)
	if !iter.Valid() {
		if v := st.Get(start); v != nil {
			return cmn.KVPair{Key: cp(start), Value: cp(v)}, true
		}
		return kv, false
	}
	defer iter.Close()

	if bytes.Equal(iter.Key(), end) {
		// Skip this one, end is exclusive.
		iter.Next()
		if !iter.Valid() {
			return kv, false
		}
	}

	return cmn.KVPair{Key: iter.Key(), Value: iter.Value()}, true
}

// copied from iavl/store.go
func cp(bz []byte) (ret []byte) {
	if bz == nil {
		return nil
	}
	ret = make([]byte, len(bz))
	copy(ret, bz)
	return ret
}
