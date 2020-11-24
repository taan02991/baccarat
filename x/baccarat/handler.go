package baccarat

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding
		case types.MsgCreateGame:
			return handleMsgCreateGame(ctx, k, msg)
<<<<<<< HEAD
=======
		case types.MsgStartGame:
			return handleMsgStartGame(ctx, k, msg)
		case types.MsgBet:
			return handleMsgBet(ctx, k, msg)
		case types.MsgEditParticipant:
			return handleMsgEditParticipant(ctx, k, msg)
		case types.MsgCreateUser:
			return handleMsgCreateUser(ctx, k, msg)
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
