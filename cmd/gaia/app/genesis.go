package app

import (
	"encoding/json"
	"errors"
	"github.com/spf13/pflag"
	crypto "github.com/tendermint/go-crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/server"
	gc "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/stake"
)

// State to Unmarshal
type GenesisState struct {
	Accounts  []GenesisAccount   `json:"accounts"`
	StakeData stake.GenesisState `json:"stake"`
}

// GenesisAccount doesn't need pubkey or sequence
type GenesisAccount struct {
	Address sdk.Address `json:"address"`
	Coins   sdk.Coins   `json:"coins"`
}

func NewGenesisAccount(acc *auth.BaseAccount) GenesisAccount {
	return GenesisAccount{
		Address: acc.Address,
		Coins:   acc.Coins,
	}
}

func NewGenesisAccountI(acc auth.Account) GenesisAccount {
	return GenesisAccount{
		Address: acc.GetAddress(),
		Coins:   acc.GetCoins(),
	}
}

// convert GenesisAccount to auth.BaseAccount
func (ga *GenesisAccount) ToAccount() (acc *auth.BaseAccount) {
	return &auth.BaseAccount{
		Address: ga.Address,
		Coins:   ga.Coins.Sort(),
	}
}

// get app init parameters for server init command
func GaiaAppInit() server.AppInit {
	fsAppGenState := pflag.NewFlagSet("", pflag.ContinueOnError)

	fsAppGenTx := pflag.NewFlagSet("", pflag.ContinueOnError)
	fsAppGenTx.String(server.FlagName, "", "validator moniker, required")
	fsAppGenTx.String(server.FlagClientHome, DefaultCLIHome,
		"home directory for the client, used for key generation")
	fsAppGenTx.Bool(server.FlagOWK, false, "overwrite the accounts created")
<<<<<<< HEAD

	flagNames := server.GenTxFlagNames{
		FlagName:       flagName,
		FlagClientHome: flagClientHome,
		FlagOWK:        flagOWK,
	}
=======
>>>>>>> Fixes requested by Rigel

	return server.AppInit{
		FlagsAppGenState: fsAppGenState,
		FlagsAppGenTx:    fsAppGenTx,
		AppGenTx:         GaiaAppGenTx,
		AppGenState:      GaiaAppGenStateJSON,
	}
}

// simple genesis tx
type GaiaGenTx struct {
	Name    string        `json:"name"`
	Address sdk.Address   `json:"address"`
	PubKey  crypto.PubKey `json:"pub_key"`
}

// Generate a gaia genesis transaction with flags
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func GaiaAppGenTx(cdc *wire.Codec, pk crypto.PubKey, genTxConfig gc.GenTxConfig) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {
	if genTxConfig.Name == "" {
=======
func GaiaAppGenTx(cdc *wire.Codec, pk crypto.PubKey, config *gaiacfg.Config) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {
	if config == nil {
		config = gaiacfg.DefaultConfig()
		config.Name = viper.GetString(flagName)
		config.CliRoot = viper.GetString(flagClientHome)
		config.GenTx.Overwrite = viper.GetBool(flagOWK)
	}
	if config.Name == "" {
>>>>>>> Added testnet command
=======
func GaiaAppGenTx(cdc *wire.Codec, pk crypto.PubKey, genTxConfig server.GenTxConfig) (
=======
func GaiaAppGenTx(cdc *wire.Codec, pk crypto.PubKey, genTxConfig gc.GenTxConfig) (
>>>>>>> Adding server/config to app package
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {
	if genTxConfig.Name == "" {
>>>>>>> Finished testnet command and introduced localnet targets in Makefile, together with gaiadnode Docker image
		return nil, nil, tmtypes.GenesisValidator{}, errors.New("Must specify --name (validator moniker)")
	}

	var addr sdk.Address
	var secret string
<<<<<<< HEAD
<<<<<<< HEAD
	addr, secret, err = server.GenerateSaveCoinKey(genTxConfig.CliRoot, genTxConfig.Name, "1234567890", genTxConfig.Overwrite)
=======
	addr, secret, err = server.GenerateSaveCoinKey(config.CliRoot, config.Name, "1234567890", config.GenTx.Overwrite)
>>>>>>> Added testnet command
=======
	addr, secret, err = server.GenerateSaveCoinKey(genTxConfig.CliRoot, genTxConfig.Name, "1234567890", genTxConfig.Overwrite)
>>>>>>> Finished testnet command and introduced localnet targets in Makefile, together with gaiadnode Docker image
	if err != nil {
		return
	}
	mm := map[string]string{"secret": secret}
	var bz []byte
	bz, err = cdc.MarshalJSON(mm)
	if err != nil {
		return
	}

	cliPrint = json.RawMessage(bz)
<<<<<<< HEAD
<<<<<<< HEAD
	appGenTx, _, validator, err = GaiaAppGenTxNF(cdc, pk, addr, genTxConfig.Name, genTxConfig.Overwrite)
=======
	appGenTx, _, validator, err = GaiaAppGenTxNF(cdc, pk, addr, config.Name, config.GenTx.Overwrite)
>>>>>>> Added testnet command
=======
	appGenTx, _, validator, err = GaiaAppGenTxNF(cdc, pk, addr, genTxConfig.Name, genTxConfig.Overwrite)
>>>>>>> Finished testnet command and introduced localnet targets in Makefile, together with gaiadnode Docker image
	return
}

// Generate a gaia genesis transaction without flags
func GaiaAppGenTxNF(cdc *wire.Codec, pk crypto.PubKey, addr sdk.Address, name string, overwrite bool) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {

	var bz []byte
	gaiaGenTx := GaiaGenTx{
		Name:    name,
		Address: addr,
		PubKey:  pk,
	}
	bz, err = wire.MarshalJSONIndent(cdc, gaiaGenTx)
	if err != nil {
		return
	}
	appGenTx = json.RawMessage(bz)

	validator = tmtypes.GenesisValidator{
		PubKey: pk,
		Power:  server.FreeFermionVal,
	}
	return
}

// Create the core parameters for genesis initialization for gaia
// note that the pubkey input is this machines pubkey
func GaiaAppGenState(cdc *wire.Codec, appGenTxs []json.RawMessage) (genesisState GenesisState, err error) {

	if len(appGenTxs) == 0 {
		err = errors.New("must provide at least genesis transaction")
		return
	}

	// start with the default staking genesis state
	stakeData := stake.DefaultGenesisState()

	// get genesis flag account information
	genaccs := make([]GenesisAccount, len(appGenTxs))
	for i, appGenTx := range appGenTxs {

		var genTx GaiaGenTx
		err = cdc.UnmarshalJSON(appGenTx, &genTx)
		if err != nil {
			return
		}

		// create the genesis account, give'm few steaks and a buncha token with there name
		accAuth := auth.NewBaseAccountWithAddress(genTx.Address)
		accAuth.Coins = sdk.Coins{
			{genTx.Name + "Token", 1000},
			{"steak", server.FreeFermionsAcc},
		}
		acc := NewGenesisAccount(&accAuth)
		genaccs[i] = acc
		stakeData.Pool.LooseUnbondedTokens += server.FreeFermionsAcc // increase the supply

		// add the validator
		if len(genTx.Name) > 0 {
			desc := stake.NewDescription(genTx.Name, "", "", "")
			validator := stake.NewValidator(genTx.Address, genTx.PubKey, desc)
			validator.PoolShares = stake.NewBondedShares(sdk.NewRat(server.FreeFermionVal))
			stakeData.Validators = append(stakeData.Validators, validator)

			// pool logic
			stakeData.Pool.BondedTokens += server.FreeFermionVal
			stakeData.Pool.BondedShares = sdk.NewRat(stakeData.Pool.BondedTokens)
		}
	}

	// create the final app state
	genesisState = GenesisState{
		Accounts:  genaccs,
		StakeData: stakeData,
	}
	return
}

// GaiaAppGenState but with JSON
func GaiaAppGenStateJSON(cdc *wire.Codec, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	// create the final app state
	genesisState, err := GaiaAppGenState(cdc, appGenTxs)
	if err != nil {
		return nil, err
	}
	appState, err = wire.MarshalJSONIndent(cdc, genesisState)
	return
}
