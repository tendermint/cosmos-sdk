package keys

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

	keys "github.com/tendermint/go-crypto/keys"
	"github.com/tendermint/tmlibs/cli"
	dbm "github.com/tendermint/tmlibs/db"

	"github.com/cosmos/cosmos-sdk/client"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// KeyDBName is the directory under root where we store the keys
const KeyDBName = "keys"

// keybase is used to make GetKeyBase a singleton
var keybase keys.Keybase

// initialize a keybase based on the configuration
func GetKeyBase() (keys.Keybase, error) {
	rootDir := viper.GetString(cli.HomeFlag)
	return GetKeyBaseFromDir(rootDir)
}

// initialize a keybase based on the configuration
func GetKeyBaseFromDir(rootDir string) (keys.Keybase, error) {
	if keybase == nil {
		db, err := dbm.NewGoLevelDB(KeyDBName, filepath.Join(rootDir, "keys"))
		if err != nil {
			return nil, err
		}
		keybase = client.GetKeyBase(db)
	}
	return keybase, nil
}

// used to set the keybase manually in test
func SetKeyBase(kb keys.Keybase) {
	keybase = kb
}

// used for outputting keys.Info over REST
type KeyOutput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
}

func Bech32CosmosKeysOutput(infos []keys.Info) ([]KeyOutput, error) {
	kos := make([]KeyOutput, len(infos))
	for i, info := range infos {
		ko, err := Bech32CosmosKeyOutput(info)
		if err != nil {
			return nil, err
		}
		kos[i] = ko
	}
	return kos, nil
}

func Bech32CosmosKeyOutput(info keys.Info) (KeyOutput, error) {
	bechAccount, err := sdk.Bech32CosmosifyAcc(sdk.Address(info.PubKey.Address().Bytes()))
	if err != nil {
		return KeyOutput{}, err
	}
	bechPubKey, err := sdk.Bech32CosmosifyAccPub(info.PubKey)
	if err != nil {
		return KeyOutput{}, err
	}
	return KeyOutput{
		Name:    info.Name,
		Address: bechAccount,
		PubKey:  bechPubKey,
	}, nil
}

func printInfo(info keys.Info) {
	ko, err := Bech32CosmosKeyOutput(info)
	if err != nil {
		panic(err)
	}
	switch viper.Get(cli.OutputFlag) {
	case "text":
		fmt.Printf("NAME:\tADDRESS:\t\t\t\t\t\tPUBKEY:\n")
		printKeyOutput(ko)
	case "json":
		out, err := json.MarshalIndent(ko, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
}

func printInfos(infos []keys.Info) {
	kos, err := Bech32CosmosKeysOutput(infos)
	if err != nil {
		panic(err)
	}
	switch viper.Get(cli.OutputFlag) {
	case "text":
		fmt.Printf("NAME:\tADDRESS:\t\t\t\t\t\tPUBKEY:\n")
		for _, ko := range kos {
			printKeyOutput(ko)
		}
	case "json":
		out, err := json.MarshalIndent(kos, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
}

func printKeyOutput(ko KeyOutput) {
	fmt.Printf("%s\t%s\t%s\n", ko.Name, ko.Address, ko.PubKey)
}
