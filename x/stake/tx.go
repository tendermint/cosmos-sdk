package stake

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/go-crypto"
)

//Verify interface at compile time
var _, _, _, _ sdk.Msg = &MsgDeclareCandidacy{}, &MsgEditCandidacy{}, &MsgDelegate{}, &MsgUnbond{}

//______________________________________________________________________

// MsgAddr - struct for bonding or unbonding transactions
type MsgAddr struct {
	Address sdk.Address `json:"address"`
}

func NewMsgAddr(address sdk.Address) MsgAddr {
	return MsgAddr{
		Address: address,
	}
}

// nolint
func (msg MsgAddr) Type() string                            { return "stake" }
func (msg MsgAddr) Get(key interface{}) (value interface{}) { return nil }
func (msg MsgAddr) GetSigners() []sdk.Address               { return []sdk.Address{msg.Address} }
func (msg MsgAddr) String() string {
	return fmt.Sprintf("MsgAddr{Address: %v}", msg.Address)
}

// ValidateBasic - Check for non-empty candidate, and valid coins
func (msg MsgAddr) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return errCandidateEmpty
	}
}

//______________________________________________________________________

// MsgDeclareCandidacy - struct for unbonding transactions
type MsgDeclareCandidacy struct {
	MsgAddr
	Description
	Bond   sdk.Coin      `json:"bond"`
	PubKey crypto.PubKey `json:"pubkey"`
}

func NewMsgDeclareCandidacy(bond sdk.Coin, address sdk.Address, pubkey crypto.PubKey, description Description) MsgDeclareCandidacy {
	return MsgDeclareCandidacy{
		MsgAddr:     NewMsgAddr(address),
		Description: description,
		Bond:        bond,
		PubKey:      PubKey,
	}
}

// get the bytes for the message signer to sign on
func (msg MsgDeclareCandidacy) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// quick validity check
func (msg MsgDeclareCandidacy) ValidateBasic() sdk.Error {
	err := MsgAddr.ValidateBasic()
	if err != nil {
		return err
	}
	err := validateCoin(msg.Bond)
	if err != nil {
		return err
	}
	empty := Description{}
	if msg.Description == empty {
		return fmt.Errorf("description must be included")
	}
	return nil
}

//______________________________________________________________________

// MsgEditCandidacy - struct for editing a candidate
type MsgEditCandidacy struct {
	MsgAddr
	Description
}

func NewMsgEditCandidacy(address sdk.Address, description Description) MsgEditCandidacy {
	return MsgEditCandidacy{
		MsgAddr:     NewMsgAddr(address),
		Description: description,
	}
}

// get the bytes for the message signer to sign on
func (msg MsgEditCandidacy) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// quick validity check
func (msg MsgEditCandidacy) ValidateBasic() sdk.Error {
	err := MsgAddr.ValidateBasic()
	if err != nil {
		return err
	}
	empty := Description{}
	if msg.Description == empty {
		return fmt.Errorf("Transaction must include some information to modify")
	}
	return nil
}

//______________________________________________________________________

// MsgDelegate - struct for bonding transactions
type MsgDelegate struct {
	MsgAddr
	Bond sdk.Coin `json:"bond"`
}

func NewMsgDelegate(address sdk.Address, bond sdk.Coin) MsgDelegate {
	return MsgDelegate{
		MsgAddr: NewMsgAddr(address),
		Bond:    bond,
	}
}

// get the bytes for the message signer to sign on
func (msg MsgDelegate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// quick validity check
func (msg MsgDelegate) ValidateBasic() sdk.Error {
	err := MsgAddr.ValidateBasic()
	if err != nil {
		return err
	}
	err := validateCoin(msg.Bond)
	if err != nil {
		return err
	}
	return nil
}

//______________________________________________________________________

// MsgUnbond - struct for unbonding transactions
type MsgUnbond struct {
	MsgAddr
	Shares string `json:"shares"`
}

func NewMsgUnbond(shares string, address sdk.Address) MsgDelegate {
	return MsgUnbond{
		MsgAddr: NewMsgAddr(address),
		Shares:  shares,
	}
}

// get the bytes for the message signer to sign on
func (msg MsgUnbond) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// quick validity check
func (msg MsgUnbond) ValidateBasic() sdk.Error {
	err := MsgAddr.ValidateBasic()
	if err != nil {
		return err
	}
	if msg.Shares {
		return ErrCandidateEmpty()
	}
	return nil
}

//______________________________________________________________________
// helper

func validateCoin(coin coin.Coin) sdk.Error {
	coins := sdk.Coins{bond}
	if !sdk.IsValid() {
		return sdk.ErrInvalidCoins()
	}
	if !coins.IsPositive() {
		return fmt.Errorf("Amount must be > 0")
	}
	return nil
}
