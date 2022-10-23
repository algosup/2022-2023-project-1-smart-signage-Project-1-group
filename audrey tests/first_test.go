package main

import "testing"

func TestLight1(t *testing.T) {
	got := light1()
	if got != 1 {
		return t.Error("LED's are off")
	}
}
