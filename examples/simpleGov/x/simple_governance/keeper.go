package simpleGovernance

import (
	stake "github.com/cosmos/cosmos-sdk/examples/democoin/x/simplestake"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
)

// nolint
type Keeper struct {
	ProposalStoreKey sdk.StoreKey
	codespace        sdk.CodespaceType
	cdc              *wire.Codec

	ck bank.Keeper
	sm stake.Keeper
}

// nolint
type KeeperRead struct {
	Keeper
}

// NewKeeper crates a new keeper with write and read access
func NewKeeper(proposalStoreKey sdk.StoreKey, ck bank.Keeper, sm stake.Keeper, codespace sdk.CodespaceType) Keeper {
	cdc := wire.NewCodec()

	return Keeper{
		ProposalStoreKey: proposalStoreKey,
		cdc:              cdc,
		ck:               ck,
		sm:               sm,
		codespace:        codespace,
	}
}

// NewKeeperRead crates a new keeper with read access
func NewKeeperRead(proposalStoreKey sdk.StoreKey, ck bank.Keeper, sm stake.Keeper, codespace sdk.CodespaceType) KeeperRead {
	cdc := wire.NewCodec()

	return KeeperRead{Keeper{
		ProposalStoreKey: proposalStoreKey,
		cdc:              cdc,
		ck:               ck,
		sm:               sm,
		codespace:        codespace,
	}}
}

// GetProposal gets the proposal with the given id from the context
func (k Keeper) GetProposal(ctx sdk.Context, proposalID int64) (Proposal, sdk.Error) {
	store := ctx.KVStore(k.ProposalStoreKey)

	bpi, err := k.cdc.MarshalBinary(proposalID)
	if err != nil {
		panic(err)
	}
	bp := store.Get(bpi)
	if bp == nil {
		return Proposal{}, ErrProposalNotFound(proposalID)
	}

	proposal := Proposal{}

	err = k.cdc.UnmarshalBinary(bp, proposal)
	if err != nil {
		panic(err)
	}

	return proposal, nil
}

// SetProposal sets a proposal to the context
func (k Keeper) SetProposal(ctx sdk.Context, proposalID int64, proposal Proposal) sdk.Error {
	store := ctx.KVStore(k.ProposalStoreKey)

	bp, err := k.cdc.MarshalBinary(proposal)
	if err != nil {
		panic(err) // return proper error
	}

	bpi, err := k.cdc.MarshalBinary(proposalID)
	if err != nil {
		panic(err) // return proper error
	}

	store.Set(bpi, bp)
	return nil
}

// func (k KeeperRead) SetProposal(ctx sdk.Context, proposalID int64, proposal Proposal) sdk.Error {
// 	return sdk.ErrUnauthorized("").Trace("This keeper does not have write access for the simple governance store")
// }

// NewProposalID creates a new id for a proposal
func (k Keeper) NewProposalID(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.ProposalStoreKey)

	bid := store.Get([]byte("TotalID"))
	if bid == nil {
		return 0
	}

	totalID := new(int64)
	err := k.cdc.UnmarshalBinary(bid, totalID)
	if err != nil {
		panic(err)
	}

	return (*totalID + 1)
}

//--------------------------------------------------------------------------------------

// GetOption returns the given option of a proposal stored in the keeper
func (k Keeper) GetOption(ctx sdk.Context, key []byte) (string, sdk.Error) {
	store := ctx.KVStore(k.ProposalStoreKey)

	bv := store.Get(key)
	if bv == nil {
		return "", ErrOptionNotFound()
	}

	option := new(string)

	err := k.cdc.UnmarshalBinary(bv, option)
	if err != nil {
		panic(err)
	}

	return *option, nil
}

// SetOption sets the option to the propposal stored in the context store
func (k Keeper) SetOption(ctx sdk.Context, key []byte, option string) {
	store := ctx.KVStore(k.ProposalStoreKey)

	bv, err := k.cdc.MarshalBinary(option)
	if err != nil {
		panic(err)
	}
	store.Set(key, bv)
}

// IMO not even necessary
// func (k KeeperRead) SetOption(ctx sdk.Context, key []byte, option string) sdk.Error {
// 	return sdk.ErrUnauthorized("").Trace("This keeper does not have write access for the simple governance store")
// }

//--------------------------------------------------------------------------------------

// getProposalQueue gets the ProposalQueue from the context
func (k Keeper) getProposalQueue(ctx sdk.Context) (ProposalQueue, sdk.Error) {
	store := ctx.KVStore(k.ProposalStoreKey)
	bpq := store.Get([]byte("proposalQueue"))
	if bpq == nil {
		return ProposalQueue{}, ErrProposalQueueNotFound()
	}

	proposalQueue := ProposalQueue{}
	err := k.cdc.UnmarshalBinaryBare(bpq, proposalQueue)
	if err != nil {
		panic(err)
	}

	return proposalQueue, nil
}

// setProposalQueue sets the ProposalQueue to the context
func (k Keeper) setProposalQueue(ctx sdk.Context, proposalQueue ProposalQueue) {
	store := ctx.KVStore(k.ProposalStoreKey)
	bpq, err := k.cdc.MarshalBinaryBare(proposalQueue)
	if err != nil {
		panic(err)
	}
	store.Set([]byte("proposalQueue"), bpq)
}

// func (k KeeperRead) setProposalQueue(ctx sdk.Context, proposalQueue ProposalQueue) sdk.Error {
// 	return sdk.ErrUnauthorized("").Trace("This keeper does not have write access for the simple governance store")
// }

// ProposalQueueHead returns the head of the FIFO Proposal queue
func (k Keeper) ProposalQueueHead(ctx sdk.Context) (Proposal, sdk.Error) {
	proposalQueue, err := k.getProposalQueue(ctx)
	if err != nil {
		return Proposal{}, err
	}
	if proposalQueue.IsEmpty() {
		return Proposal{}, ErrEmptyProposalQueue()
	}
	proposal, err := k.GetProposal(ctx, proposalQueue[0])
	if err != nil {
		return Proposal{}, err
	}
	return proposal, nil
}

// ProposalQueuePop pops the head from the Proposal queue
func (k Keeper) ProposalQueuePop(ctx sdk.Context) (Proposal, sdk.Error) {
	proposalQueue, err := k.getProposalQueue(ctx)
	if err != nil {
		return Proposal{}, err
	}
	if proposalQueue.IsEmpty() {
		return Proposal{}, ErrEmptyProposalQueue()
	}
	headElement, tailProposalQueue := proposalQueue[0], proposalQueue[1:]
	k.setProposalQueue(ctx, tailProposalQueue)
	proposal, err := k.GetProposal(ctx, headElement)
	if err != nil {
		return Proposal{}, err
	}
	return proposal, nil
}

// ProposalQueuePush pushes a proposal to the tail of the FIFO Proposal queue
func (k Keeper) ProposalQueuePush(ctx sdk.Context, proposaID int64) sdk.Error {
	proposalQueue, err := k.getProposalQueue(ctx)
	if err != nil {
		return err
	}
	proposalQueue = append(proposalQueue, proposaID)
	k.setProposalQueue(ctx, proposalQueue)
	return nil
}
