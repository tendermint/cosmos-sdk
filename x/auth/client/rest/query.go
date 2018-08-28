package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"

	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	"github.com/cosmos/cosmos-sdk/client/httputils"
)

// register REST routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec, storeName string) {
	r.HandleFunc(
		"/auth/accounts/{address}",
		QueryAccountRequestHandlerFn(storeName, cdc, authcmd.GetAccountDecoder(cdc), cliCtx),
	).Methods("GET")
}

// query accountREST Handler
func QueryAccountRequestHandlerFn(
	storeName string, cdc *wire.Codec,
	decoder auth.AccountDecoder, cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32addr := vars["address"]

		addr, err := sdk.AccAddressFromBech32(bech32addr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := cliCtx.QueryStore(auth.AddressStoreKey(addr), storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query account. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// decode the value
		account, err := decoder(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		// print out whole account
		output, err := cdc.MarshalJSON(account)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}

// RegisterSwaggerRoutes registers account query related routes to Gaia-lite server
func RegisterSwaggerRoutes(routerGroup *gin.RouterGroup, ctx context.CLIContext, cdc *wire.Codec, storeName string) {
	routerGroup.GET("/auth/accounts/:address",queryAccountRequestHandler(storeName,cdc,authcmd.GetAccountDecoder(cdc),ctx))
}

func queryAccountRequestHandler(storeName string, cdc *wire.Codec, decoder auth.AccountDecoder, ctx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {

		bech32addr := gtx.Param("address")

		addr, err := sdk.AccAddressFromBech32(bech32addr)
		if err != nil {
			httputils.NewError(gtx, http.StatusConflict, err)
			return
		}

		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), storeName)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query account. Error: %s", err.Error()))
			return
		}

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			httputils.NewError(gtx, http.StatusNoContent, fmt.Errorf("this account info is nil+"))
			return
		}

		// decode the value
		account, err := decoder(res)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't parse query result. Result: %s. Error: %s", res, err.Error()))
			return
		}

		// print out whole account
		output, err := cdc.MarshalJSON(account)
		if err != nil {
			httputils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't marshall query result. Error: %s", err.Error()))
			return
		}

		httputils.NormalResponse(gtx,output)
	}
}