package types

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pool - dynamic parameters of the current state
type Pool struct {
	LooseTokens       sdk.Dec `json:"loose_tokens"`        // tokens which are not bonded in a validator
	BondedTokens      sdk.Dec `json:"bonded_tokens"`       // reserve of bonded tokens
	InflationLastTime int64   `json:"inflation_last_time"` // block which the last inflation was processed // TODO make time
	Inflation         sdk.Dec `json:"inflation"`           // current annual inflation rate

	DateLastCommissionReset int64 `json:"date_last_commission_reset"` // unix timestamp for last commission accounting reset (daily)

	// Fee Related
	PrevBondedShares sdk.Dec `json:"prev_bonded_shares"` // last recorded bonded shares - for fee calculations
}

// nolint
func (p Pool) Equal(p2 Pool) bool {
	bz1 := MsgCdc.MustMarshalBinary(&p)
	bz2 := MsgCdc.MustMarshalBinary(&p2)
	return bytes.Equal(bz1, bz2)
}

// initial pool for testing
func InitialPool() Pool {
	return Pool{
		LooseTokens:             sdk.ZeroDec(),
		BondedTokens:            sdk.ZeroDec(),
		InflationLastTime:       0,
		Inflation:               sdk.NewDec(7, 2),
		DateLastCommissionReset: 0,
		PrevBondedShares:        sdk.ZeroDec(),
	}
}

//____________________________________________________________________

// Sum total of all staking tokens in the pool
func (p Pool) TokenSupply() sdk.Dec {
	return p.LooseTokens.Add(p.BondedTokens)
}

//____________________________________________________________________

// get the bond ratio of the global state
func (p Pool) BondedRatio() sdk.Dec {
	supply := p.TokenSupply()
	if supply.GT(sdk.ZeroDec()) {
		return p.BondedTokens.Quo(supply)
	}
	return sdk.ZeroDec()
}

//_______________________________________________________________________

func (p Pool) looseTokensToBonded(bondedTokens sdk.Dec) Pool {
	p.BondedTokens = p.BondedTokens.Add(bondedTokens)
	p.LooseTokens = p.LooseTokens.Sub(bondedTokens)
	if p.LooseTokens.LT(sdk.ZeroDec()) {
		panic(fmt.Sprintf("sanity check: loose tokens negative, pool: %v", p))
	}
	return p
}

func (p Pool) bondedTokensToLoose(bondedTokens sdk.Dec) Pool {
	p.BondedTokens = p.BondedTokens.Sub(bondedTokens)
	p.LooseTokens = p.LooseTokens.Add(bondedTokens)
	if p.BondedTokens.LT(sdk.ZeroDec()) {
		panic(fmt.Sprintf("sanity check: bonded tokens negative, pool: %v", p))
	}
	return p
}

//_______________________________________________________________________
// Inflation

const precision = 10000               // increased to this precision for accuracy
var hrsPerYrDec = sdk.NewDec(8766, 0) // as defined by a julian year of 365.25 days

// process provisions for an hour period
func (p Pool) ProcessProvisions(params Params) Pool {
	p.Inflation = p.NextInflation(params)
	provisions := p.Inflation.
		Mul(p.TokenSupply()).
		Quo(hrsPerYrDec)

	// TODO add to the fees provisions
	p.LooseTokens = p.LooseTokens.Add(provisions)
	return p
}

// get the next inflation rate for the hour
func (p Pool) NextInflation(params Params) (inflation sdk.Dec) {

	// The target annual inflation rate is recalculated for each previsions cycle. The
	// inflation is also subject to a rate change (positive or negative) depending on
	// the distance from the desired ratio (67%). The maximum rate change possible is
	// defined to be 13% per year, however the annual inflation is capped as between
	// 7% and 20%.

	// (1 - bondedRatio/GoalBonded) * InflationDeceChange
	inflationDeceChangePerYear := sdk.OneDec().
		Sub(p.BondedRatio().
			Quo(params.GoalBonded)).
		Mul(params.InflationDeceChange)
	inflationDeceChange := inflationDeceChangePerYear.Quo(hrsPerYrDec)

	// increase the new annual inflation for this next cycle
	inflation = p.Inflation.Add(inflationDeceChange)
	if inflation.GT(params.InflationMax) {
		inflation = params.InflationMax
	}
	if inflation.LT(params.InflationMin) {
		inflation = params.InflationMin
	}

	return inflation
}
