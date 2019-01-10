package gov

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote
type Vote struct {
	Voter      sdk.AccAddress `json:"voter"`       //  address of the voter
	ProposalID uint64         `json:"proposal_id"` //  proposalID of the proposal
	Option     VoteOption     `json:"option"`      //  option from OptionSet chosen by the voter
}

// HumanReadableString implements client.Printable
func (v Vote) HumanReadableString() string {
	return fmt.Sprintf("%s voted with option %s for proposal %d...", v.Voter, v.Option, v.ProposalID)
}

// Votes is an array of vote
type Votes []Vote

// HumanReadableString implements client.Printable
func (v Votes) HumanReadableString() (out string) {
	if len(v) < 1 {
		return ""
	}
	out += fmt.Sprintf("Votes for Proposal %d:\n", v[0].ProposalID)
	for _, vot := range v {
		out += fmt.Sprintf("  %s: %s\n", vot.Voter, vot.Option)
	}
	return strings.TrimSpace(out)
}

// Returns whether 2 votes are equal
func (v Vote) Equals(comp Vote) bool {
	return v.Voter.Equals(comp.Voter) && v.ProposalID == comp.ProposalID && v.Option == comp.Option
}

// Returns whether a vote is empty
func (v Vote) Empty() bool {
	return v.Equals(Vote{})
}

// Deposit
type Deposit struct {
	Depositor  sdk.AccAddress `json:"depositor"`   //  Address of the depositor
	ProposalID uint64         `json:"proposal_id"` //  proposalID of the proposal
	Amount     sdk.Coins      `json:"amount"`      //  Deposit amount
}

// HumanReadableString implements client.Printable
func (d Deposit) HumanReadableString() string {
	return fmt.Sprintf("%s deposited %s on proposal %d...", d.Depositor, d.Amount, d.ProposalID)
}

// Deposits is a collection of deposit
type Deposits []Deposit

// HumanReadableString implements client.Printable
func (d Deposits) HumanReadableString() (out string) {
	if len(d) < 1 {
		return ""
	}
	out += fmt.Sprintf("Deposits for Proposal %d:\n", d[0].ProposalID)
	for _, dep := range d {
		out += fmt.Sprintf("  %s: %s\n", dep.Depositor, dep.Amount)
	}
	return strings.TrimSpace(out)
}

// Returns whether 2 deposits are equal
func (d Deposit) Equals(comp Deposit) bool {
	return d.Depositor.Equals(comp.Depositor) && d.ProposalID == comp.ProposalID && d.Amount.IsEqual(comp.Amount)
}

// Returns whether a deposit is empty
func (d Deposit) Empty() bool {
	return d.Equals(Deposit{})
}

// Type that represents VoteOption as a byte
type VoteOption byte

//nolint
const (
	OptionEmpty      VoteOption = 0x00
	OptionYes        VoteOption = 0x01
	OptionAbstain    VoteOption = 0x02
	OptionNo         VoteOption = 0x03
	OptionNoWithVeto VoteOption = 0x04
)

// String to proposalType byte.  Returns ff if invalid.
func VoteOptionFromString(str string) (VoteOption, error) {
	switch str {
	case "Yes":
		return OptionYes, nil
	case "Abstain":
		return OptionAbstain, nil
	case "No":
		return OptionNo, nil
	case "NoWithVeto":
		return OptionNoWithVeto, nil
	default:
		return VoteOption(0xff), errors.Errorf("'%s' is not a valid vote option", str)
	}
}

// Is defined VoteOption
func validVoteOption(option VoteOption) bool {
	if option == OptionYes ||
		option == OptionAbstain ||
		option == OptionNo ||
		option == OptionNoWithVeto {
		return true
	}
	return false
}

// Marshal needed for protobuf compatibility
func (vo VoteOption) Marshal() ([]byte, error) {
	return []byte{byte(vo)}, nil
}

// Unmarshal needed for protobuf compatibility
func (vo *VoteOption) Unmarshal(data []byte) error {
	*vo = VoteOption(data[0])
	return nil
}

// Marshals to JSON using string
func (vo VoteOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(vo.String())
}

// Unmarshals from JSON assuming Bech32 encoding
func (vo *VoteOption) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return nil
	}

	bz2, err := VoteOptionFromString(s)
	if err != nil {
		return err
	}
	*vo = bz2
	return nil
}

// Turns VoteOption byte to String
func (vo VoteOption) String() string {
	switch vo {
	case OptionYes:
		return "Yes"
	case OptionAbstain:
		return "Abstain"
	case OptionNo:
		return "No"
	case OptionNoWithVeto:
		return "NoWithVeto"
	default:
		return ""
	}
}

// For Printf / Sprintf, returns bech32 when using %s
// nolint: errcheck
func (vo VoteOption) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(fmt.Sprintf("%s", vo.String())))
	default:
		s.Write([]byte(fmt.Sprintf("%v", byte(vo))))
	}
}
