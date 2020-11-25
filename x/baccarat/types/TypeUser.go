package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type User struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Addr sdk.AccAddress `json:"addr" yaml:"addr"`
}

type GetUsersRequest struct {
	Addr []string `json:"addr"`	
}