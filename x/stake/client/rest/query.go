package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/stake"
	"github.com/cosmos/cosmos-sdk/x/stake/tags"
	"github.com/cosmos/cosmos-sdk/x/stake/types"

	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	"github.com/cosmos/cosmos-sdk/client/httputils"
)

const storeName = "stake"

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec) {

	// Get all delegations (delegation, undelegation and redelegation) from a delegator
	r.HandleFunc(
		"/stake/delegators/{delegatorAddr}",
		delegatorHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Get all staking txs (i.e msgs) from a delegator
	r.HandleFunc(
		"/stake/delegators/{delegatorAddr}/txs",
		delegatorTxsHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Query all validators that a delegator is bonded to
	r.HandleFunc(
		"/stake/delegators/{delegatorAddr}/validators",
		delegatorValidatorsHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Query a validator that a delegator is bonded to
	r.HandleFunc(
		"/stake/delegators/{delegatorAddr}/validators/{validatorAddr}",
		delegatorValidatorHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Query a delegation between a delegator and a validator
	r.HandleFunc(
		"/stake/delegators/{delegatorAddr}/delegations/{validatorAddr}",
		delegationHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Query all unbonding_delegations between a delegator and a validator
	r.HandleFunc(
		"/stake/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr}",
		unbondingDelegationsHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Get all validators
	r.HandleFunc(
		"/stake/validators",
		validatorsHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Get a single validator info
	r.HandleFunc(
		"/stake/validators/{addr}",
		validatorHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Get the current state of the staking pool
	r.HandleFunc(
		"/stake/pool",
		poolHandlerFn(cliCtx, cdc),
	).Methods("GET")

	// Get the current staking parameter values
	r.HandleFunc(
		"/stake/parameters",
		paramsHandlerFn(cliCtx, cdc),
	).Methods("GET")

}

// already resolve the rational shares to not handle this in the client

// defines a delegation without type Rat for shares
type DelegationWithoutRat struct {
	DelegatorAddr sdk.AccAddress `json:"delegator_addr"`
	ValidatorAddr sdk.AccAddress `json:"validator_addr"`
	Shares        string         `json:"shares"`
	Height        int64          `json:"height"`
}

// aggregation of all delegations, unbondings and redelegations
type DelegationSummary struct {
	Delegations          []DelegationWithoutRat      `json:"delegations"`
	UnbondingDelegations []stake.UnbondingDelegation `json:"unbonding_delegations"`
	Redelegations        []stake.Redelegation        `json:"redelegations"`
}

// HTTP request handler to query a delegator delegations
func delegatorHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var validatorAddr sdk.AccAddress
		var delegationSummary = DelegationSummary{}

		// read parameters
		vars := mux.Vars(r)
		bech32delegator := vars["delegatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Get all validators using key
		validators, statusCode, errMsg, err := getBech32Validators(storeName, cliCtx, cdc)
		if err != nil {
			w.WriteHeader(statusCode)
			w.Write([]byte(fmt.Sprintf("%s%s", errMsg, err.Error())))
			return
		}

		for _, validator := range validators {
			validatorAddr = validator.Operator

			// Delegations
			delegations, statusCode, errMsg, err := getDelegatorDelegations(cliCtx, cdc, delegatorAddr, validatorAddr)
			if err != nil {
				w.WriteHeader(statusCode)
				w.Write([]byte(fmt.Sprintf("%s%s", errMsg, err.Error())))
				return
			}
			if statusCode != http.StatusNoContent {
				delegationSummary.Delegations = append(delegationSummary.Delegations, delegations)
			}

			// Undelegations
			unbondingDelegation, statusCode, errMsg, err := getDelegatorUndelegations(cliCtx, cdc, delegatorAddr, validatorAddr)
			if err != nil {
				w.WriteHeader(statusCode)
				w.Write([]byte(fmt.Sprintf("%s%s", errMsg, err.Error())))
				return
			}
			if statusCode != http.StatusNoContent {
				delegationSummary.UnbondingDelegations = append(delegationSummary.UnbondingDelegations, unbondingDelegation)
			}

			// Redelegations
			// only querying redelegations to a validator as this should give us already all relegations
			// if we also would put in redelegations from, we would have every redelegation double
			redelegations, statusCode, errMsg, err := getDelegatorRedelegations(cliCtx, cdc, delegatorAddr, validatorAddr)
			if err != nil {
				w.WriteHeader(statusCode)
				w.Write([]byte(fmt.Sprintf("%s%s", errMsg, err.Error())))
				return
			}
			if statusCode != http.StatusNoContent {
				delegationSummary.Redelegations = append(delegationSummary.Redelegations, redelegations)
			}
		}

		output, err := cdc.MarshalJSON(delegationSummary)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}

// nolint gocyclo
// HTTP request handler to query all staking txs (msgs) from a delegator
func delegatorTxsHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var output []byte
		var typesQuerySlice []string
		vars := mux.Vars(r)
		delegatorAddr := vars["delegatorAddr"]

		_, err := sdk.AccAddressFromBech32(delegatorAddr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		node, err := cliCtx.GetNode()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Couldn't get current Node information. Error: %s", err.Error())))
			return
		}

		// Get values from query

		typesQuery := r.URL.Query().Get("type")
		trimmedQuery := strings.TrimSpace(typesQuery)
		if len(trimmedQuery) != 0 {
			typesQuerySlice = strings.Split(trimmedQuery, " ")
		}

		noQuery := len(typesQuerySlice) == 0
		isBondTx := contains(typesQuerySlice, "bond")
		isUnbondTx := contains(typesQuerySlice, "unbond")
		isRedTx := contains(typesQuerySlice, "redelegate")
		var txs = []tx.Info{}
		var actions []string

		switch {
		case isBondTx:
			actions = append(actions, string(tags.ActionDelegate))
		case isUnbondTx:
			actions = append(actions, string(tags.ActionBeginUnbonding))
			actions = append(actions, string(tags.ActionCompleteUnbonding))
		case isRedTx:
			actions = append(actions, string(tags.ActionBeginRedelegation))
			actions = append(actions, string(tags.ActionCompleteRedelegation))
		case noQuery:
			actions = append(actions, string(tags.ActionDelegate))
			actions = append(actions, string(tags.ActionBeginUnbonding))
			actions = append(actions, string(tags.ActionCompleteUnbonding))
			actions = append(actions, string(tags.ActionBeginRedelegation))
			actions = append(actions, string(tags.ActionCompleteRedelegation))
		default:
			w.WriteHeader(http.StatusNoContent)
			return
		}

		for _, action := range actions {
			foundTxs, errQuery := queryTxs(node, cdc, action, delegatorAddr)
			if errQuery != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("error querying transactions. Error: %s", errQuery.Error())))
			}
			txs = append(txs, foundTxs...)
		}

		output, err = cdc.MarshalJSON(txs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(output)
	}
}

// HTTP request handler to query an unbonding-delegation
func unbondingDelegationsHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32delegator := vars["delegatorAddr"]
		bech32validator := vars["validatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		validatorAddr, err := sdk.AccAddressFromBech32(bech32validator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		validatorAddrAcc := sdk.AccAddress(validatorAddr)

		key := stake.GetUBDKey(delegatorAddr, validatorAddrAcc)

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query unbonding-delegation. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this record
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ubd, err := types.UnmarshalUBD(cdc, key, res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't unmarshall unbonding-delegation. Error: %s", err.Error())))
			return
		}

		// unbondings will be a list in the future but is not yet, but we want to keep the API consistent
		ubdArray := []stake.UnbondingDelegation{ubd}

		output, err := cdc.MarshalJSON(ubdArray)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't marshall unbonding-delegation. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}

// HTTP request handler to query a bonded validator
func delegationHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read parameters
		vars := mux.Vars(r)
		bech32delegator := vars["delegatorAddr"]
		bech32validator := vars["validatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		validatorAddr, err := sdk.AccAddressFromBech32(bech32validator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		validatorAddrAcc := sdk.AccAddress(validatorAddr)

		key := stake.GetDelegationKey(delegatorAddr, validatorAddrAcc)

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query delegation. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this record
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		delegation, err := types.UnmarshalDelegation(cdc, key, res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		outputDelegation := DelegationWithoutRat{
			DelegatorAddr: delegation.DelegatorAddr,
			ValidatorAddr: delegation.ValidatorAddr,
			Height:        delegation.Height,
			Shares:        delegation.Shares.String(),
		}

		output, err := cdc.MarshalJSON(outputDelegation)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}

// HTTP request handler to query all delegator bonded validators
func delegatorValidatorsHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var validatorAccAddr sdk.AccAddress
		var bondedValidators []types.BechValidator

		// read parameters
		vars := mux.Vars(r)
		bech32delegator := vars["delegatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Get all validators using key
		kvs, err := cliCtx.QuerySubspace(stake.ValidatorsKey, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query validators. Error: %s", err.Error())))
			return
		} else if len(kvs) == 0 {
			// the query will return empty if there are no validators
			w.WriteHeader(http.StatusNoContent)
			return
		}

		validators, err := getValidators(kvs, cdc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
			return
		}

		for _, validator := range validators {
			// get all transactions from the delegator to val and append
			validatorAccAddr = validator.Operator

			validator, statusCode, errMsg, errRes := getDelegatorValidator(cliCtx, cdc, delegatorAddr, validatorAccAddr)
			if errRes != nil {
				w.WriteHeader(statusCode)
				w.Write([]byte(fmt.Sprintf("%s%s", errMsg, errRes.Error())))
				return
			} else if statusCode == http.StatusNoContent {
				continue
			}

			bondedValidators = append(bondedValidators, validator)
		}
		output, err := cdc.MarshalJSON(bondedValidators)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(output)
	}
}

// HTTP request handler to get information from a currently bonded validator
func delegatorValidatorHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read parameters
		var output []byte
		vars := mux.Vars(r)
		bech32delegator := vars["delegatorAddr"]
		bech32validator := vars["validatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		validatorAccAddr, err := sdk.AccAddressFromBech32(bech32validator)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
			return
		}

		// Check if there if the delegator is bonded or redelegated to the validator

		validator, statusCode, errMsg, err := getDelegatorValidator(cliCtx, cdc, delegatorAddr, validatorAccAddr)
		if err != nil {
			w.WriteHeader(statusCode)
			w.Write([]byte(fmt.Sprintf("%s%s", errMsg, err.Error())))
			return
		} else if statusCode == http.StatusNoContent {
			w.WriteHeader(statusCode)
			return
		}
		output, err = cdc.MarshalJSON(validator)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(output)
	}
}

// TODO bech32
// http request handler to query list of validators
func validatorsHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kvs, err := cliCtx.QuerySubspace(stake.ValidatorsKey, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query validators. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there are no validators
		if len(kvs) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		validators, err := getValidators(kvs, cdc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
			return
		}

		output, err := cdc.MarshalJSON(validators)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}

// HTTP request handler to query the validator information from a given validator address
func validatorHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var output []byte
		// read parameters
		vars := mux.Vars(r)
		bech32validatorAddr := vars["addr"]
		valAddress, err := sdk.AccAddressFromBech32(bech32validatorAddr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
			return
		}

		key := stake.GetValidatorKey(valAddress)

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query validator, error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this record
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		validator, err := types.UnmarshalValidator(cdc, valAddress, res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		bech32Validator, err := validator.Bech32Validator()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		output, err = cdc.MarshalJSON(bech32Validator)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
			return
		}

		if output == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Write(output)
	}
}

func registerSwaggerQueryRoutes(routerGroup *gin.RouterGroup, ctx context.CLIContext, cdc *wire.Codec) {
	routerGroup.GET("/stake/delegators/:delegatorAddr", delegatorHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/delegators/:delegatorAddr/txs", delegatorTxsHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/delegators/:delegatorAddr/validators", delegatorValidatorsHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/delegators/:delegatorAddr/validators/:validatorAddr", delegatorValidatorHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/delegators/:delegatorAddr/delegations/:validatorAddr", delegationHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/delegators/:delegatorAddr/unbonding_delegations/:validatorAddr", unbondingDelegationsHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/validators", validatorsHandlerCreation(cdc, ctx))
	routerGroup.GET("/stake/validators/:addr", validatorHandlerCreation(cdc, ctx))
}

func delegatorHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		var validatorAddr sdk.AccAddress
		var delegationSummary = DelegationSummary{}

		// read parameters
		bech32delegator := gtx.Param("delegatorAddr")

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		// Get all validators using key
		validators, statusCode, errMsg, err := getBech32Validators(storeName, cliCtx, cdc)
		if err != nil {
			httputils.NewError(gtx, statusCode, fmt.Errorf("%s%s", errMsg, err.Error()))
			return
		}

		for _, validator := range validators {
			validatorAddr = validator.Operator

			// Delegations
			delegations, statusCode, errMsg, err := getDelegatorDelegations(cliCtx, cdc, delegatorAddr, validatorAddr)
			if err != nil {
				httputils.NewError(gtx, statusCode, fmt.Errorf("%s%s", errMsg, err.Error()))
				return
			}
			if statusCode != http.StatusNoContent {
				delegationSummary.Delegations = append(delegationSummary.Delegations, delegations)
			}

			// Undelegations
			unbondingDelegation, statusCode, errMsg, err := getDelegatorUndelegations(cliCtx, cdc, delegatorAddr, validatorAddr)
			if err != nil {
				httputils.NewError(gtx, statusCode, fmt.Errorf("%s%s", errMsg, err.Error()))
				return
			}
			if statusCode != http.StatusNoContent {
				delegationSummary.UnbondingDelegations = append(delegationSummary.UnbondingDelegations, unbondingDelegation)
			}

			// Redelegations
			// only querying redelegations to a validator as this should give us already all relegations
			// if we also would put in redelegations from, we would have every redelegation double
			redelegations, statusCode, errMsg, err := getDelegatorRedelegations(cliCtx, cdc, delegatorAddr, validatorAddr)
			if err != nil {
				httputils.NewError(gtx, statusCode, fmt.Errorf("%s%s", errMsg, err.Error()))
				return
			}
			if statusCode != http.StatusNoContent {
				delegationSummary.Redelegations = append(delegationSummary.Redelegations, redelegations)
			}
		}

		output, err := cdc.MarshalJSON(delegationSummary)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}

		httputils.NormalResponse(gtx, output)
	}
}

// nolint: gocyclo
func delegatorTxsHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		var output []byte
		var typesQuerySlice []string

		delegatorAddr := gtx.Param("delegatorAddr")

		_, err := sdk.AccAddressFromBech32(delegatorAddr)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		node, err := cliCtx.GetNode()
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("Couldn't get current Node information. Error: %s", err.Error()))
			return
		}

		// Get values from query
		typesQuery := gtx.Query("type")
		trimmedQuery := strings.TrimSpace(typesQuery)
		if len(trimmedQuery) != 0 {
			typesQuerySlice = strings.Split(trimmedQuery, " ")
		}

		noQuery := len(typesQuerySlice) == 0
		isBondTx := contains(typesQuerySlice, "bond")
		isUnbondTx := contains(typesQuerySlice, "unbond")
		isRedTx := contains(typesQuerySlice, "redelegate")
		var txs = []tx.Info{}
		var actions []string

		switch {
		case isBondTx:
			actions = append(actions, string(tags.ActionDelegate))
		case isUnbondTx:
			actions = append(actions, string(tags.ActionBeginUnbonding))
			actions = append(actions, string(tags.ActionCompleteUnbonding))
		case isRedTx:
			actions = append(actions, string(tags.ActionBeginRedelegation))
			actions = append(actions, string(tags.ActionCompleteRedelegation))
		case noQuery:
			actions = append(actions, string(tags.ActionDelegate))
			actions = append(actions, string(tags.ActionBeginUnbonding))
			actions = append(actions, string(tags.ActionCompleteUnbonding))
			actions = append(actions, string(tags.ActionBeginRedelegation))
			actions = append(actions, string(tags.ActionCompleteRedelegation))
		default:
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}

		for _, action := range actions {
			foundTxs, errQuery := queryTxs(node, cdc, action, delegatorAddr)
			if errQuery != nil {
				httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("error querying transactions. Error: %s", errQuery.Error()))
			}
			txs = append(txs, foundTxs...)
		}

		output, err = cdc.MarshalJSON(txs)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}
		httputils.NormalResponse(gtx, output)
	}
}

func delegatorValidatorsHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		var validatorAccAddr sdk.AccAddress
		var bondedValidators []types.BechValidator

		// read parameters
		bech32delegator := gtx.Param("delegatorAddr")

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		// Get all validators using key
		kvs, err := cliCtx.QuerySubspace(stake.ValidatorsKey, storeName)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query validators. Error: %s", err.Error()))
			return
		} else if len(kvs) == 0 {
			// the query will return empty if there are no validators
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}

		validators, err := getValidators(kvs, cdc)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}

		for _, validator := range validators {
			// get all transactions from the delegator to val and append
			validatorAccAddr = validator.Operator

			validator, statusCode, errMsg, errRes := getDelegatorValidator(cliCtx, cdc, delegatorAddr, validatorAccAddr)
			if errRes != nil {
				httputils.NewError(gtx, statusCode, fmt.Errorf("%s%s", errMsg, errRes.Error()))
				return
			} else if statusCode == http.StatusNoContent {
				continue
			}

			bondedValidators = append(bondedValidators, validator)
		}
		output, err := cdc.MarshalJSON(bondedValidators)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}
		httputils.NormalResponse(gtx, output)
	}
}

func delegatorValidatorHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		// read parameters
		var output []byte
		bech32delegator := gtx.Param("delegatorAddr")
		bech32validator := gtx.Param("validatorAddr")

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		validatorAccAddr, err := sdk.AccAddressFromBech32(bech32validator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		// Check if there if the delegator is bonded or redelegated to the validator

		validator, statusCode, errMsg, err := getDelegatorValidator(cliCtx, cdc, delegatorAddr, validatorAccAddr)
		if err != nil {
			httputils.NewError(gtx, statusCode, fmt.Errorf("%s%s", errMsg, err.Error()))
			return
		} else if statusCode == http.StatusNoContent {
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}
		output, err = cdc.MarshalJSON(validator)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}
		httputils.NormalResponse(gtx, output)
	}
}

func delegationHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		// read parameters
		bech32delegator := gtx.Param("delegatorAddr")
		bech32validator := gtx.Param("validatorAddr")

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		validatorAddr, err := sdk.AccAddressFromBech32(bech32validator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}
		validatorAddrAcc := sdk.AccAddress(validatorAddr)

		key := stake.GetDelegationKey(delegatorAddr, validatorAddrAcc)

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query delegation. Error: %s", err.Error()))
			return
		}

		// the query will return empty if there is no data for this record
		if len(res) == 0 {
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}

		delegation, err := types.UnmarshalDelegation(cdc, key, res)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		outputDelegation := DelegationWithoutRat{
			DelegatorAddr: delegation.DelegatorAddr,
			ValidatorAddr: delegation.ValidatorAddr,
			Height:        delegation.Height,
			Shares:        delegation.Shares.String(),
		}

		output, err := cdc.MarshalJSON(outputDelegation)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}

		httputils.NormalResponse(gtx, output)
	}
}

func unbondingDelegationsHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		bech32delegator := gtx.Param("delegatorAddr")
		bech32validator := gtx.Param("validatorAddr")

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		validatorAddr, err := sdk.AccAddressFromBech32(bech32validator)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}
		validatorAddrAcc := sdk.AccAddress(validatorAddr)

		key := stake.GetUBDKey(delegatorAddr, validatorAddrAcc)

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query unbonding-delegation. Error: %s", err.Error()))
			return
		}

		// the query will return empty if there is no data for this record
		if len(res) == 0 {
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}

		ubd, err := types.UnmarshalUBD(cdc, key, res)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't unmarshall unbonding-delegation. Error: %s", err.Error()))
			return
		}

		// unbondings will be a list in the future but is not yet, but we want to keep the API consistent
		ubdArray := []stake.UnbondingDelegation{ubd}

		output, err := cdc.MarshalJSON(ubdArray)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't marshall unbonding-delegation. Error: %s", err.Error()))
			return
		}

		httputils.NormalResponse(gtx, output)
	}
}

func validatorsHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		kvs, err := cliCtx.QuerySubspace(stake.ValidatorsKey, storeName)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query validators. Error: %s", err.Error()))
			return
		}

		// the query will return empty if there are no validators
		if len(kvs) == 0 {
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}

		validators, err := getValidators(kvs, cdc)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}

		output, err := cdc.MarshalJSON(validators)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, err)
			return
		}

		httputils.NormalResponse(gtx, output)
	}
}

func validatorHandlerCreation(cdc *wire.Codec, cliCtx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		var output []byte
		// read parameters
		bech32validatorAddr := gtx.Param("addr")
		valAddress, err := sdk.AccAddressFromBech32(bech32validatorAddr)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		key := stake.GetValidatorKey(valAddress)

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query validator, error: %s", err.Error()))
			return
		}

		// the query will return empty if there is no data for this record
		if len(res) == 0 {
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}

		validator, err := types.UnmarshalValidator(cdc, valAddress, res)
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		bech32Validator, err := validator.Bech32Validator()
		if err != nil {
			httputils.NewError(gtx, http.StatusBadRequest, err)
			return
		}

		output, err = cdc.MarshalJSON(bech32Validator)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("Error: %s", err.Error()))
			return
		}

		if output == nil {
			httputils.NewError(gtx, http.StatusNoContent, nil)
			return
		}
		httputils.NormalResponse(gtx, output)
	}
}
=======
// HTTP request handler to query the pool information
func poolHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := stake.PoolKey

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query pool. Error: %s", err.Error())))
			return
		}

		pool, err := types.UnmarshalPool(cdc, res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		output, err := cdc.MarshalJSON(pool)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}

// HTTP request handler to query the staking params values
func paramsHandlerFn(cliCtx context.CLIContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := stake.ParamKey

		res, err := cliCtx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query parameters. Error: %s", err.Error())))
			return
		}

		params, err := types.UnmarshalParams(cdc, res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		output, err := cdc.MarshalJSON(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}
