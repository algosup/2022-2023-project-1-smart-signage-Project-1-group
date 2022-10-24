// on / off  for ligths

package main

var LEDS = [5]bool{true, true, true, true, true}
var light int 

// a refaire 

func Light3(LEDS) {
	if light == "ON" {
		for c := 0; c < len(LEDS); c++ {
			LEDS[c] = false
		}
		light := "OFF"
		return LEDS
	} else if light == "OFF" {
		for c := 0; c < len(LEDS); c++ {
			LEDS[c] = true
		}
		light := "ON"
		return LEDS
	} else {
		return Error("certaines LEDS sont allumées et d'autres éteintes voir état des leds  ")
	}

}

func main() {
	light3()
}
