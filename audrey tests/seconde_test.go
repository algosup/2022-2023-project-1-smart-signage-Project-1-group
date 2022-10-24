package main

import "testing"
import "fmt"
import "math/rand"

func TestLight2(t *testing.T) {
	got := rand.Intn(10, 100),
	want := 10
	for c := 0; c < len(LEDS); c++ {
		if got[c]<= wan {
			return Error ("leds off ")
		}
	}
}

func main() {
	TestLight2()
}