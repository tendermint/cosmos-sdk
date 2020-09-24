package ibctesting

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/exported"
	solomachinetypes "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/solomachine/types"
)

// Solomachine is a testing helper used to simulate a counterparty
// solo machine client.
type Solomachine struct {
	t *testing.T

	cdc         codec.BinaryMarshaler
	ClientID    string
	PrivateKeys []crypto.PrivKey // keys used for signing
	PublicKeys  []crypto.PubKey  // keys used for generating solo machine pub key
	PublicKey   crypto.PubKey    // key used for verification
	Sequence    uint64
	Time        uint64
	Diversifier string
}

// NewSolomachine returns a new solomachine instance with an `nKeys` amount of
// generated private/public key pairs and a sequence starting at 1. If nKeys
// is greater than 1 then a multisig public key is used.
func NewSolomachine(t *testing.T, cdc codec.BinaryMarshaler, clientID, diversifier string, nKeys uint64) *Solomachine {
	privKeys, pubKeys, pk := GenerateKeys(nKeys)

	return &Solomachine{
		t:           t,
		cdc:         cdc,
		ClientID:    clientID,
		PrivateKeys: privKeys,
		PublicKeys:  pubKeys,
		PublicKey:   pk,
		Sequence:    1,
		Time:        10,
		Diversifier: diversifier,
	}
}

// GenerateKeys generates a new set of private keys and public keys. If the
// number of keys is greater than one then the public key returned represents
// a multisig public key. The private keys are used for signing, the public
// keys are used for generating the public key and the public key is used for
// solo machine verification.
func GenerateKeys(n uint64) ([]crypto.PrivKey, []crypto.PubKey, crypto.PubKey) {
	if n == 0 {
		panic("cannot use zero private keys")
	}

	privKeys := make([]crypto.PrivKey, n)
	pubKeys := make([]crypto.PubKey, n)
	for i := uint64(0); i < n; i++ {
		privKeys[i] = secp256k1.GenPrivKey()
		pubKeys[i] = privKeys[i].PubKey()
	}

	var pk crypto.PubKey
	if len(privKeys) > 1 {
		// generate multi sig pk
		pk = kmultisig.NewLegacyAminoPubKey(int(n), pubKeys)
	} else {
		pk = privKeys[0].PubKey()
	}

	return privKeys, pubKeys, pk
}

// ClientState returns a new solo machine ClientState instance. Default usage does not allow update
// after governance proposal
func (solo *Solomachine) ClientState() *solomachinetypes.ClientState {
	return solomachinetypes.NewClientState(solo.Sequence, solo.ConsensusState(), false)
}

// ConsensusState returns a new solo machine ConsensusState instance
func (solo *Solomachine) ConsensusState() *solomachinetypes.ConsensusState {
	publicKey, err := tx.PubKeyToAny(solo.PublicKey)
	require.NoError(solo.t, err)

	return &solomachinetypes.ConsensusState{
		PublicKey:   publicKey,
		Diversifier: solo.Diversifier,
		Timestamp:   solo.Time,
	}
}

// GetHeight returns an exported.Height with Sequence as EpochHeight
func (solo *Solomachine) GetHeight() exported.Height {
	return clienttypes.NewHeight(0, solo.Sequence)
}

// CreateHeader generates a new private/public key pair and creates the
// necessary signature to construct a valid solo machine header.
func (solo *Solomachine) CreateHeader() *solomachinetypes.Header {
	// generate new private keys and signature for header
	newPrivKeys, newPubKeys, newPubKey := GenerateKeys(uint64(len(solo.PrivateKeys)))

	publicKey, err := tx.PubKeyToAny(newPubKey)
	require.NoError(solo.t, err)

	data := &solomachinetypes.HeaderData{
		NewPubKey:      publicKey,
		NewDiversifier: solo.Diversifier,
	}

	dataBz, err := solo.cdc.MarshalBinaryBare(data)
	require.NoError(solo.t, err)

	signBytes := &solomachinetypes.SignBytes{
		Sequence:    solo.Sequence,
		Timestamp:   solo.Time,
		Diversifier: solo.Diversifier,
		Data:        dataBz,
	}

	bz, err := solo.cdc.MarshalBinaryBare(signBytes)
	require.NoError(solo.t, err)

	sig := solo.GenerateSignature(bz)

	header := &solomachinetypes.Header{
		Sequence:       solo.Sequence,
		Timestamp:      solo.Time,
		Signature:      sig,
		NewPublicKey:   publicKey,
		NewDiversifier: solo.Diversifier,
	}

	// assumes successful header update
	solo.Sequence++
	solo.PrivateKeys = newPrivKeys
	solo.PublicKeys = newPubKeys
	solo.PublicKey = newPubKey

	return header
}

// CreateMisbehaviour constructs testing misbehaviour for the solo machine client
// by signing over two different data bytes at the same sequence.
func (solo *Solomachine) CreateMisbehaviour() *solomachinetypes.Misbehaviour {
	dataOne := []byte("DATA ONE")
	dataTwo := []byte("DATA TWO")

	signBytes := &solomachinetypes.SignBytes{
		Sequence:    solo.Sequence,
		Timestamp:   solo.Time,
		Diversifier: solo.Diversifier,
		Data:        dataOne,
	}

	bz, err := solo.cdc.MarshalBinaryBare(signBytes)
	require.NoError(solo.t, err)

	sig := solo.GenerateSignature(bz)
	signatureOne := solomachinetypes.SignatureAndData{
		Signature: sig,
		Data:      dataOne,
	}

	signBytes = &solomachinetypes.SignBytes{
		Sequence:    solo.Sequence,
		Timestamp:   solo.Time,
		Diversifier: solo.Diversifier,
		Data:        dataTwo,
	}

	bz, err = solo.cdc.MarshalBinaryBare(signBytes)
	require.NoError(solo.t, err)

	sig = solo.GenerateSignature(bz)
	signatureTwo := solomachinetypes.SignatureAndData{
		Signature: sig,
		Data:      dataTwo,
	}

	return &solomachinetypes.Misbehaviour{
		ClientId:     solo.ClientID,
		Sequence:     solo.Sequence,
		SignatureOne: &signatureOne,
		SignatureTwo: &signatureTwo,
	}
}

// GenerateSignature uses the stored private keys to generate a signature
// over the sign bytes with each key. If the amount of keys is greater than
// 1 then a multisig data type is returned.
func (solo *Solomachine) GenerateSignature(signBytes []byte) []byte {
	sigs := make([]signing.SignatureData, len(solo.PrivateKeys))
	for i, key := range solo.PrivateKeys {
		sig, err := key.Sign(signBytes)
		require.NoError(solo.t, err)

		sigs[i] = &signing.SingleSignatureData{
			Signature: sig,
		}
	}

	var sigData signing.SignatureData
	if len(sigs) == 1 {
		// single public key
		sigData = sigs[0]
	} else {
		// generate multi signature data
		multiSigData := multisig.NewMultisig(len(sigs))
		for i, sig := range sigs {
			multisig.AddSignature(multiSigData, sig, i)
		}

		sigData = multiSigData
	}

	protoSigData := signing.SignatureDataToProto(sigData)
	bz, err := solo.cdc.MarshalBinaryBare(protoSigData)
	require.NoError(solo.t, err)

	return bz
}
