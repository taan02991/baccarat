package baccarat

import (
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgEditParticipant(ctx sdk.Context, k keeper.Keeper, msg types.MsgEditParticipant) (*sdk.Result, error) {
	k.EditParticipant(ctx, msg.ID, msg.Creator, msg.Action)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
