package function

import (
	"testing"
)

func TestLightstatus(t *testing.T) {
	got := lightstatus(3)
	want := false
	if got != want {
		t.Error("LED's should be off")
	}
}

func TestLightsvoltage(t *testing.T) {
	got := lightsvoltage(230)
	want := true
	if got != want {
		t.Error(" a LED should be broken ")
	}
}

func TestLightsonoff(t *testing.T) {
	got := lightsonoff(230)
	want := 0
	if got != want {
		t.Error(" a LED should be broken ")
	}
}
