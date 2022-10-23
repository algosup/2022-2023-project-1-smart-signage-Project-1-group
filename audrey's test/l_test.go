package main

import "testing"

func TestLight1(t *testing.T,Time.endTime, Time.startTime, Times.actualTime) int8 {
	got := light1(Time.endTime, Time.startTime, Times.actualTime)
	if got != 1 {
		return t.Error("la led est éteinte")
	}
}
