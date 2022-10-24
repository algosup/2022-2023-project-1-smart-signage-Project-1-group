package main

import "testing"

func TestLight2(t *testing.T) {
	got := {false,false,true,true,false}
	want := {true,true,true,true,true}
	for c := 0; c < len(LEDS); c++ {
		if got[c]!= want[c]{
			return Error ("leds éteinte ")
		}
	}
}

func main() {
	TestLight2()
}