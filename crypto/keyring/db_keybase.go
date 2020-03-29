package keyring

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ LegacyKeybase = dbKeybase{}

// dbKeybase combines encryption and storage implementation to provide a
// full-featured key manager.
//
// NOTE: dbKeybase will be deprecated in favor of keyringKeybase.
type dbKeybase struct {
	base baseKeybase
	db   dbm.DB
}

// newDBKeybase creates a new dbKeybase instance using the provided DB for
// reading and writing keys.
func newDBKeybase(db dbm.DB, opts ...KeybaseOption) dbKeybase {
	return dbKeybase{
		base: newBaseKeybase(opts...),
		db:   db,
	}
}

// List returns the keys from storage in alphabetical order.
func (kb dbKeybase) List() ([]Info, error) {
	var res []Info

	iter, err := kb.db.Iterator(nil, nil)
	if err != nil {
		return nil, err
	}

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		key := string(iter.Key())

		// need to include only keys in storage that have an info suffix
		if strings.HasSuffix(key, infoSuffix) {
			info, err := unmarshalInfo(iter.Value())
			if err != nil {
				return nil, err
			}

			res = append(res, info)
		}
	}

	return res, nil
}

// Get returns the public information about one key.
func (kb dbKeybase) Get(name string) (Info, error) {
	bs, err := kb.db.Get(infoKey(name))
	if err != nil {
		return nil, err
	}

	if len(bs) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, name)
	}

	return unmarshalInfo(bs)
}

// ExportPrivateKeyObject returns a PrivKey object given the key name and
// passphrase. An error is returned if the key does not exist or if the Info for
// the key is invalid.
func (kb dbKeybase) ExportPrivateKeyObject(name string, passphrase string) (tmcrypto.PrivKey, error) {
	info, err := kb.Get(name)
	if err != nil {
		return nil, err
	}

	var priv tmcrypto.PrivKey

	switch i := info.(type) {
	case localInfo:
		linfo := i
		if linfo.PrivKeyArmor == "" {
			err = fmt.Errorf("private key not available")
			return nil, err
		}

		priv, _, err = crypto.UnarmorDecryptPrivKey(linfo.PrivKeyArmor, passphrase)
		if err != nil {
			return nil, err
		}

	case ledgerInfo, offlineInfo, multiInfo:
		return nil, errors.New("only works on local private keys")
	}

	return priv, nil
}

func (kb dbKeybase) Export(name string) (armor string, err error) {
	bz, err := kb.db.Get(infoKey(name))
	if err != nil {
		return "", err
	}

	if bz == nil {
		return "", fmt.Errorf("no key to export with name %s", name)
	}

	return crypto.ArmorInfoBytes(bz), nil
}

// ExportPubKey returns public keys in ASCII armored format. It retrieves a Info
// object by its name and return the public key in a portable format.
func (kb dbKeybase) ExportPubKey(name string) (armor string, err error) {
	bz, err := kb.db.Get(infoKey(name))
	if err != nil {
		return "", err
	}

	if bz == nil {
		return "", fmt.Errorf("no key to export with name %s", name)
	}

	info, err := unmarshalInfo(bz)
	if err != nil {
		return
	}

	return crypto.ArmorPubKeyBytes(info.GetPubKey().Bytes(), string(info.GetAlgo())), nil
}

// ExportPrivKey returns a private key in ASCII armored format.
// It returns an error if the key does not exist or a wrong encryption passphrase
// is supplied.
func (kb dbKeybase) ExportPrivKey(name string, decryptPassphrase string,
	encryptPassphrase string) (armor string, err error) {
	priv, err := kb.ExportPrivateKeyObject(name, decryptPassphrase)
	if err != nil {
		return "", err
	}

	info, err := kb.Get(name)
	if err != nil {
		return "", err
	}

	return crypto.EncryptArmorPrivKey(priv, encryptPassphrase, string(info.GetAlgo())), nil
}

// Update changes the passphrase with which an already stored key is
// encrypted.
//
// oldpass must be the current passphrase used for encryption,
// getNewpass is a function to get the passphrase to permanently replace
// the current passphrase
func (kb dbKeybase) Update(name, oldpass string, getNewpass func() (string, error)) error {
	info, err := kb.Get(name)
	if err != nil {
		return err
	}

	switch i := info.(type) {
	case localInfo:
		linfo := i

		key, _, err := crypto.UnarmorDecryptPrivKey(linfo.PrivKeyArmor, oldpass)
		if err != nil {
			return err
		}

		newpass, err := getNewpass()
		if err != nil {
			return err
		}

		kb.writeLocalKey(name, key, newpass, i.GetAlgo())
		return nil

	default:
		return fmt.Errorf("locally stored key required. Received: %v", reflect.TypeOf(info).String())
	}
}

// Close the underlying storage.
func (kb dbKeybase) Close() error {
	return kb.db.Close()
}

func (kb dbKeybase) writeLocalKey(name string, priv tmcrypto.PrivKey, passphrase string, algo SigningAlgo) Info {
	// encrypt private key using passphrase
	privArmor := crypto.EncryptArmorPrivKey(priv, passphrase, string(algo))

	// make Info
	pub := priv.PubKey()
	info := newLocalInfo(name, pub, privArmor, algo)

	kb.writeInfo(name, info)
	return info
}

func (kb dbKeybase) writeInfo(name string, info Info) {
	// write the info by key
	key := infoKey(name)
	serializedInfo := marshalInfo(info)

	kb.db.SetSync(key, serializedInfo)

	// store a pointer to the infokey by address for fast lookup
	kb.db.SetSync(addrKey(info.GetAddress()), key)
}

func addrKey(address sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("%s.%s", address.String(), addressSuffix))
}

func infoKey(name string) []byte {
	return []byte(fmt.Sprintf("%s.%s", name, infoSuffix))
}
