package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
)

var _ FeeAllowanceI = (*FilteredFeeAllowance)(nil)
var _ types.UnpackInterfacesMessage = (*FilteredFeeAllowance)(nil)

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (a *FilteredFeeAllowance) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var allowance FeeAllowanceI
	return unpacker.UnpackAny(a.Allowance, &allowance)
}

// NewFilteredFeeAllowance creates new filtered fee allowance.
func NewFilteredFeeAllowance(allowance FeeAllowanceI, allowedMsgs []string) (*FilteredFeeAllowance, error) {
	msg, ok := allowance.(proto.Message)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	return &FilteredFeeAllowance{
		Allowance:       any,
		AllowedMessages: allowedMsgs,
	}, nil
}

// GetAllowance returns allowed fee allowance.
func (a *FilteredFeeAllowance) GetAllowance() (FeeAllowanceI, error) {
	allowance, ok := a.Allowance.GetCachedValue().(FeeAllowanceI)
	if !ok {
		return nil, sdkerrors.Wrap(ErrNoAllowance, "failed to get allowance")
	}

	return allowance, nil
}

// Accept method checks for the filtered messages has valid expiry
func (a *FilteredFeeAllowance) Accept(fee sdk.Coins, blockTime time.Time, blockHeight int64, msgs []sdk.Msg) (bool, error) {
	if !a.isMsgTypesAllowed(msgs) {
		return false, sdkerrors.Wrap(ErrMessageNotAllowed, "message does not exist in allowed messages")
	}

	allowance, err := a.GetAllowance()
	if err != nil {
		return false, err
	}

	return allowance.Accept(fee, blockTime, blockHeight, msgs)
}

func (a *FilteredFeeAllowance) isMsgTypesAllowed(msgs []sdk.Msg) bool {
	found := false

	for _, msg := range msgs {
		for _, allowedMsg := range a.AllowedMessages {
			if allowedMsg == msg.Type() {
				found = true
				break
			}
		}

		if !found {
			return false
		}

		found = false
	}

	return true
}

// PrepareForExport will adjust the expiration based on export time. In particular,
// it will subtract the dumpHeight from any height-based expiration to ensure that
// the elapsed number of blocks this allowance is valid for is fixed.
func (a *FilteredFeeAllowance) PrepareForExport(dumpTime time.Time, dumpHeight int64) FeeAllowanceI {
	allowance, err := a.GetAllowance()
	if err != nil {
		panic("failed to get allowance")
	}

	f, err := NewFilteredFeeAllowance(allowance.PrepareForExport(dumpTime, dumpHeight), a.AllowedMessages)
	if err != nil {
		panic("failed to export filtered fee allowance")
	}

	return f
}

// ValidateBasic implements FeeAllowance and enforces basic sanity checks
func (a *FilteredFeeAllowance) ValidateBasic() error {
	if a.Allowance == nil {
		return sdkerrors.Wrap(ErrNoAllowance, "allowance should not be empty")
	}
	if len(a.AllowedMessages) == 0 {
		return sdkerrors.Wrap(ErrNoMessages, "allowed messages shouldn't be empty")
	}

	allowance, err := a.GetAllowance()
	if err != nil {
		return err
	}

	return allowance.ValidateBasic()
}
