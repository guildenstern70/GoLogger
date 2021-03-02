/*
 * Project GoLogger
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package utils

import (
	"time"
)

func CurDateTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05.000")
}
