package baccarat

import (
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgStartGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartGame) (*sdk.Result, error) {
	game, err := k.GetGame(ctx, msg.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game does not exists")
	}
	if game.State != types.Waiting {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game is already start")
	}
	if game.Participant[0].Equals(msg.Creator) != true {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You're not host of this room")
	}

	k.StartGame(ctx, msg.ID, msg.CurrentTime, msg.UpdateTime)
	// temp <- struct {
	// 	*sdk.Result
	// 	error
	// }{&sdk.Result{Events: ctx.EventManager().Events()}, nil}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
