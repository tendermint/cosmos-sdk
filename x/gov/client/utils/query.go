package utils

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	defaultPage  = 1
	defaultLimit = 30 // should be consistent with tendermint/tendermint/rpc/core/pipe.go:19
)

// Proposer contains metadata of a governance proposal used for querying a
// proposer.
type Proposer struct {
	ProposalID uint64 `json:"proposal_id" yaml:"proposal_id"`
	Proposer   string `json:"proposer" yaml:"proposer"`
}

// NewProposer returns a new Proposer given id and proposer
func NewProposer(proposalID uint64, proposer string) Proposer {
	return Proposer{proposalID, proposer}
}

func (p Proposer) String() string {
	return fmt.Sprintf("Proposal with ID %d was proposed by %s", p.ProposalID, p.Proposer)
}

// QueryDepositsByTxQuery will query for deposits via a direct txs tags query. It
// will fetch and build deposits directly from the returned txs and return a
// JSON marshalled result or any error that occurred.
//
// NOTE: SearchTxs is used to facilitate the txs query which does not currently
// support configurable pagination.
func QueryDepositsByTxQuery(clientCtx client.Context, params types.QueryProposalParams) ([]byte, error) {
	searchResult, err := searchEvents(
		clientCtx, types.TypeMsgDeposit, types.TypeSvcMsgDeposit,
		fmt.Sprintf("%s.%s='%s'", types.EventTypeProposalDeposit, types.AttributeKeyProposalID, []byte(fmt.Sprintf("%d", params.ProposalID))),
	)
	if err != nil {
		return nil, err
	}

	var deposits []types.Deposit

	for _, info := range searchResult.Txs {
		for _, msg := range info.GetTx().GetMsgs() {
			var depMsg *types.MsgDeposit
			if msg.Type() == types.TypeSvcMsgDeposit {
				depMsg = msg.(sdk.ServiceMsg).Request.(*types.MsgDeposit)
			} else if msg.Type() == types.TypeMsgDeposit {
				depMsg = msg.(*types.MsgDeposit)
			}

			if depMsg != nil {
				deposits = append(deposits, types.Deposit{
					Depositor:  depMsg.Depositor,
					ProposalId: params.ProposalID,
					Amount:     depMsg.Amount,
				})
			}
		}
	}

	bz, err := clientCtx.LegacyAmino.MarshalJSON(deposits)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// QueryVotesByTxQuery will query for votes via a direct txs tags query. It
// will fetch and build votes directly from the returned txs and return a JSON
// marshalled result or any error that occurred.
func QueryVotesByTxQuery(clientCtx client.Context, params types.QueryProposalVotesParams) ([]byte, error) {
	var (
		votes      []types.Vote
		nextTxPage = defaultPage
		totalLimit = params.Limit * params.Page
	)

	// query interrupted either if we collected enough votes or tx indexer run out of relevant txs
	for len(votes) < totalLimit {
		searchResult, err := searchEvents(
			clientCtx, types.TypeMsgVote, types.TypeSvcMsgVote,
			fmt.Sprintf("%s.%s='%s'", types.EventTypeProposalVote, types.AttributeKeyProposalID, []byte(fmt.Sprintf("%d", params.ProposalID))),
		)
		if err != nil {
			return nil, err
		}

		nextTxPage++
		for _, info := range searchResult.Txs {
			for _, msg := range info.GetTx().GetMsgs() {
				var voteMsg *types.MsgVote
				if msg.Type() == types.TypeSvcMsgVote {
					voteMsg = msg.(sdk.ServiceMsg).Request.(*types.MsgVote)
				} else if msg.Type() == types.TypeMsgVote {
					voteMsg = msg.(*types.MsgVote)
				}

				if voteMsg != nil {
					votes = append(votes, types.Vote{
						Voter:      voteMsg.Voter,
						ProposalId: params.ProposalID,
						Option:     voteMsg.Option,
					})
				}
			}
		}
		if len(searchResult.Txs) != defaultLimit {
			break
		}
	}
	start, end := client.Paginate(len(votes), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		votes = []types.Vote{}
	} else {
		votes = votes[start:end]
	}

	bz, err := clientCtx.LegacyAmino.MarshalJSON(votes)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// QueryVoteByTxQuery will query for a single vote via a direct txs tags query.
func QueryVoteByTxQuery(clientCtx client.Context, params types.QueryVoteParams) ([]byte, error) {
	searchResult, err := searchEvents(
		clientCtx, types.TypeMsgVote, types.TypeSvcMsgVote,
		fmt.Sprintf("%s.%s='%s'", types.EventTypeProposalVote, types.AttributeKeyProposalID, []byte(fmt.Sprintf("%d", params.ProposalID))),
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeySender, []byte(params.Voter.String())),
	)
	if err != nil {
		return nil, err
	}

	for _, info := range searchResult.Txs {
		for _, msg := range info.GetTx().GetMsgs() {
			var voteMsg *types.MsgVote
			// there should only be a single vote under the given conditions
			if msg.Type() == types.TypeSvcMsgVote {
				voteMsg = msg.(sdk.ServiceMsg).Request.(*types.MsgVote)
			} else if msg.Type() == types.TypeMsgVote {
				voteMsg = msg.(*types.MsgVote)
			}

			if voteMsg != nil {
				vote := types.Vote{
					Voter:      voteMsg.Voter,
					ProposalId: params.ProposalID,
					Option:     voteMsg.Option,
				}

				bz, err := clientCtx.JSONMarshaler.MarshalJSON(&vote)
				if err != nil {
					return nil, err
				}

				return bz, nil
			}
		}
	}

	return nil, fmt.Errorf("address '%s' did not vote on proposalID %d", params.Voter, params.ProposalID)
}

// QueryDepositByTxQuery will query for a single deposit via a direct txs tags
// query.
func QueryDepositByTxQuery(clientCtx client.Context, params types.QueryDepositParams) ([]byte, error) {
	searchResult, err := searchEvents(
		clientCtx, types.TypeMsgDeposit, types.TypeSvcMsgDeposit,
		fmt.Sprintf("%s.%s='%s'", types.EventTypeProposalDeposit, types.AttributeKeyProposalID, []byte(fmt.Sprintf("%d", params.ProposalID))),
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeySender, []byte(params.Depositor.String())),
	)
	if err != nil {
		return nil, err
	}

	for _, info := range searchResult.Txs {
		for _, msg := range info.GetTx().GetMsgs() {
			var depMsg *types.MsgDeposit
			// there should only be a single deposit under the given conditions
			if msg.Type() == types.TypeSvcMsgDeposit {
				depMsg = msg.(sdk.ServiceMsg).Request.(*types.MsgDeposit)
			} else if msg.Type() == types.TypeMsgDeposit {
				depMsg = msg.(*types.MsgDeposit)
			}

			if depMsg != nil {
				deposit := types.Deposit{
					Depositor:  depMsg.Depositor,
					ProposalId: params.ProposalID,
					Amount:     depMsg.Amount,
				}

				bz, err := clientCtx.JSONMarshaler.MarshalJSON(&deposit)
				if err != nil {
					return nil, err
				}

				return bz, nil
			}
		}
	}

	return nil, fmt.Errorf("address '%s' did not deposit to proposalID %d", params.Depositor, params.ProposalID)
}

// QueryProposerByTxQuery will query for a proposer of a governance proposal by
// ID.
func QueryProposerByTxQuery(clientCtx client.Context, proposalID uint64) (Proposer, error) {
	searchResult, err := searchEvents(
		clientCtx, types.TypeMsgSubmitProposal, types.TypeSvcMsgSubmitProposal,
		fmt.Sprintf("%s.%s='%s'", types.EventTypeSubmitProposal, types.AttributeKeyProposalID, []byte(fmt.Sprintf("%d", proposalID))),
	)
	if err != nil {
		return Proposer{}, err
	}

	for _, info := range searchResult.Txs {
		for _, msg := range info.GetTx().GetMsgs() {
			// there should only be a single proposal under the given conditions
			if msg.Type() == types.TypeSvcMsgSubmitProposal {
				subMsg := msg.(sdk.ServiceMsg).Request.(*types.MsgSubmitProposal)

				return NewProposer(proposalID, subMsg.Proposer), nil
			} else if msg.Type() == types.TypeMsgSubmitProposal {
				subMsg := msg.(*types.MsgSubmitProposal)

				return NewProposer(proposalID, subMsg.Proposer), nil
			}
		}
	}

	return Proposer{}, fmt.Errorf("failed to find the proposer for proposalID %d", proposalID)
}

// QueryProposalByID takes a proposalID and returns a proposal
func QueryProposalByID(proposalID uint64, clientCtx client.Context, queryRoute string) ([]byte, error) {
	params := types.NewQueryProposalParams(proposalID)
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/proposal", queryRoute), bz)
	if err != nil {
		return nil, err
	}

	return res, err
}

func queryTxsByEvents(clientCtx client.Context, msgType string, otherEvents ...string) (*sdk.SearchTxsResult, error) {
	events := append([]string{
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeyAction, msgType),
	}, otherEvents...)

	// NOTE: SearchTxs is used to facilitate the txs query which does not currently
	// support configurable pagination.
	return authclient.QueryTxsByEvents(clientCtx, events, defaultPage, defaultLimit, "")
}

// searchEvents queries txs by events with both `oldMsgType` and `newMsgtype`,
// merges the results into one *sdk.SearchTxsResult.
func searchEvents(clientCtx client.Context, oldMsgType, newMsgType string, otherEvents ...string) (*sdk.SearchTxsResult, error) {
	// Tx are indexed in tendermint via their Msgs `Type()`, which can be:
	// - via legacy Msgs (amino or proto), their `Type()` is a custom string,
	// - via ADR-031 service msgs, their `Type()` is the protobuf FQ method name.
	// In searching for events, we search for both `Type()`s.
	oldsearchEvents, err := queryTxsByEvents(clientCtx, oldMsgType, otherEvents...)
	if err != nil {
		return nil, err
	}
	newsearchEvents, err := queryTxsByEvents(clientCtx, newMsgType, otherEvents...)
	if err != nil {
		return nil, err
	}

	return &sdk.SearchTxsResult{
		// Only the Txs field will be used by other functions, so we don't need
		// to populate the other fields.
		Txs: append(oldsearchEvents.Txs, newsearchEvents.Txs...),
	}, nil
}
