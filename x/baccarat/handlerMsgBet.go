package baccarat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/blockchain/baccarat/x/baccarat/types"
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgBet(ctx sdk.Context, k keeper.Keeper, msg types.MsgBet) (*sdk.Result, error) {
	var bet = types.Bet{
		Creator:	msg.Creator,
		Side:		msg.Side,
		Amount:		msg.Amount,
	}
	game, err := k.GetGame(ctx, msg.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game does not exists")
	}
	for _, e := range game.Bet[len(game.ResultHash) - 1] {
		if e.Creator.Equals(bet.Creator) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You already bet")
		}
	}

	k.Bet(ctx, msg.ID, bet)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, bet.Creator, moduleAcct, bet.Amount)
	if sdkError != nil {
		return nil, sdkError
	}

	game, err = k.GetGame(ctx, msg.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game does not exists")
	}

	var isAllBet = true
	for _, p := range game.Participant {
		var isBet = false
		for _, e := range game.Bet[len(game.ResultHash) - 1] {
			if e.Creator.Equals(p) {
				isBet = true
				break
			}
		}
		if isBet == false {
			isAllBet = false
			break
		}
	}

	if isAllBet {
		k.RevealResult(ctx, msg.ID)
		if len(game.ResultHash) < 8 {
			k.AppendResultHash(ctx, msg.ID)
		} else {
			k.EndGame(ctx, msg.ID)
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
