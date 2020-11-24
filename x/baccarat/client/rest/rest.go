package rest

import (
	"flag"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

var addr = flag.String("addr", "localhost:4000", "http service address")

// RegisterRoutes registers baccarat-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
<<<<<<< HEAD
	// this line is used by starport scaffolding

	r.HandleFunc("/baccarat/game", listGameHandler(cliCtx, "baccarat")).Methods("GET")
	r.HandleFunc("/baccarat/game", createGameHandler(cliCtx)).Methods("POST")
=======
  // this line is used by starport scaffolding
	r.HandleFunc("/baccarat/game", listGameHandler(cliCtx, "baccarat")).Methods("GET")
	r.HandleFunc("/baccarat/game/{id}", getGameHandler(cliCtx, "baccarat")).Methods("GET")
	r.HandleFunc("/baccarat/game", createGameHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/baccarat/game/start", startGameHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/baccarat/game/bet", betHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/baccarat/game/participant", editParticipantHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/baccarat/user", listUserHandler(cliCtx, "baccarat")).Methods("GET")
	r.HandleFunc("/baccarat/user", createUserHandler(cliCtx)).Methods("POST")
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
}
