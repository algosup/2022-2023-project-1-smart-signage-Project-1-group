package main

func light1(hourvar int) bool {
	if hourvar >= 6 && hourvar <= 1 {
		return true
	}
	return false
}

// // This function is used to see if a light is not working when it's supposed to be on

// var LEDS = [5]bool{true, true, true, false, true}

// func light2(LEDS){
// 	for c := 0; c < len(LEDS); c++ {
// 		got := true
// 		if got != LEDS[c]{
// 			return Error ("A LED is not working")
// 			for c := 0; c < len(LEDS); c++ {
// 				LEDS[c]:= false
// 				return LEDS
// 		}
// 		else{
// 			return LEDS
// 		}
// 	}
// 	}
// }

// // This function check if lights are on or off

// var LEDS = [5]bool{true, true, true, true, true}
// var light int

// func light3(LEDS) {
// 	if light == "ON" {
// 		for c := 0; c < len(LEDS); c++ {
// 			LEDS[c] = false
// 		}
// 		light := "OFF"
// 		return LEDS
// 	} else if light == "OFF" {
// 		for c := 0; c < len(LEDS); c++ {
// 			LEDS[c] = true
// 		}
// 		light := "ON"
// 		return LEDS
// 	} else {
// 		return Error("Some LEDS are on and others are off see LEDS state ")
// 	}

// }

// func main() {
// 	light1()
// 	light2()
// 	light3()
// }
