package baccarat

import (
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgGenerateCard(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	var game = types.Game{
		Creator: msg.Creator,
		ID:      msg.ID,
		Participant: []sdk.AccAddress{
			msg.Creator,
		},
		State:      0,
		Result:     []types.Card{},
		ResultHash: []string{},
	}
	k.CreateGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
