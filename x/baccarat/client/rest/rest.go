package rest

import (
	"flag"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

var addr = flag.String("addr", "localhost:4000", "http service address")

// RegisterRoutes registers baccarat-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding

	r.HandleFunc("/baccarat/game", listGameHandler(cliCtx, "baccarat")).Methods("GET")
	r.HandleFunc("/baccarat/game", createGameHandler(cliCtx)).Methods("POST")
}
