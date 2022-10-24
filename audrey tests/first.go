package main

var actualtime float32 = 15.58

// var startTime float32 = 8.00
// var endTime float32 = 19.30

func light1() float32 {
	if actualtime > 6.00 {
		if actualtime < 1.00 {
			return 1
		}
	}
	return 0
}

func main() {
	light1()
}
