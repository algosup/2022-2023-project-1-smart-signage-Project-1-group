package main

import (
	"fmt"
	"machine"
	"time"
)

// Creating a function to blink the led. As input we pass the pin number and the delay time
func blinkLed() {
	led := machine.A12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		time.Sleep(time.Second)
		led.Low()
		time.Sleep(time.Second)
	}
}

// Creating a function to Light On the LED.
func ledON() {
	led := machine.A12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.High()
}

// Creating a function to Turn Off the LED.
func ledOFF() {
	led := machine.A12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()
}

// Creating a function to print the value of the A7 pin
func printA7() {
	a7 := machine.A7
	a7.Configure(machine.PinConfig{Mode: machine.PinInput})
	for {
		fmt.Println(a7.Get())
		time.Sleep(time.Second)
	}
}

// Function that return the current time in the format: 2019-01-23 23:00:00.000000000 +0000 UTC m=+0.000000001
func getTime() string {
	return time.Now().String()
}

func main() {
	//blinkLed()
	//ledON()
	ledOFF()
	//printA7()
	fmt.Println(getTime())
}
