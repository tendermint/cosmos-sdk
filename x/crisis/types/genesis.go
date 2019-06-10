package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - crisis genesis state
type GenesisState struct {
	ConstantFee sdk.Coin `json:"constant_fee"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(constantFee sdk.Coin) GenesisState {
	return GenesisState{
		ConstantFee: constantFee,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		ConstantFee: sdk.NewCoin(sdk.BondDenom, sdk.NewInt(1000)),
	}
}
