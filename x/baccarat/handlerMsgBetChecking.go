package baccarat

import (
	"time"

	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgBetChecking(ctx sdk.Context, k keeper.Keeper, msg types.MsgBet) (*sdk.Result, error) {
	game, err := k.GetGame(ctx, msg.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game does not exists")
	}
	if game.State != types.Playing {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game is not Playing")
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	if game.Creator.String() != moduleAcct.String() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You are not host.")
	}
	var timeUP bool
	if time.Now().Unix() >= game.BettingTime[len(game.BettingTime)-1] {
		timeUP = true
	}

	if timeUP {
		k.RevealResult(ctx, msg.ID)
		if len(game.ResultHash) < 8 {
			k.AppendResultHash(ctx, msg.ID)
		} else {
			k.EndGame(ctx, msg.ID)
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
