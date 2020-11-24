package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBet{}

type MsgBet struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Side    BetSide        `json:"side" yaml:"side"`
	Amount  sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgBet(creator sdk.AccAddress, id string, side BetSide, amount sdk.Coins) MsgBet {
	return MsgBet{
		Creator: creator,
		ID:      id,
		Side:    side,
		Amount:  amount,
	}
}

func (msg MsgBet) Route() string {
	return RouterKey
}

func (msg MsgBet) Type() string {
	return "Bet"
}

func (msg MsgBet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBet) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
