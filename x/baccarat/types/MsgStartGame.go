package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgStartGame{}

type MsgStartGame struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	ID          string         `json:"id" yaml:"id"`
	CurrentTime int64          `json:"current_time"`
	UpdateTime  int64          `json:"update_time"`
}

func NewMsgStartGame(creator sdk.AccAddress, id string, currentTime int64, updateTime int64) MsgStartGame {
	return MsgStartGame{
		Creator:     creator,
		ID:          id,
		CurrentTime: currentTime,
		UpdateTime:  updateTime,
	}
}

func (msg MsgStartGame) Route() string {
	return RouterKey
}

func (msg MsgStartGame) Type() string {
	return "StartGame"
}

func (msg MsgStartGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgStartGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgStartGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
