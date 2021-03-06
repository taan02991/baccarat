package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding
		cdc.RegisterConcrete(MsgCreateGame{}, "baccarat/CreateGame", nil)
		cdc.RegisterConcrete(MsgStartGame{}, "baccarat/StartGame", nil)
		cdc.RegisterConcrete(MsgBet{}, "baccarat/Bet", nil)
		cdc.RegisterConcrete(MsgCreateUser{}, "baccarat/CreateUser", nil)
		cdc.RegisterConcrete(MsgEditParticipant{}, "baccarat/EditParticipant", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
