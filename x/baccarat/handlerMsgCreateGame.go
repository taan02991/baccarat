package baccarat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/types"
	"github.com/blockchain/baccarat/x/baccarat/keeper"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	var game = types.Game{
		Creator: 		msg.Creator,
		ID:      		msg.ID,
		State:			types.Waiting,
		Participant: 	[]sdk.AccAddress {msg.Creator},
		Result: 		make([]string, 0, 8),
		ResultHash: 	make([]string, 0, 8),
		Bet: 			make([][]types.Bet, 8, 8),
	}
	k.CreateGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
