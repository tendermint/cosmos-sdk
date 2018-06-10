package config

//_____________________________________________________________________

// Configuration structure for command functions that share configuration.
// For example: init, init gen-tx and testnet commands need similar input and run the same code

<<<<<<< HEAD
// Storage for init gen-tx command input parameters
=======
>>>>>>> Separated GenTxConfig into a server/config package so both the server package and the mock package can use it
type GenTxConfig struct {
	Name      string
	CliRoot   string
	Overwrite bool
	IP        string
}
