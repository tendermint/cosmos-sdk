package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	crypto "github.com/tendermint/go-crypto"
	"github.com/tendermint/go-crypto/keys"
	"github.com/tendermint/go-crypto/keys/words"
	cfg "github.com/tendermint/tendermint/config"
	gaiacfg "github.com/cosmos/cosmos-sdk/config"
	"github.com/tendermint/tendermint/p2p"
	tmtypes "github.com/tendermint/tendermint/types"
	pvm "github.com/tendermint/tendermint/privval"
	tmcli "github.com/tendermint/tmlibs/cli"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"

	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	gc "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
)

// genesis piece structure for creating combined genesis
type GenesisTx struct {
	NodeID    string                   `json:"node_id"`
	IP        string                   `json:"ip"`
	Validator tmtypes.GenesisValidator `json:"validator"`
	AppGenTx  json.RawMessage          `json:"app_gen_tx"`
}

// Names of input parameters coming from app
/*type GenTxFlagNames struct {
	FlagName       string
	FlagClientHome string
	FlagOWK        string
	FlagIP         string
}*/

var (
	FlagName       = "name"
	FlagClientHome = "home-client"
	FlagOWK        = "owk"

	// bonded tokens given to genesis validators/accounts
	FreeFermionVal  = int64(100)
	FreeFermionsAcc = int64(50)
)

// Storage for init command input parameters
type InitConfig struct {
	ChainID   string
	GenTxs    bool
	GenTxsDir string
	Overwrite bool
}

var (
	FlagOverwrite = "overwrite"
	FlagGenTxs    = "gen-txs"
	FlagIP        = "ip"
	FlagChainID   = "chain-id"
)

// get cmd to initialize all files for tendermint and application
func GenTxCmd(ctx *Context, cdc *wire.Codec, appInit AppInit) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-tx",
		Short: "Create genesis transaction file (under [--home]/config/gentx/gentx-[nodeID].json)",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, args []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(tmcli.HomeFlag))
<<<<<<< HEAD

			ip := viper.GetString(FlagIP)
			if len(ip) == 0 {
				eip, err := externalIP()
				if err != nil {
					return err
				}
				ip = eip
			}

			genTxConfig := gc.GenTxConfig{
				viper.GetString(FlagName),
				viper.GetString(FlagClientHome),
				viper.GetBool(FlagOWK),
				ip,
			}
			cliPrint, genTxFile, err := gentxWithConfig(ctx, cdc, appInit, config, genTxConfig)
=======
			cliPrint, genTxFile, err := gentxWithConfig(config,nil,ctx,cdc,appInit)
>>>>>>> Added testnet command
			if err != nil {
				return err
			}
			toPrint := struct {
				AppMessage json.RawMessage `json:"app_message"`
				GenTxFile  json.RawMessage `json:"gen_tx_file"`
			}{
				cliPrint,
				genTxFile,
			}
			out, err := wire.MarshalJSONIndent(cdc, toPrint)
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}
	cmd.Flags().String(FlagIP, "", "external facing IP to use if left blank IP will be retrieved from this machine")
	cmd.Flags().AddFlagSet(appInit.FlagsAppGenTx)
	return cmd
}

<<<<<<< HEAD
func gentxWithConfig(ctx *Context, cdc *wire.Codec, appInit AppInit, config *cfg.Config, genTxConfig gc.GenTxConfig) (
=======
func gentxWithConfig(config *cfg.Config, gaiaConfig *gaiacfg.Config, ctx *Context, cdc *wire.Codec, appInit AppInit) (
>>>>>>> Added testnet command
	cliPrint json.RawMessage, genTxFile json.RawMessage, err error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return
	}
	nodeID := string(nodeKey.ID())
	pubKey := readOrCreatePrivValidator(config)

<<<<<<< HEAD
	appGenTx, cliPrint, validator, err := appInit.AppGenTx(cdc, pubKey, genTxConfig)
=======
	appGenTx, cliPrint, validator, err := appInit.AppGenTx(cdc, pubKey, gaiaConfig)
>>>>>>> Added testnet command
	if err != nil {
		return
	}

<<<<<<< HEAD
	tx := GenesisTx{
		NodeID:    nodeID,
		IP:        genTxConfig.IP,
=======
	ip := viper.GetString(flagIP)
	if len(ip) == 0 {
		ip, err = externalIP()
		if err != nil {
			return
		}
	}

	tx := GenesisTx{
		NodeID:    nodeID,
		IP:        ip,
>>>>>>> Added testnet command
		Validator: validator,
		AppGenTx:  appGenTx,
	}
	bz, err := wire.MarshalJSONIndent(cdc, tx)
	if err != nil {
		return
	}
	genTxFile = json.RawMessage(bz)
	name := fmt.Sprintf("gentx-%v.json", nodeID)
	writePath := filepath.Join(config.RootDir, "config", "gentx")
	file := filepath.Join(writePath, name)
	err = cmn.EnsureDir(writePath, 0700)
	if err != nil {
		return
	}
	err = cmn.WriteFile(file, bz, 0644)
	if err != nil {
		return
	}

	return
}

// get cmd to initialize all files for tendermint and application
func InitCmd(ctx *Context, cdc *wire.Codec, appInit AppInit) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize genesis config, priv-validator file, and p2p-node file",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(tmcli.HomeFlag))
<<<<<<< HEAD
			initConfig := InitConfig{
				viper.GetString(flagChainID),
				viper.GetBool(flagGenTxs),
				filepath.Join(config.RootDir, "config", "gentx"),
				viper.GetBool(flagOverwrite),
			}

			chainID, nodeID, appMessage, err := initWithConfig(ctx, cdc, appInit, config, initConfig)
=======
			chainID, nodeID, appMessage, err := initWithConfig(config,nil, ctx,cdc,appInit)
>>>>>>> Added testnet command
			if err != nil {
				return err
			}
			// print out some key information
			toPrint := struct {
				ChainID    string          `json:"chain_id"`
				NodeID     string          `json:"node_id"`
				AppMessage json.RawMessage `json:"app_message"`
			}{
				chainID,
				nodeID,
				appMessage,
			}
			out, err := wire.MarshalJSONIndent(cdc, toPrint)
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}
	cmd.Flags().BoolP(flagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().String(flagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().Bool(flagGenTxs, false, "apply genesis transactions from [--home]/config/gentx/")
	cmd.Flags().AddFlagSet(appInit.FlagsAppGenState)
	cmd.Flags().AddFlagSet(appInit.FlagsAppGenTx) // need to add this flagset for when no GenTx's provided
	cmd.AddCommand(GenTxCmd(ctx, cdc, appInit))
	return cmd
}

<<<<<<< HEAD
func initWithConfig(ctx *Context, cdc *wire.Codec, appInit AppInit, config *cfg.Config, initConfig InitConfig) (
=======
func initWithConfig(config *cfg.Config, gaiaConfig *gaiacfg.Config, ctx *Context, cdc *wire.Codec, appInit AppInit) (
>>>>>>> Added testnet command
	chainID string, nodeID string, appMessage json.RawMessage, err error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return
	}
	nodeID = string(nodeKey.ID())
	pubKey := readOrCreatePrivValidator(config)
<<<<<<< HEAD

	if initConfig.ChainID == "" {
		initConfig.ChainID = fmt.Sprintf("test-chain-%v", cmn.RandStr(6))
	}
	chainID = initConfig.ChainID

	genFile := config.GenesisFile()
	if !initConfig.Overwrite && cmn.FileExists(genFile) {
=======
	if gaiaConfig == nil {
		gaiaConfig = gaiacfg.DefaultConfig()
		gaiaConfig.Init.ChainID = viper.GetString(flagChainID)
		gaiaConfig.Init.Overwrite = viper.GetBool(flagOverwrite)
		gaiaConfig.Init.GenTxs = viper.GetBool(flagGenTxs)
		gaiaConfig.Init.GenTxsDir = filepath.Join(config.RootDir, "config", "gentx")
	}

	if gaiaConfig.Init.ChainID == "" {
		gaiaConfig.Init.ChainID = cmn.Fmt("test-chain-%v", cmn.RandStr(6))
	}
	chainID = gaiaConfig.Init.ChainID

	genFile := config.GenesisFile()
	if !gaiaConfig.Init.Overwrite && cmn.FileExists(genFile) {
>>>>>>> Added testnet command
		err = fmt.Errorf("genesis.json file already exists: %v", genFile)
		return
	}

	// process genesis transactions, or otherwise create one for defaults
	var appGenTxs []json.RawMessage
	var validators []tmtypes.GenesisValidator
	var persistentPeers string

<<<<<<< HEAD
	if initConfig.GenTxs {
		validators, appGenTxs, persistentPeers, err = processGenTxs(initConfig.GenTxsDir, cdc, appInit)
=======
	if gaiaConfig.Init.GenTxs {
		genTxsDir := gaiaConfig.Init.GenTxsDir
		validators, appGenTxs, persistentPeers, err = processGenTxs(genTxsDir, cdc, appInit)
>>>>>>> Added testnet command
		if err != nil {
			return
		}
		config.P2P.PersistentPeers = persistentPeers
		configFilePath := filepath.Join(config.RootDir, "config", "config.toml")
		cfg.WriteConfigFile(configFilePath, config)
	} else {
<<<<<<< HEAD
		genTxConfig := gc.GenTxConfig{
			viper.GetString(FlagName),
			viper.GetString(FlagClientHome),
			viper.GetBool(FlagOWK),
			"127.0.0.1",
		}
		appGenTx, am, validator, err := appInit.AppGenTx(cdc, pubKey, genTxConfig)
=======
		appGenTx, am, validator, err := appInit.AppGenTx(cdc, pubKey, nil)
>>>>>>> Added testnet command
		appMessage = am
		if err != nil {
			return "", "", nil, err
		}
		validators = []tmtypes.GenesisValidator{validator}
		appGenTxs = []json.RawMessage{appGenTx}
	}

	appState, err := appInit.AppGenState(cdc, appGenTxs)
	if err != nil {
		return
	}

<<<<<<< HEAD
	err = writeGenesisFile(cdc, genFile, initConfig.ChainID, validators, appState)
=======
	err = writeGenesisFile(cdc, genFile, gaiaConfig.Init.ChainID, validators, appState)
>>>>>>> Added testnet command
	if err != nil {
		return
	}

	return
}

// append a genesis-piece
func processGenTxs(genTxsDir string, cdc *wire.Codec, appInit AppInit) (
	validators []tmtypes.GenesisValidator, appGenTxs []json.RawMessage, persistentPeers string, err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return
	}

	genTxs := make(map[string]GenesisTx)
	var nodeIDs []string
	for _, fo := range fos {
		filename := path.Join(genTxsDir, fo.Name())
		if !fo.IsDir() && (path.Ext(filename) != ".json") {
			continue
		}

		// get the genTx
		var bz []byte
		bz, err = ioutil.ReadFile(filename)
		if err != nil {
			return
		}
		var genTx GenesisTx
		err = cdc.UnmarshalJSON(bz, &genTx)
		if err != nil {
			return
		}

		genTxs[genTx.NodeID] = genTx
		nodeIDs = append(nodeIDs, genTx.NodeID)
	}

	sort.Strings(nodeIDs)

	for _, nodeID := range nodeIDs {
		genTx := genTxs[nodeID]

		// combine some stuff
		validators = append(validators, genTx.Validator)
		appGenTxs = append(appGenTxs, genTx.AppGenTx)

		// Add a persistent peer
		comma := ","
		if len(persistentPeers) == 0 {
			comma = ""
		}
		persistentPeers += fmt.Sprintf("%s%s@%s:46656", comma, genTx.NodeID, genTx.IP)
	}

	return
}

//________________________________________________________________________________________

// read of create the private key file for this config
func readOrCreatePrivValidator(tmConfig *cfg.Config) crypto.PubKey {
	// private validator
	privValFile := tmConfig.PrivValidatorFile()
	var privValidator *pvm.FilePV
	if cmn.FileExists(privValFile) {
		privValidator = pvm.LoadFilePV(privValFile)
	} else {
		privValidator = pvm.GenFilePV(privValFile)
		privValidator.Save()
	}
	return privValidator.GetPubKey()
}

// create the genesis file
func writeGenesisFile(cdc *wire.Codec, genesisFile, chainID string, validators []tmtypes.GenesisValidator, appState json.RawMessage) error {
	genDoc := tmtypes.GenesisDoc{
		ChainID:    chainID,
		Validators: validators,
	}
	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}
	if err := genDoc.SaveAs(genesisFile); err != nil {
		return err
	}
	return addAppStateToGenesis(cdc, genesisFile, appState)
}

// Add one line to the genesis file
func addAppStateToGenesis(cdc *wire.Codec, genesisConfigPath string, appState json.RawMessage) error {
	bz, err := ioutil.ReadFile(genesisConfigPath)
	if err != nil {
		return err
	}
	out, err := AppendJSON(cdc, bz, "app_state", appState)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(genesisConfigPath, out, 0600)
}

//_____________________________________________________________________

// Core functionality passed from the application to the server init command
type AppInit struct {

	// flags required for application init functions
	FlagsAppGenState *pflag.FlagSet
	FlagsAppGenTx    *pflag.FlagSet

	// create the application genesis tx
<<<<<<< HEAD
	AppGenTx func(cdc *wire.Codec, pk crypto.PubKey, genTxConfig gc.GenTxConfig) (
=======
	AppGenTx func(cdc *wire.Codec, pk crypto.PubKey, gaiaConfig *gaiacfg.Config) (
>>>>>>> Added testnet command
		appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error)

	// AppGenState creates the core parameters initialization. It takes in a
	// pubkey meant to represent the pubkey of the validator of this machine.
	AppGenState func(cdc *wire.Codec, appGenTxs []json.RawMessage) (appState json.RawMessage, err error)
}

//_____________________________________________________________________

// simple default application init
var DefaultAppInit = AppInit{
	AppGenTx:    SimpleAppGenTx,
	AppGenState: SimpleAppGenState,
}

// simple genesis tx
type SimpleGenTx struct {
	Addr sdk.Address `json:"addr"`
}

// Generate a genesis transaction
<<<<<<< HEAD
func SimpleAppGenTx(cdc *wire.Codec, pk crypto.PubKey, genTxConfig gc.GenTxConfig) (
=======
func SimpleAppGenTx(cdc *wire.Codec, pk crypto.PubKey, gaiaConfig *gaiacfg.Config) (
>>>>>>> Added testnet command
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {

	var addr sdk.Address
	var secret string
	addr, secret, err = GenerateCoinKey()
	if err != nil {
		return
	}

	var bz []byte
	simpleGenTx := SimpleGenTx{addr}
	bz, err = cdc.MarshalJSON(simpleGenTx)
	if err != nil {
		return
	}
	appGenTx = json.RawMessage(bz)

	mm := map[string]string{"secret": secret}
	bz, err = cdc.MarshalJSON(mm)
	if err != nil {
		return
	}
	cliPrint = json.RawMessage(bz)

	validator = tmtypes.GenesisValidator{
		PubKey: pk,
		Power:  10,
	}
	return
}

// create the genesis app state
func SimpleAppGenState(cdc *wire.Codec, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	if len(appGenTxs) != 1 {
		err = errors.New("must provide a single genesis transaction")
		return
	}

	var genTx SimpleGenTx
	err = cdc.UnmarshalJSON(appGenTxs[0], &genTx)
	if err != nil {
		return
	}

	appState = json.RawMessage(fmt.Sprintf(`{
  "accounts": [{
    "address": "%s",
    "coins": [
      {
        "denom": "mycoin",
        "amount": 9007199254740992
      }
    ]
  }]
}`, genTx.Addr.String()))
	return
}

//___________________________________________________________________________________________

// GenerateCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func GenerateCoinKey() (sdk.Address, string, error) {

	// construct an in-memory key store
	codec, err := words.LoadCodec("english")
	if err != nil {
		return nil, "", err
	}
	keybase := keys.New(
		dbm.NewMemDB(),
		codec,
	)

	// generate a private key, with recovery phrase
	info, secret, err := keybase.Create("name", "pass", keys.AlgoEd25519)
	if err != nil {
		return nil, "", err
	}
	addr := info.PubKey.Address()
	return addr, secret, nil
}

// GenerateSaveCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func GenerateSaveCoinKey(clientRoot, keyName, keyPass string, overwrite bool) (sdk.Address, string, error) {

	// get the keystore from the client
	keybase, err := clkeys.GetKeyBaseFromDir(clientRoot)
	if err != nil {
		return nil, "", err
	}

	// ensure no overwrite
	if !overwrite {
		_, err := keybase.Get(keyName)
		if err == nil {
			return nil, "", errors.New("key already exists, overwrite is disabled")
		}
	}

	// generate a private key, with recovery phrase
	info, secret, err := keybase.Create(keyName, keyPass, keys.AlgoEd25519)
	if err != nil {
		return nil, "", err
	}
	addr := info.PubKey.Address()
	return addr, secret, nil
}
