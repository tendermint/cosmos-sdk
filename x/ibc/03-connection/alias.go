package connection

// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types"
)

const (
	AttributeKeyConnectionID         = types.AttributeKeyConnectionID
	AttributeKeyCounterpartyClientID = types.AttributeKeyCounterpartyClientID
	SubModuleName                    = types.SubModuleName
	StoreKey                         = types.StoreKey
	RouterKey                        = types.RouterKey
	QuerierRoute                     = types.QuerierRoute
	QueryAllConnections              = types.QueryAllConnections
	QueryAllClientConnections        = types.QueryAllClientConnections
	QueryClientConnections           = types.QueryClientConnections
	UNINITIALIZED                    = types.UNINITIALIZED
	INIT                             = types.INIT
	TRYOPEN                          = types.TRYOPEN
	OPEN                             = types.OPEN
)

var (
	// functions aliases
	NewKeeper                        = keeper.NewKeeper
	QuerierConnections               = keeper.QuerierConnections
	QuerierClientConnections         = keeper.QuerierClientConnections
	QuerierAllClientConnections      = keeper.QuerierAllClientConnections
	RegisterCodec                    = types.RegisterCodec
	RegisterInterfaces               = types.RegisterInterfaces
	NewConnectionEnd                 = types.NewConnectionEnd
	NewCounterparty                  = types.NewCounterparty
	ErrConnectionExists              = types.ErrConnectionExists
	ErrConnectionNotFound            = types.ErrConnectionNotFound
	ErrClientConnectionPathsNotFound = types.ErrClientConnectionPathsNotFound
	ErrConnectionPath                = types.ErrConnectionPath
	ErrInvalidConnectionState        = types.ErrInvalidConnectionState
	ErrInvalidCounterparty           = types.ErrInvalidCounterparty
	NewMsgConnectionOpenInit         = types.NewMsgConnectionOpenInit
	NewMsgConnectionOpenTry          = types.NewMsgConnectionOpenTry
	NewMsgConnectionOpenAck          = types.NewMsgConnectionOpenAck
	NewMsgConnectionOpenConfirm      = types.NewMsgConnectionOpenConfirm
	NewConnectionResponse            = types.NewConnectionResponse
	NewClientConnectionsResponse     = types.NewClientConnectionsResponse
	NewQueryClientConnectionsParams  = types.NewQueryClientConnectionsParams
	GetCompatibleVersions            = types.GetCompatibleVersions
	LatestVersion                    = types.LatestVersion
	PickVersion                      = types.PickVersion
	NewConnectionPaths               = types.NewConnectionPaths
	DefaultGenesisState              = types.DefaultGenesisState
	NewGenesisState                  = types.NewGenesisState

	// variable aliases
	SubModuleCdc                   = types.SubModuleCdc
	EventTypeConnectionOpenInit    = types.EventTypeConnectionOpenInit
	EventTypeConnectionOpenTry     = types.EventTypeConnectionOpenTry
	EventTypeConnectionOpenAck     = types.EventTypeConnectionOpenAck
	EventTypeConnectionOpenConfirm = types.EventTypeConnectionOpenConfirm
	AttributeValueCategory         = types.AttributeValueCategory
)

type (
	Keeper                       = keeper.Keeper
	End                          = types.ConnectionEnd
	State                        = types.State
	Counterparty                 = types.Counterparty
	ClientKeeper                 = types.ClientKeeper
	MsgConnectionOpenInit        = types.MsgConnectionOpenInit
	MsgConnectionOpenTry         = types.MsgConnectionOpenTry
	MsgConnectionOpenAck         = types.MsgConnectionOpenAck
	MsgConnectionOpenConfirm     = types.MsgConnectionOpenConfirm
	Response                     = types.ConnectionResponse
	ClientConnectionsResponse    = types.ClientConnectionsResponse
	QueryClientConnectionsParams = types.QueryClientConnectionsParams
	GenesisState                 = types.GenesisState
	Paths                        = types.ConnectionPaths
)
