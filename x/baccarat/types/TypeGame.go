package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
	Creator      sdk.AccAddress   `json:"creator" yaml:"creator"`
	ID           string           `json:"id" yaml:"id"`
	State        int              `json:"state" yaml:"state"`
	Participant  []sdk.AccAddress `json:"participant" yaml:"participant"`
	Stake        []Stake          `json:"cardResource"`
	CardResource []string         `json:"cardResource"`
	Result       []Card           `json:"result" yaml:"result"`
	ResultHash   []string         `json:"resultHash" yaml:"resultHash"`
}
