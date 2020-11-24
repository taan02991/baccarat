package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEditParticipant{}

type MsgEditParticipant struct {
	Creator sdk.AccAddress    `json:"creator" yaml:"creator"`
	ID      string            `json:"id" yaml:"id"`
	Action  ParticipantAction `json:"action" yaml:"action"`
}

func NewMsgEditParticipant(creator sdk.AccAddress, id string, action ParticipantAction) MsgEditParticipant {
	return MsgEditParticipant{
		Creator: creator,
		ID:      id,
		Action:  action,
	}
}

func (msg MsgEditParticipant) Route() string {
	return RouterKey
}

func (msg MsgEditParticipant) Type() string {
	return "EditParticipant"
}

func (msg MsgEditParticipant) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgEditParticipant) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgEditParticipant) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
