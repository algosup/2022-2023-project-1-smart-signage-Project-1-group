package main

import "testing"

func Testlight1(t *testing.T) {
	got := light1(3)
	want := true
	if got != want {
		return Error("LED's should be off")
	}
}

// // TestLight2 is a test function for light2
// func TestLight2(t *testing.T) {
// 	got := rand.Intn(10, 100),
// 	want := 10
// 	for c := 0; c < len(LEDS), c++ {
// 		if got[c]<= wan {
// 			return Error ("LED's are off")
// 		}
// 	}
// }

// // TestLight3 is a test function for light3
// func TestLight3(t *testing.T) {
// 	if  light == 10 {
// 		return 0
// 	}else if light == light >10{
// 		return light "Power of light"
// 	}else if light < 10 {
// 		return Error("There is a problem with a LED")
// 	}
// }

// func main() {
// 	TestLight1()
// 	TestLight2()
// 	TestLight3()
// }
