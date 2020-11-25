package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/types"
  "github.com/cosmos/cosmos-sdk/codec"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
  "fmt"
  abci "github.com/tendermint/tendermint/abci/types"
)

var INITIAL_AMOUNT, _ = sdk.ParseCoins("1000token")

func (k Keeper) CreateUser(ctx sdk.Context, user types.User) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.UserPrefix + user.Addr.String())
	value := k.cdc.MustMarshalBinaryLengthPrefixed(user)
  store.Set(key, value)
  sdkError := k.CoinKeeper.SetCoins(ctx, user.Addr, INITIAL_AMOUNT)
  if sdkError != nil {
		fmt.Printf("could not set initial coin\n%s\n", sdkError.Error())
	}
}

func (k Keeper) GetUser(ctx sdk.Context, addr string) (types.User, error) {
  var user types.User
  store := ctx.KVStore(k.storeKey)
  key := []byte(types.UserPrefix + addr)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &user)
  if err != nil {
      return user, err
  }
  return user, nil
}

func listUser(ctx sdk.Context, k Keeper) ([]byte, error) {
  var userList []types.User
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var user types.User
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &user)
    userList = append(userList, user)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, userList)
  return res, nil
}

func getUser(ctx sdk.Context, k Keeper, path []string) ([]byte, error) {
  user, err := k.GetUser(ctx, path[0])
  if err != nil {
    return nil, err
	}
  res, err := codec.MarshalJSONIndent(k.cdc, user)
  if err != nil {
    return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func getUsers(ctx sdk.Context, k Keeper, req abci.RequestQuery) ([]byte, error) {
  userMap := make(map[string]string)
  var data types.GetUsersRequest
  k.cdc.MustUnmarshalJSON(req.Data, &data)
  for _, addr := range data.Addr {
    user, err := k.GetUser(ctx, addr)
    if err != nil {
      return nil, err
    }
    userMap[addr] = user.Name
  }
  res, err := codec.MarshalJSONIndent(k.cdc, userMap)
  if err != nil {
    return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}