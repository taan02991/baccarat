package keeper

import (
  // this line is used by starport scaffolding
	"github.com/blockchain/baccarat/x/baccarat/types"
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for baccarat clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListGame:
			return listGame(ctx, k)
		case types.QueryGetGame:
			return getGame(ctx, k, path[1])
		case types.QueryListUser:
			return listUser(ctx, k)
		case types.QueryGetUser:
			return getUser(ctx, k, path[1:])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown baccarat query endpoint")
		}
	}
}