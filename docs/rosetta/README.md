<!--
order: false
parent:
  order: 2
-->

# Rosetta

This folder contains instructions about how to use the Coinbase Rosetta API integration.

## Motivations and design

For the motivation an design choices refer to the [ADR 035](../architecture/adr-035-rosetta-api-support.md)

## Usage

The Rosetta API server is a stand-alone server that connects to a node of a chain developed with the Cosmos-SDK. 

To enable the Rosetta API server support it's required to Add the `RosettaCommand` to your application root command file.
After the 

```
initRootCmd(rootCmd, encodingConfig)
```

line, add the following lines:

```
rootCmd.AddCommand(
  server.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Marshaler)
)
```


The application root command file is in the server package in `github.com/cosmos/cosmos-sdk/server`.

Since we’ve updated the Cosmos SDK to work with the Rosetta API, updating the application root command file is all you need to do.

To run Rosetta in your application CLI, use the following command:
`appd rosetta --help`.
To test and run Rosetta API endpoints for applications that are running and exposed, use the following command:

```
appd rosetta
     --blockchain "your application name (ex: gaia)"
     --network "your chain identifier (ex: testnet-1)"
     --tendermint "tendermint endpoint (ex: localhost:26657)"
     --grpc "gRPC endpoint (ex: localhost:9090)"
     --addr "rosetta binding address (ex: :8080)"
```
