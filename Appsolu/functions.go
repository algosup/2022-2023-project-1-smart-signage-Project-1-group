package main

import (
	"fmt"
	"machine"
	"time"
)

// Create a struct to define a new sensor
type Sensor struct {
	pin machine.Pin
}

// Function to create a new sensor
func newSensor(pin int) *Sensor {
	sensor := machine.Pin(pin)
	sensor.Configure(machine.PinConfig{Mode: machine.PinInput})
	return &Sensor{pin: sensor}
}

// Function to get the value of the sensor
func (sensor *Sensor) getValue() bool {
	return sensor.pin.Get()
}

// Function to print the value of the sensor
func (sensor *Sensor) printValue() {
	fmt.Println(sensor.getValue())
}

// Function to print the current time
func printTime() {
	fmt.Println(time.Now())
}

func main() {
	sensor := newSensor(7)
	for {
		sensor.printValue()
		printTime()
		time.Sleep(time.Second)
	}
}
