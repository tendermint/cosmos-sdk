package connection

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/store/mapping"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client"
)

// XXX: all panic -> err
// XXX: Try -> TryOpen

type Manager struct {
	protocol mapping.Mapping

	client client.Manager

	// CONTRACT: remote/self should not be used when remote
	remote *Manager
	self   mapping.Indexer
}

func NewManager(protocol, free mapping.Base, client client.Manager) Manager {
	return Manager{
		protocol: mapping.NewMapping(protocol, []byte("/")),

		client: client,

		self: mapping.NewIndexer(free, []byte("/self"), mapping.Dec),
	}
}

// TODO: return newtyped manager
func NewRemoteManager(protocol mapping.Base, client client.Manager) Manager {
	return NewManager(protocol, mapping.EmptyBase(), client)
}

func (man Manager) Exec(remote Manager, fn func(Manager)) {
	fn(Manager{
		protocol: man.protocol,
		client:   man.client,
		self:     man.self,
		remote:   &remote,
	})
}

// CONTRACT: client and remote must be filled by the caller
func (man Manager) object(id string) Object {
	return Object{
		id:          id,
		connection:  man.protocol.Value([]byte(id)),
		state:       man.protocol.Value([]byte(id + "/state")).Enum(),
		nexttimeout: man.protocol.Value([]byte(id + "/timeout")).Integer(),

		self: man.self,
	}
}

// Init Try
func (man Manager) Create(ctx sdk.Context, id string, client client.Object, remote Object) (Object, error) {
	obj := man.object(id)
	if obj.exists(ctx) {
		return Object{}, errors.New("connection already exists for the provided id")
	}
	obj.client = client
	obj.remote = &remote
	return obj, nil
}

// Ack Confirm
func (man Manager) Query(ctx sdk.Context, key string) (obj Object, err error) {
	obj = man.object(key)
	if !obj.exists(ctx) {
		return Object{}, errors.New("connection not exists for the provided id")
	}
	conn := obj.Value(ctx)
	obj.client, err = man.client.Query(ctx, conn.Client)
	if err != nil {
		return
	}
	remote := man.remote.object(conn.Counterparty)
	obj.remote = &remote
	return
}

type Object struct {
	id          string
	connection  mapping.Value
	state       mapping.Enum
	nexttimeout mapping.Integer

	client client.Object

	// CONTRACT: remote/self should not be used when remote
	remote *Object
	self   mapping.Indexer
}

func (obj Object) create(ctx sdk.Context, c Connection) error {
	if obj.exists(ctx) {
		return errors.New("Create connection on already existing id")
	}
	obj.connection.Set(ctx, c)
	return nil
}

func (obj Object) exists(ctx sdk.Context) bool {
	return obj.connection.Exists(ctx)
}

func (obj Object) remove(ctx sdk.Context) {
	obj.connection.Delete(ctx)
	obj.state.Delete(ctx)
	obj.nexttimeout.Delete(ctx)
}

func (obj Object) assertSymmetric(ctx sdk.Context) error {
	conn := obj.Value(ctx)
	if !obj.remote.Value(ctx).Equal(Connection{
		Counterparty:       obj.id,
		Client:             conn.CounterpartyClient,
		CounterpartyClient: conn.Client,
	}) {
		return errors.New("unexpected counterparty connection value")
	}

	return nil
}

func assertTimeout(ctx sdk.Context, timeoutHeight uint64) error {
	if ctx.BlockHeight() > int64(timeoutHeight) {
		return errors.New("timeout")
	}

	return nil
}

func (obj Object) Value(ctx sdk.Context) (res Connection) {
	obj.connection.Get(ctx, &res)
	return
}

func (obj Object) OpenInit(ctx sdk.Context, desiredCounterparty, client, counterpartyClient string, nextTimeoutHeight uint64) error {
	if obj.exists(ctx) {
		panic("init on existing connection")
	}

	if !obj.state.Transit(ctx, Idle, Init) {
		panic("init on non-idle connection")
	}

	obj.connection.Set(ctx, Connection{
		Counterparty:       desiredCounterparty,
		Client:             client,
		CounterpartyClient: counterpartyClient,
	})

	obj.nexttimeout.Set(ctx, int64(nextTimeoutHeight))

	return nil
}

func (obj Object) OpenTry(ctx sdk.Context, expheight uint64, timeoutHeight, nextTimeoutHeight uint64) error {
	if !obj.state.Transit(ctx, Idle, Try) {
		return errors.New("invalid state")
	}

	err := assertTimeout(ctx, timeoutHeight)
	if err != nil {
		return err
	}

	if obj.remote.state.Get(ctx) != Init {
		return errors.New("counterparty state not init")
	}

	err = obj.assertSymmetric(ctx)
	if err != nil {
		return err
	}

	if obj.remote.nexttimeout.Get(ctx) != int64(timeoutHeight) {
		return errors.New("unexpected counterparty timeout value")
	}

	var expected client.Client
	obj.self.Get(ctx, expheight, &expected)
	if !client.Equal(obj.remote.client.Value(ctx), expected) {
		return errors.New("unexpected counterparty client value")
	}

	// CONTRACT: OpenTry() should be called after man.Create(), not man.Query(),
	// which will ensure
	// assert(get("connections/{desiredIdentifier}") === null) and
	// set("connections{identifier}", connection)

	obj.nexttimeout.Set(ctx, int64(nextTimeoutHeight))

	return nil
}

func (obj Object) OpenAck(ctx sdk.Context, expheight uint64, timeoutHeight, nextTimeoutHeight uint64) error {
	if !obj.state.Transit(ctx, Init, Open) {
		panic("ack on non-init connection")
	}

	err := assertTimeout(ctx, timeoutHeight)
	if err != nil {
		return err
	}

	if obj.remote.state.Get(ctx) != Try {
		return errors.New("counterparty state not try")
	}

	err = obj.assertSymmetric(ctx)
	if err != nil {
		return err
	}

	if obj.remote.nexttimeout.Get(ctx) != int64(timeoutHeight) {
		return errors.New("unexpected counterparty timeout value")
	}

	var expected client.Client
	obj.self.Get(ctx, expheight, &expected)
	if !client.Equal(obj.remote.client.Value(ctx), expected) {
		return errors.New("unexpected counterparty client value")
	}

	obj.nexttimeout.Set(ctx, int64(nextTimeoutHeight))

	return nil
}

func (obj Object) OpenConfirm(ctx sdk.Context, timeoutHeight uint64) error {
	if !obj.state.Transit(ctx, Try, Open) {
		return errors.New("confirm on non-try connection")
	}

	err := assertTimeout(ctx, timeoutHeight)
	if err != nil {
		return err
	}

	if obj.remote.state.Get(ctx) != Open {
		return errors.New("counterparty state not open")
	}

	err = obj.assertSymmetric(ctx)
	if err != nil {
		return err
	}

	if obj.remote.nexttimeout.Get(ctx) != int64(timeoutHeight) {
		return errors.New("unexpected counterparty timeout value")
	}

	obj.nexttimeout.Set(ctx, 0)

	return nil
}

func (obj Object) OpenTimeout(ctx sdk.Context) error {
	if !(obj.client.Value(ctx).Base().Height() > obj.nexttimeout.Get(ctx)) {
		return errors.New("timeout height not yet reached")
	}

	switch obj.state.Get(ctx) {
	case Init:
		if obj.remote.exists(ctx) {
			return errors.New("counterparty connection existw")
		}
	case Try:
		// XXX
	case Open:
		// XXX
	}

	obj.remove(ctx)

	return nil
}
