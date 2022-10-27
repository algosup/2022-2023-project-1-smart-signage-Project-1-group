package function

import (
	"testing"
)

// This test uses the function lightstatus
func TestLightstatus(t *testing.T) {
	got := lightstatus(3)
	want := false
	if got != want {
		t.Error("LED's should be off")
	}
}

// This test uses the function lightvoltage
func TestLightsvoltage(t *testing.T) {
	got := lightsvoltage(230)
	want := true
	if got != want {
		t.Error(" a LED should be broken ")
	}
}

// This test uses the function lightsonoff
func TestLightsonoff(t *testing.T) {
	got := lightsonoff(230)
	want := 0
	if got != want {
		t.Error(" a LED should be broken ")
	}
}
