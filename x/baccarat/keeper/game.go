package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/types"
  "github.com/cosmos/cosmos-sdk/codec"
  "fmt"
)

func (k Keeper) CreateGame(ctx sdk.Context, game types.Game) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + game.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)
}

func (k Keeper) StartGame(ctx sdk.Context, id string) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to start game:\n%s\n", err.Error())
    return
  }
  game.State = types.Playing
  //Todo: generate card deck & draw first 4 cards
  game.ResultHash = append(game.ResultHash, "1A,2B;4S,5K")
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
  store.Set(key, value)
}

func (k Keeper) Bet(ctx sdk.Context, id string, bet types.Bet) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to bet:\n%s\n", err.Error())
    return
  }
  game.Bet[len(game.ResultHash ) - 1] = append(game.Bet[len(game.ResultHash) - 1], bet)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
  store.Set(key, value)
}

func listGame(ctx sdk.Context, k Keeper) ([]byte, error) {
  var gameList []types.Game
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
  for ; iterator.Valid(); iterator.Next() {
    var game types.Game
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &game)
    gameList = append(gameList, game)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, gameList)
  return res, nil
}