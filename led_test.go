package main

import (
	"Time"
)

atm := Time.Now()

func Now() Time


func Timecheck(atm float64) bool {
	if atm <1.00 && atm >6.00{
		return false
	}
	return true
}



