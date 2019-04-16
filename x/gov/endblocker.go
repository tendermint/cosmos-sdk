package gov

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/tags"
)

// Called every block, process inflation, update validator set
func EndBlocker(ctx sdk.Context, keeper Keeper) sdk.Tags {
	logger := keeper.Logger(ctx)
	resTags := sdk.NewTags()

	inactiveIterator := keeper.InactiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	defer inactiveIterator.Close()
	for ; inactiveIterator.Valid(); inactiveIterator.Next() {
		var proposalID uint64

		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(inactiveIterator.Value(), &proposalID)
		inactiveProposal, ok := keeper.GetProposal(ctx, proposalID)
		if !ok {
			panic(fmt.Sprintf("proposal %d does not exist", proposalID))
		}

		keeper.DeleteProposal(ctx, proposalID)
		keeper.DeleteDeposits(ctx, proposalID) // delete any associated deposits (burned)

		resTags = resTags.AppendTag(tags.ProposalID, fmt.Sprintf("%d", proposalID))
		resTags = resTags.AppendTag(tags.ProposalResult, tags.ActionProposalDropped)

		logger.Info(
			fmt.Sprintf("proposal %d (%s) didn't meet minimum deposit of %s (had only %s); deleted",
				inactiveProposal.ProposalID,
				inactiveProposal.GetTitle(),
				keeper.GetDepositParams(ctx).MinDeposit,
				inactiveProposal.TotalDeposit,
			),
		)
	}

	// fetch active proposals whose voting periods have ended (are passed the block time)
	activeIterator := keeper.ActiveProposalQueueIterator(ctx, ctx.BlockHeader().Time)
	defer activeIterator.Close()
	for ; activeIterator.Valid(); activeIterator.Next() {
		var proposalID uint64

		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(activeIterator.Value(), &proposalID)
		activeProposal, ok := keeper.GetProposal(ctx, proposalID)
		if !ok {
			panic(fmt.Sprintf("proposal %d does not exist", proposalID))
		}
		passes, tallyResults := tally(ctx, keeper, activeProposal)

		var tagValue string
		var logMsg string
		if passes {
			keeper.RefundDeposits(ctx, activeProposal.ProposalID)
			activeProposal.Status = StatusPassed

			handler := keeper.router.GetRoute(activeProposal.ProposalRoute())

			cacheCtx, writeCache := ctx.CacheContext()

			// The handler is state mutating logic depending on the proposal content.
			// The handler may mutate the state or not
			// if handler fails no state mutation happened and the err msg is logged.
			err := handler(cacheCtx, activeProposal.Content)
			if err == nil {
				tagValue = tags.ActionProposalPassed
				logMsg = "passed"

				// writes state mutation to the underlying multistore
				writeCache()
			} else {
				logMsg = fmt.Sprintf("passed, but failed on execution: %s", err.ABCILog())
				tagValue = tags.ActionProposalFailed
			}
		} else {
			keeper.DeleteDeposits(ctx, activeProposal.ProposalID)
			activeProposal.Status = StatusRejected
			logMsg = "rejected"
			tagValue = tags.ActionProposalRejected
		}

		activeProposal.FinalTallyResult = tallyResults
		keeper.SetProposal(ctx, activeProposal)
		keeper.RemoveFromActiveProposalQueue(ctx, activeProposal.VotingEndTime, activeProposal.ProposalID)

		logger.Info(
			fmt.Sprintf(
				"proposal %d (%s) %s",
				activeProposal.ProposalID, activeProposal.GetTitle(), logMsg,
			),
		)

		resTags = resTags.AppendTag(tags.ProposalID, fmt.Sprintf("%d", proposalID))
		resTags = resTags.AppendTag(tags.ProposalResult, tagValue)
	}

	return resTags
}
