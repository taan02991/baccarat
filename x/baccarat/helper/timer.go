package helper

import (
	"time"
)

func getTimeToStart() (int64, int64) {
	currentTime := time.Now().Unix()
	return currentTime, currentTime + 15
}
