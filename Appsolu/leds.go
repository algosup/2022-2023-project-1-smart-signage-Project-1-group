package main

import (
	"fmt"
	"machine"
	"time"
)

type LED struct { // Creating a new type LED
	pin machine.Pin
}

// Create a struct to define a new sensor
type Sensor struct {
	pin machine.Pin
}

// Function to create a new LED
func NewLED(pin int) *LED {
	led := machine.Pin(pin)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &LED{pin: led}
}

// Function to turn on the LED
func (led *LED) On() {
	led.pin.High()
}

// Function to turn off the LED
func (led *LED) Off() {
	led.pin.Low()
}

// Function to blink the LED
func (led *LED) Blink(delay time.Duration) {
	led.On()
	time.Sleep(delay)
	led.Off()
	time.Sleep(delay)
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
	led := NewLED(12)
	for {
		led.Blink(time.Second)
	}

	machine.UART1.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A3, RX: machine.A2})
	machine.UART1.Write([]byte("AT+VERSION\r\n"))
	time.Sleep(time.Second)

}
