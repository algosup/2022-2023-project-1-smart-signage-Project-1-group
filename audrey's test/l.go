package main

import "time"

var T float32 = time.Now()
var startTime float32 = 8
var endTime float32 = 19.30

func light1(c Times) int8 {
	if c.actualTime > 6.00 {
		return 1
	} else {
		return 0
	}
}
