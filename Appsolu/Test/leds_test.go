package main

import (
	"reflect"
	"testing"
)

func TestNewLED(t *testing.T) {
	type args struct {
		pin int
	}
	tests := []struct {
		name string
		args args
		want *LED
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLED(tt.args.pin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLED() = %v, want %v", got, tt.want)
			}
		})
	}
}
