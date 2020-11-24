package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
<<<<<<< HEAD
	Creator      sdk.AccAddress   `json:"creator" yaml:"creator"`
	ID           string           `json:"id" yaml:"id"`
	State        int              `json:"state" yaml:"state"`
	Participant  []sdk.AccAddress `json:"participant" yaml:"participant"`
	Stake        []Stake          `json:"cardResource"`
	CardResource []string         `json:"cardResource"`
	Result       []Card           `json:"result" yaml:"result"`
	ResultHash   []string         `json:"resultHash" yaml:"resultHash"`
}
=======
	Creator     sdk.AccAddress    `json:"creator" yaml:"creator"`
	ID          string            `json:"id" yaml:"id"`
  State       StateType         `json:"state" yaml:"state"`
  Participant []sdk.AccAddress  `json:"participant" yaml:"participant"`
  Result      []string            `json:"result" yaml:"result"`
  ResultHash  []string            `json:"resultHash" yaml:"resultHash"`
  Bet         [][]Bet            `json:"bet" yaml:"bet"`
}

type Bet struct {
  Creator sdk.AccAddress  `json:"creator" yaml:"creator"`
  Side    BetSide         `json:"side" yaml:"side"`
  Amount  sdk.Coins             `json:"amount" yaml:"amount"`
}

type StateType string

const (
  Waiting StateType = "Waiting"
  Playing           = "Playing"
  End               = "End"
)


type BetSide string

const (
  Player BetSide = "Player"
  Banker         = "Banker"
  Tie            = "Tie"
)

type ParticipantAction string

const (
  Join ParticipantAction  = "Join"
  Leave                   = "Leave"
)
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
