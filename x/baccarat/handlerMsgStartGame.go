package baccarat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/types"
	"github.com/blockchain/baccarat/x/baccarat/keeper"
)

func handleMsgStartGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartGame) (*sdk.Result, error) {
	k.StartGame(ctx, msg.ID)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
