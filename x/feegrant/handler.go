package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/feegrant/exported"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgGrantFeeAllowanceBase:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%T must be extended to support feegrant", msg)

		case MsgRevokeFeeAllowance:
			return handleRevokeFee(ctx, k, msg)

		default:
			msgGrantFa, ok := msg.(exported.MsgGrantFeeAllowance)
			if ok {
				return handleGrantFee(ctx, k, msgGrantFa)
			}
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleGrantFee(ctx sdk.Context, k Keeper, msg exported.MsgGrantFeeAllowance) (*sdk.Result, error) {
	grant := exported.FeeAllowanceGrant(msg)
	k.GrantFeeAllowance(ctx, grant)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleRevokeFee(ctx sdk.Context, k Keeper, msg MsgRevokeFeeAllowance) (*sdk.Result, error) {
	k.RevokeFeeAllowance(ctx, msg.Granter, msg.Grantee)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
