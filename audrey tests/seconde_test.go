package main

// a refaire 

import "testing"

func TestLight2(t *testing.T) {
	got := 
	want := 
	for c := 0; c < len(LEDS); c++ {
		if got[c]!= want[c]{
			return Error ("leds éteinte ")
		}
	}
}

func main() {
	TestLight2()
}