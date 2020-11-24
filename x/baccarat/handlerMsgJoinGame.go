package baccarat

import (
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgJoinGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	gameId := msg.ID
	game, err := k.GetGame(ctx, gameId)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "No game id")
	}
	if len(game.Participant) == 4 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Room is full!")
	}
	game.Participant = append(game.Participant, msg.Creator)
	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
