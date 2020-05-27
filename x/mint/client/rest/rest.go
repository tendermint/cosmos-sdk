package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
}
