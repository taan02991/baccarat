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
  //Todo: generate card deck & draw first 4 cards and save hash in cache
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

func (k Keeper) EditParticipant(ctx sdk.Context, id string, creator sdk.AccAddress, action types.ParticipantAction) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to edit participant:\n%s\n", err.Error())
    return
  }
  if action == types.Join && len(game.Participant) < 6 {
      for _, e := range game.Participant {
        if e.Equals(creator) == true {
          return
        }
        game.Participant = append(game.Participant, creator)
      }
  } else if action == types.Leave {
      for i, e := range game.Participant {
        if e.Equals(creator) == true {
          game.Participant = append(game.Participant[:i], game.Participant[i+1:]...)
          if len(game.Participant) == 0 {
            game.State = types.End
          }
        }
      }
  } else {
    fmt.Printf("Edit participant action doesn't match\n")
    return
  }
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
  store.Set(key, value)
}

func (k Keeper) RevealResult(ctx sdk.Context, id string) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to start game:\n%s\n", err.Error())
    return
  }
  //Todo: get card from cahce
  game.Result = append(game.Result, "1A,2B;4S,5K")
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
  store.Set(key, value)
}

func (k Keeper) AppendResultHash(ctx sdk.Context, id string) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to start game:\n%s\n", err.Error())
    return
  }

  game.ResultHash = append(game.ResultHash, "1A,2B;4S,5K")
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
  store.Set(key, value)
}

func (k Keeper) EndGame(ctx sdk.Context, id string) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to start game:\n%s\n", err.Error())
    return
  }
  //Todo: draw 4 cards & save hash in cache
  game.State = types.End
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
  store.Set(key, value)
}

func (k Keeper) GetGame(ctx sdk.Context, id string) (types.Game, error) {
  var game types.Game
	store := ctx.KVStore(k.storeKey)
  key := []byte(types.GamePrefix + id)
  err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
  if err != nil {
    fmt.Printf("Failed to get game:\n%s\n", err.Error())
    return game, err
  }
  return game, nil
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