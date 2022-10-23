package main

import (
	"testing"
)

func TestTimecheck(t *testing.T) {

	Timecheck(9.00)

	want := true

	if got != want {
		t.Error("got %d want %d", got, want)
		t.Errorf("got %d want %d", got, want)
	}
}
