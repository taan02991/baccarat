package baccarat

import (
<<<<<<< HEAD
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
=======
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/types"
	"github.com/blockchain/baccarat/x/baccarat/keeper"
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	var game = types.Game{
<<<<<<< HEAD
		Creator: msg.Creator,
		ID:      msg.ID,
		Participant: []sdk.AccAddress{
			msg.Creator,
		},
		State:      0,
		Result:     []types.Card{},
		ResultHash: []string{},
=======
		Creator: 		msg.Creator,
		ID:      		msg.ID,
		State:			types.Waiting,
		Participant: 	[]sdk.AccAddress {msg.Creator},
		Result: 		make([]string, 0, 8),
		ResultHash: 	make([]string, 0, 8),
		Bet: 			make([][]types.Bet, 8, 8),
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
	}
	k.CreateGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
