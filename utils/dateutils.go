package utils

import (
	"time"
)

func CurDateTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05.000")
}
