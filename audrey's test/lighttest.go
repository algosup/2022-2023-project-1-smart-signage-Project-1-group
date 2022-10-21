package main

import "testing"

func TestLight(t *testing.T) {
	got := string(light1(8.00, 20.00, 7.01))
	want := "the light is on "
	if got != want {
		t.Errorf("You need to switch on the light")
	}
}
