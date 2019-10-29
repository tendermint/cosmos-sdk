package transfer

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types"
)

const (
	DefaultPacketTimeout    = keeper.DefaultPacketTimeout
	DefaultCodespace        = types.DefaultCodespace
	CodeInvalidAddress      = types.CodeInvalidAddress
	CodeErrSendPacket       = types.CodeErrSendPacket
	CodeInvalidPacketData   = types.CodeInvalidPacketData
	CodeInvalidChannelOrder = types.CodeInvalidChannelOrder
	CodeInvalidPort         = types.CodeInvalidPort
	CodeInvalidVersion      = types.CodeInvalidVersion
	AttributeKeyReceiver    = types.AttributeKeyReceiver
	SubModuleName           = types.SubModuleName
	StoreKey                = types.StoreKey
	RouterKey               = types.RouterKey
	QuerierRoute            = types.QuerierRoute
	BoundPortID             = types.BoundPortID
)

var (
	// functions aliases
	NewKeeper              = keeper.NewKeeper
	RegisterCodec          = types.RegisterCodec
	ErrInvalidAddress      = types.ErrInvalidAddress
	ErrSendPacket          = types.ErrSendPacket
	ErrInvalidPacketData   = types.ErrInvalidPacketData
	ErrInvalidChannelOrder = types.ErrInvalidChannelOrder
	ErrInvalidPort         = types.ErrInvalidPort
	ErrInvalidVersion      = types.ErrInvalidVersion
	GetEscrowAddress       = types.GetEscrowAddress
	GetDenomPrefix         = types.GetDenomPrefix
	GetModuleAccountName   = types.GetModuleAccountName
	NewMsgTransfer         = types.NewMsgTransfer

	// variable aliases
	ModuleCdc              = types.ModuleCdc
	AttributeValueCategory = types.AttributeValueCategory
)

type (
	Keeper                  = keeper.Keeper
	BankKeeper              = types.BankKeeper
	ChannelKeeper           = types.ChannelKeeper
	ClientKeeper            = types.ClientKeeper
	ConnectionKeeper        = types.ConnectionKeeper
	SupplyKeeper            = types.SupplyKeeper
	MsgTransfer             = types.MsgTransfer
	PacketData      = types.PacketData
	PacketDataAlias = types.PacketDataAlias
)
