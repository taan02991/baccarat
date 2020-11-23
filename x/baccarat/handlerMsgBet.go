package baccarat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	k.Bet(ctx, msg.ID, bet)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	sdkError := k.CoinKeeper.SendCoins(ctx, bet.Creator, moduleAcct, bet.Amount)
	if sdkError != nil {
		return nil, sdkError
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
