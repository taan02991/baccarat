package keeper

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/blockchain/baccarat/x/baccarat/helper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) CreateGame(ctx sdk.Context, game types.Game) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + game.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)
}

func (k Keeper) StartGame(ctx sdk.Context, id string, currentTime int64, updateTime int64) {
	var game types.Game
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
	if err != nil {
		fmt.Printf("Failed to start game:\n%s\n", err.Error())
		return
	}
	game.State = types.Playing
	deck := helper.GenerateDeck()
	hand, deck := helper.DrawCard(deck)
	helper.SetCache(id+"-deck", deck)
	helper.SetCache(id, hand)
	// Hash
	handHashed := sha256.Sum256([]byte(hand))
	game.ResultHash = append(game.ResultHash, string(handHashed[:]))
	game.CurrentTime = append(game.CurrentTime, currentTime)
	game.StartingTime = append(game.StartingTime, updateTime)
	game.BettingTime = append(game.BettingTime, updateTime+30)
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
	game.Bet[len(game.ResultHash)-1] = append(game.Bet[len(game.ResultHash)-1], bet)
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
    fmt.Printf("Failed to reveal result game:\n%s\n", err.Error())
    return
  }
  v, _ := helper.GetCache(id)
  game.Result = append(game.Result, v)
  winner, ratio := helper.Winner(v)
  for _, e := range game.Bet[len(game.ResultHash ) - 1] {
    reward, _ := strconv.Atoi(e.Amount.AmountOf("token").String())
    if e.Side == winner {
      coins, _ := sdk.ParseCoins(strconv.Itoa(reward * ratio) + "token")
      moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
      sdkError := k.CoinKeeper.SendCoins(ctx, moduleAcct, e.Creator, coins)
      if sdkError != nil {
        return
      }
      ctx.EventManager().EmitEvent(
        sdk.NewEvent(
          types.EventTypeRevealResult,
          sdk.NewAttribute(types.AttributeKeyGameID, id),
          sdk.NewAttribute(types.AttributeKeyWinner, string(winner)),
          sdk.NewAttribute(types.AttributeKeyAddress, e.Creator.String()),
          sdk.NewAttribute(types.AttributeKeyReward, strconv.Itoa(reward * ratio - reward)),
          sdk.NewAttribute(types.AttributeKeyBetSide, string(e.Side)),
          sdk.NewAttribute(types.AttributeKeyCard, string(v)),
          sdk.NewAttribute(types.AttributeKeyResultHash, game.ResultHash[len(game.ResultHash) - 1]),
        ),
      )
    } else {
      ctx.EventManager().EmitEvent(
        sdk.NewEvent(
          types.EventTypeRevealResult,
          sdk.NewAttribute(types.AttributeKeyGameID, id),
          sdk.NewAttribute(types.AttributeKeyWinner, string(winner)),
          sdk.NewAttribute(types.AttributeKeyAddress, e.Creator.String()),
          sdk.NewAttribute(types.AttributeKeyReward, strconv.Itoa(-reward)),
          sdk.NewAttribute(types.AttributeKeyBetSide, string(e.Side)),
          sdk.NewAttribute(types.AttributeKeyCard, string(v)),
          sdk.NewAttribute(types.AttributeKeyResultHash, game.ResultHash[len(game.ResultHash) - 1]),
        ),
      )
    }
  }
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)
}

func (k Keeper) AppendResultHash(ctx sdk.Context, id string) {
	var game types.Game
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
	if err != nil {
		fmt.Printf("Failed to append result hash:\n%s\n", err.Error())
		return
	}
	deck, _ := helper.GetCache(id + "-deck")
	hand, deck := helper.DrawCard(deck)
	helper.SetCache(id+"-deck", deck)
	helper.SetCache(id, hand)
	//Edit here
	handHashed := sha256.Sum256([]byte(hand))
	game.ResultHash = append(game.ResultHash, string(handHashed[:]))
	time := helper.GetTime("Betting")
	game.CurrentTime = append(game.CurrentTime, time.CurrentTime)
	game.StartingTime = append(game.StartingTime, time.CurrentTime)
	game.BettingTime = append(game.BettingTime, time.UpdateTime)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)
}

func (k Keeper) EndGame(ctx sdk.Context, id string) {
	var game types.Game
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &game)
	if err != nil {
		fmt.Printf("Failed to end game:\n%s\n", err.Error())
		return
	}
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

func getGame(ctx sdk.Context, k Keeper, id string) ([]byte, error) {
	game, err := k.GetGame(ctx, id)
	if err != nil {
		return nil, err
	}
	res, err := codec.MarshalJSONIndent(k.cdc, game)
	if err != nil {
		return nil, err
	}
	return res, nil
}
