package main

import (
	"machine"
	"strconv"
	"time"
)

var uart = machine.UART2 // Define the UART (UART is a serial communication protocol)

// Create a struct to define a new sensor
type Sensor struct {
	pin machine.Pin
}

type Led struct {
	pin machine.Pin
}

type ADC struct { // ADC is an analog to digital converter
	pin machine.Pin
}

// Function to create a new ADC and return it
func newADC(pin int) *ADC {
	adc := machine.Pin(pin)
	adc.Configure(machine.PinConfig{Mode: machine.PinInput})
	return &ADC{pin: adc}
}

// Function to get the value of the ADC
func (adc *ADC) getValue() string {
	return strconv.FormatBool(adc.pin.Get())
}

// Function to create a new led and return it
func newLed(pin int) *Led {
	led := &Led{pin: machine.Pin(pin)}
	led.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return led
}

// Function to turn on the led
func (led *Led) On() {
	led.pin.High()
}

// Function to turn off the led
func (led *Led) Off() {
	led.pin.Low()
}

// Function to create a new sensor and return it
func newSensor(pin int) *Sensor {
	sensor := machine.Pin(pin)
	sensor.Configure(machine.PinConfig{Mode: machine.PinInput})
	return &Sensor{pin: sensor}
}

// Write a command to the server
func WriteATCommand(command string) {
	uart.Write([]byte(command))

}

// Send a message to the server
func SendMessage(mess string) {
	var m = "\"" + mess + "\""
	uart.Write([]byte("AT+MSG=" + m + "\r\n"))
}

// Read the pin A12 (we actually getting a bool value) and returning it in a string
func (sensor *Sensor) getValue() string {
	return strconv.FormatBool(sensor.pin.Get())
}

func main() {

	sensor := machine.Pin(machine.PA7)                          // Create a new sensor (this one is the photoresistor)
	sensor.Configure(machine.PinConfig{Mode: machine.PinInput}) // Configure the sensor as an input

	// get the value of the sensor
	var value = sensor.Get()

	led := machine.Pin(machine.PA12)                          // Define the led pin
	led.Configure(machine.PinConfig{Mode: machine.PinOutput}) // Configure the pin as output

	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3}) // Configure the UART

	WriteATCommand("AT+ JOIN\r\n") // Join the network

	time.Sleep(time.Second * 10) //  We are waiting for the connection to be established
	SendMessage("Starting the program")
	time.Sleep(time.Second * 10) // We are waiting for the message to be sent

	for {
		if value == true { // However we wanted a int value instead of a bool
			led.Low() // Light off the led
			SendMessage("The sensor said true")
		} else {
			led.High() // Light on the led
			SendMessage("The sensor said false")
		}
	}
}
