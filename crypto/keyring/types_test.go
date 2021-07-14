package keyring_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Test_writeReadLedgerInfo(t *testing.T) {
	tmpKey := make([]byte, secp256k1.PubKeySize)
	hexPK := "035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E143A"
	bz, err := hex.DecodeString(hexPK)
	require.NoError(t, err)
	copy(tmpKey[:], bz)

	pk := &secp256k1.PubKey{Key: tmpKey}
	path := hd.NewFundraiserParams(5, sdk.CoinType, 1)
	ledgerRecord := keyring.NewLedgerRecord(path)
	ledgerRecordItem := keyring.NewLedgerRecordItem(ledgerRecord)
	k, err := keyring.NewRecord("some_name", pk, ledgerRecordItem)
	require.NoError(t, err)

	l := k.GetLedger()
	require.NotNil(t, l)
	path = l.Path
	require.Equal(t, "m/44'/118'/5'/0/1", path.String())
	pubKey, err := k.GetPubKey()
	require.NoError(t, err)
	require.Equal(t,
		fmt.Sprintf("PubKeySecp256k1{%s}", hexPK),
		pubKey.String())

	// Serialize and restore
	encCfg := simapp.MakeTestEncodingConfig()
	serialized, err := encCfg.Codec.Marshal(k)
	require.NoError(t, err)
	var restoredRecord keyring.Record
	err = encCfg.Codec.Unmarshal(serialized, &restoredRecord)
	require.NoError(t, err)
	require.NotNil(t, restoredRecord)

	// Check both keys match
	require.Equal(t, k.Name, restoredRecord.Name)
	require.Equal(t, k.GetType(), restoredRecord.GetType())

	restoredPubKey, err := restoredRecord.GetPubKey()
	require.NoError(t, err)
	require.Equal(t, pubKey, restoredPubKey)

	l = restoredRecord.GetLedger()
	require.NotNil(t, l)
	restoredPath := l.GetPath()
	require.NoError(t, err)
	require.Equal(t, path, restoredPath)
}
