package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers baccarat-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/baccarat/user", listUserHandler(cliCtx, "baccarat")).Methods("GET")
	r.HandleFunc("/baccarat/user", createUserHandler(cliCtx)).Methods("POST")
}
