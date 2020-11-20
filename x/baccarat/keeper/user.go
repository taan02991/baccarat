package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateUser(ctx sdk.Context, user types.User) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.UserPrefix + user.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(user)
	store.Set(key, value)
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