package types

type Card struct {
	Player1 string `json:"Player1" yaml:"roomNo"`
	Player2 string `json:"Player2" yaml:"roomNo"`
	Banker1 string `json:"Banker1" yaml:"roomNo"`
	Banker2 string `json:"Banker2" yaml:"roomNo"`
}
