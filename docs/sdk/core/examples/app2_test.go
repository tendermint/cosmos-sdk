package app

import (
	"testing"
	"github.com/tendermint/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

// Test encoding of app2Tx is correct with both msg types
func TestEncoding(t *testing.T) {
	// Create privkeys and addresses
	priv1 := crypto.GenPrivKeyEd25519()
	priv2 := crypto.GenPrivKeyEd25519()
	addr1 := priv1.PubKey().Address().Bytes()
	addr2 := priv2.PubKey().Address().Bytes()

	sendMsg := MsgSend{
		From: addr1,
		To: addr2,
		Amount: sdk.Coins{{"testCoins", sdk.NewInt(100)}},
	}

	// Construct transaction
	signBytes := sendMsg.GetSignBytes()
	sig, err := priv1.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	sendTxBefore := app2Tx{
		Msg: sendMsg,
		PubKey: priv1.PubKey(),
		Signature: sig,
	}

	cdc := NewCodec()
	testTxDecoder := tx2Decoder(cdc)

	encodedSendTx, err := cdc.MarshalBinary(sendTxBefore)

	require.Nil(t, err, "Error encoding sendTx")

	var tx1 sdk.Tx
	tx1, err = testTxDecoder(encodedSendTx)
	require.Nil(t, err, "Error decoding sendTx")

	sendTxAfter := tx1.(app2Tx)
	
	require.Equal(t, sendTxBefore, sendTxAfter, "Transaction changed after encoding/decoding")

	issueMsg := MsgIssue{
		Issuer: addr1,
		Receiver: addr2,
		Coin: sdk.Coin{"testCoin", sdk.NewInt(100)},
	}

	signBytes = issueMsg.GetSignBytes()
	sig, err = priv1.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	issueTxBefore := app2Tx{
		Msg: issueMsg,
		PubKey: priv1.PubKey(),
		Signature: sig,
	}

	encodedIssueTx, err2 := cdc.MarshalBinary(issueTxBefore)

	require.Nil(t, err2, "Error encoding issueTx")

	var tx2 sdk.Tx
	tx2, err2 = testTxDecoder(encodedIssueTx)
	require.Nil(t, err2, "Error decoding issue Tx")

	issueTxAfter := tx2.(app2Tx)

	require.Equal(t, issueTxBefore, issueTxAfter, "Transaction changed after encoding/decoding")

}