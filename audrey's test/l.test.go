package main

import "testing"

func main (){
	
}
func TestLight1(t *testing.T, T) int8 {
	got := light1(Times.actualTime)
	if got != 1 {
		return t.Error("la led est éteinte")
	}
}
