package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBetChecking{}

type MsgBetChecking struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
}

func NewMsgBetChecking(creator sdk.AccAddress, id string) MsgBetChecking {
	return MsgBetChecking{
		Creator: creator,
		ID:      id,
	}
}

func (msg MsgBetChecking) Route() string {
	return RouterKey
}

func (msg MsgBetChecking) Type() string {
	return "MsgBetChecking"
}

func (msg MsgBetChecking) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgBetChecking) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBetChecking) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
