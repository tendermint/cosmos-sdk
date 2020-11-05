package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/server/rosetta"
	"github.com/cosmos/cosmos-sdk/server/rosetta/cosmos/client"
	"github.com/cosmos/cosmos-sdk/server/rosetta/services"
	"github.com/ghodss/yaml"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// configuration defaults constants
const (
	// DefaultConfigBlockchain defines the default blockchain identifier name
	// TODO: should it be cosmos stargate?
	DefaultConfigBlockchain = "cosmos"
	// DefaultConfigAddr defines the default rosetta binding address
	DefaultConfigAddr = ":8080"
	// DefaultConfigRetries is the default number of retries
	DefaultConfigRetries = 5
)

// configuration flags
const (
	flagBlockchain         = "blockchain"
	flagNetwork            = "network"
	flagTendermintEndpoint = "tendermint"
	flagGRPCEndpoint       = "grpc"
	flagAddr               = "addr"
	flagRetries            = "retries"
	flagFile               = "file"
	flagOffline            = "offline"
)

// RosettaFromConfig builds the rosetta servicer full implementation from configurations
func RosettaFromConfig(conf *Config) (rosetta.Rosetta, error) {
	if conf.Offline {
		panic("offline mode not supported for now")
	}
	dataAPIClient, err := client.NewSingle(conf.GRPCEndpoint, conf.TendermintRPC)
	if err != nil {
		return nil, fmt.Errorf("data api client init failure: %w", err)
	}
	sn, err := services.NewSingleNetwork(dataAPIClient, &types.NetworkIdentifier{
		Blockchain: conf.Blockchain,
		Network:    conf.Network,
	})
	if err != nil {
		return nil, fmt.Errorf("rosetta network initialization failure: %w", err)
	}
	return sn, nil
}

// RetryRosettaFromConfig tries to initialize rosetta retrying
func RetryRosettaFromConfig(conf *Config) (rosetta rosetta.Rosetta, err error) {
	for i := 0; i < conf.Retries; i++ {
		rosetta, err = RosettaFromConfig(conf)
		if err == nil {
			return
		}
		time.Sleep(5 * time.Second)
	}
	return
}

// Config defines the configuration of the rosetta server
type Config struct {
	// Blockchain defines the blockchain name
	// defaults to DefaultConfigBlockchain
	Blockchain string `json:"blockchain" yaml:"blockchain" env:"ROSETTA_BLOCKCHAIN"`
	// Network defines the network name
	Network string `json:"network" yaml:"network" env:"ROSETTA_NETWORK"`
	// TendermintRPC defines the endpoint to connect to
	// tendermint RPC, specifying 'tcp://' before is not
	// required, usually it's at port 26657 of the
	TendermintRPC string `json:"tendermint_rpc" yaml:"tendermintRPC" env:"ROSETTA_TENDERMINT_RPC"`
	// GRPCEndpoint defines the cosmos application gRPC endpoint
	// usually it is located at 9090 port
	GRPCEndpoint string `json:"grpc_endpoint" yaml:"gRPCEndpoint" env:"ROSETTA_GRPC_ENDPOINT"`
	// Addr defines the default address to bind the rosetta server to
	// defaults to DefaultConfigAddr
	Addr string `json:"addr" yaml:"addr" env:"ROSETTA_ADDR"`
	// Retries defines the maximum number of retries
	// rosetta will do before quitting
	Retries int `json:"retries" yaml:"retries" env:"ROSETTA_RETRIES"`
	// Offline defines if the server must be run in offline mode
	Offline bool `json:"offline" yaml:"offline" env:"ROSETTA_OFFLINE"`
}

// Validate validates a configuration and sets
// its defaults in case they were not provided
func (c *Config) Validate() error {
	// set defaults
	if c.Addr == "" {
		c.Addr = DefaultConfigAddr
	}
	if c.Blockchain == "" {
		c.Blockchain = DefaultConfigBlockchain
	}
	if c.Retries == 0 {
		c.Retries = DefaultConfigRetries
	}
	// these are must
	if c.Network == "" {
		return fmt.Errorf("network not provided")
	}
	if c.Offline {
		return nil
	}
	// these are optional but it must be online
	if c.GRPCEndpoint == "" {
		return fmt.Errorf("grpc endpoint not provided")
	}
	if c.TendermintRPC == "" {
		return fmt.Errorf("tendermint rpc not provided")
	}
	if !strings.HasPrefix(c.TendermintRPC, "tcp://") {
		c.TendermintRPC = fmt.Sprintf("tcp://%s", c.TendermintRPC)
	}
	return nil
}

// ConfigFromEnv tries to get the configurations from the environment variable
func ConfigFromEnv() (*Config, error) {
	conf := &Config{}
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	err = conf.Validate()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// ConfigFromYaml attempts to get a configuration given a yaml file
func ConfigFromYaml(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	err = yaml.Unmarshal(b, conf)
	if err != nil {
		return nil, err
	}
	err = conf.Validate()
	if err != nil {
		return conf, err
	}
	return conf, nil
}

// ConfigFromCmdFlags gets the configuration from flags
func ConfigFromCmdFlags(flags *pflag.FlagSet) (*Config, error) {
	blockchain, err := flags.GetString(flagBlockchain)
	if err != nil {
		return nil, err
	}
	network, err := flags.GetString(flagNetwork)
	if err != nil {
		return nil, err
	}
	tendermintRPC, err := flags.GetString(flagTendermintEndpoint)
	if err != nil {
		return nil, err
	}
	gRPCEndpoint, err := flags.GetString(flagGRPCEndpoint)
	if err != nil {
		return nil, err
	}
	addr, err := flags.GetString(flagAddr)
	if err != nil {
		return nil, err
	}
	retries, err := flags.GetInt(flagRetries)
	if err != nil {
		return nil, err
	}
	offline, err := flags.GetBool(flagOffline)
	conf := &Config{
		Blockchain:    blockchain,
		Network:       network,
		TendermintRPC: tendermintRPC,
		GRPCEndpoint:  gRPCEndpoint,
		Addr:          addr,
		Retries:       retries,
		Offline:       offline,
	}
	err = conf.Validate()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// SetConfigFlags sets the configuration flags to the given flagset
func SetConfigFlags(flags *pflag.FlagSet) {
	flags.StringP(flagFile, "f", "", "the .yaml configuration file (optional, can use env or flags)")
	flags.String(flagBlockchain, DefaultConfigBlockchain, "the blockchain type")
	flags.String(flagNetwork, "", "the network name")
	flags.String(flagTendermintEndpoint, "", "the tendermint rpc endpoint, without tcp://")
	flags.String(flagGRPCEndpoint, "", "the app gRPC endpoint")
	flags.String(flagAddr, DefaultConfigAddr, "the address rosetta will bind to")
	flags.Int(flagRetries, DefaultConfigRetries, "the number of retries that will be done before quitting")
	return
}

// FindConfigs will attempt to find configurations
// giving priority to
// 1) if config is set via flags
// 2) flags
// 3) environment variables
func FindConfig(flags *pflag.FlagSet) (*Config, error) {
	// try config file
	filePath, err := flags.GetString(flagFile)
	if err == nil && filePath != "" {
		return ConfigFromYaml(filePath)
	}
	// try flags
	config, err := ConfigFromCmdFlags(flags)
	if err == nil {
		return config, nil
	}
	// try env
	config, err = ConfigFromEnv()
	if err == nil {
		return config, nil
	}
	return nil, fmt.Errorf("unable to find valid configurations")
}

// MustFindConfig is used to find configs but if it fails it panics
func MustFindConfig(flags *pflag.FlagSet) *Config {
	config, err := FindConfig(flags)
	if err != nil {
		panic(err)
	}
	return config
}
