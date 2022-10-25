package function

type hour struct {
	hourvar int
}

func light1(hh int) bool {
	h := hour{hh}
	if h.hourvar >= 6 && h.hourvar <= 1 {
		return true
	}
	return false
}

// // This function is used to see if a light is not working when it's supposed to be on

// var LEDS = [5]int{11,110,0,10,10}
func light2(LEDS int) bool {
	for c := 0; c < len(LEDS); c++ {
		d := LEDS[c]
		t:= 10
		if  t >= d || d <= 0 {
			return false
		}
		return true 
}

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
