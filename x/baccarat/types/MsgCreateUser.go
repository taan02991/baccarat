package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateUser{}

type MsgCreateUser struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	Addr    sdk.AccAddress `json:"addr" yaml:"addr"`
}

func NewMsgCreateUser(creator sdk.AccAddress, name string, addr sdk.AccAddress) MsgCreateUser {
	return MsgCreateUser{
		ID:      uuid.New().String(),
		Creator: creator,
		Name:    name,
		Addr:    addr,
	}
}

func (msg MsgCreateUser) Route() string {
	return RouterKey
}

func (msg MsgCreateUser) Type() string {
	return "CreateUser"
}

func (msg MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
