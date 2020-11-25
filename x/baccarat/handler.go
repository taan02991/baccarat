package baccarat

import (
	"fmt"

	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// wg.Add(1)
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding
		case types.MsgCreateGame:
			return handleMsgCreateGame(ctx, k, msg)
		case types.MsgStartGame:

			// temp := make(chan struct {
			// 	*sdk.Result
			// 	error
			// }, 1)
			// go func(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartGame) {
			// 	defer wg.Done()
			// 	go handleMsgStartGame(ctx, k, msg, temp)

			// }(ctx, k, msg)
			// sdkOut := <-temp
			// return sdkOut.Result, sdkOut.error
			return handleMsgStartGame(ctx, k, msg)
		case types.MsgBet:
			return handleMsgBet(ctx, k, msg)
		case types.MsgEditParticipant:
			return handleMsgEditParticipant(ctx, k, msg)
		case types.MsgCreateUser:
			return handleMsgCreateUser(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
