package baccarat

import (
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreateUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateUser) (*sdk.Result, error) {
	var user = types.User{
		Creator: msg.Creator,
		ID:      msg.ID,
		Name:    msg.Name,
		Addr:    msg.Addr,
	}
	k.CreateUser(ctx, user)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
