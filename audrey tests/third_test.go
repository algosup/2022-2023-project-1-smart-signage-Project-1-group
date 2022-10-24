package main

import "testing"

func TestLight3(t *testing.T) {
	if  light == "ON"{
		got := {true,true,true,true,true}
		want :=  {false,false,false,false,false}
	}else if light == "OFF"{
		got := {false,false,false,false,false}
		want := {true,true,true,true,true}
	}else{
		return Error("")
	}
}

func main() {
	TestLight3()
}