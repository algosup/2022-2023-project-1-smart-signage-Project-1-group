package test

import (
	"testing"
)

func TestLight1(t *testing.T) {
	got := light1(3)
	want := false
	if got != want {
		t.Error("LED's should be off")
	}
}
