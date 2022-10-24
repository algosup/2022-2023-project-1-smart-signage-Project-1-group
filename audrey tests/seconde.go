//see if a light is not working when it is supposed to be on 
package main

// a refaire 

var LEDS = [5]bool{true, true, true, false, true}

func Light2(LEDS){
	for c := 0; c < len(LEDS); c++ {
		got := true 
		if got != LEDS[c]{
			return Error ("une des LEDS est eteinte ")
			for c := 0; c < len(LEDS); c++ {
				LEDS[c]:= false
				return LEDS
		}
		else{
			return LEDS 
		}
	}
	}
}

func main() {
	light2()
}