package helper

import (
	"time"
)

type Timer struct {
	CurrentTime int64 `json:"current_time" yaml:"current_time"`
	UpdateTime  int64 `json:"update_time" yaml:"update_time"`
}

func GetTime(State string) Timer {
	currentTime := time.Now().Unix()
	var outTime Timer
	if State == "start" {
		outTime.CurrentTime = currentTime
		outTime.UpdateTime = currentTime + 15
	} else if State == "betting" {
		outTime.CurrentTime = currentTime
		outTime.UpdateTime = currentTime + 30
	}
	return outTime
}
