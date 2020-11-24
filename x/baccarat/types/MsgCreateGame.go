package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
<<<<<<< HEAD
	"github.com/google/uuid"
=======
  "github.com/google/uuid"
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
)

var _ sdk.Msg = &MsgCreateGame{}

type MsgCreateGame struct {
<<<<<<< HEAD
	ID          string
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	RoomNo      string         `json:"roomNo" yaml:"roomNo"`
	Round       string         `json:"round" yaml:"round"`
	Participant string         `json:"participant" yaml:"participant"`
	Result      string         `json:"result" yaml:"result"`
	ResultHash  string         `json:"resultHash" yaml:"resultHash"`
}

func NewMsgCreateGame(creator sdk.AccAddress, roomNo string, round string, participant string, result string, resultHash string) MsgCreateGame {
	return MsgCreateGame{
		ID:          uuid.New().String(),
		Creator:     creator,
		RoomNo:      roomNo,
		Round:       round,
		Participant: participant,
		Result:      result,
		ResultHash:  resultHash,
=======
  Creator     sdk.AccAddress    `json:"creator" yaml:"creator"`
	ID          string            `json:"id" yaml:"id"`
}

func NewMsgCreateGame(creator sdk.AccAddress) MsgCreateGame {
  return MsgCreateGame{
    ID: uuid.New().String(),
		Creator: creator,
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
	}
}

func (msg MsgCreateGame) Route() string {
<<<<<<< HEAD
	return RouterKey
}

func (msg MsgCreateGame) Type() string {
	return "CreateGame"
}

func (msg MsgCreateGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
=======
  return RouterKey
}

func (msg MsgCreateGame) Type() string {
  return "CreateGame"
}

func (msg MsgCreateGame) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGame) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
