package main

import (
	"machine"
	"strconv"
	"time"
)

var uart = machine.UART2

// Create a struct to define a new sensor
type Sensor struct {
	pin machine.Pin
}

type Led struct {
	pin machine.Pin
}

type ADC struct {
	pin machine.Pin
}

func newADC(pin int) *ADC {
	adc := machine.Pin(pin)
	adc.Configure(machine.PinConfig{Mode: machine.PinInput})
	return &ADC{pin: adc}
}

func (adc *ADC) getValue() string {
	return strconv.FormatBool(adc.pin.Get())
}

func newLed(pin int) *Led {
	led := &Led{pin: machine.Pin(pin)}
	led.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return led
}

func (led *Led) On() {
	led.pin.High()
}

func (led *Led) Off() {
	led.pin.Low()
}

// Function to create a new sensor
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

// read the pin A12 and send the data to the server
func (sensor *Sensor) getValue() string {
	return strconv.FormatBool(sensor.pin.Get())
}

func main() {

	sensor := machine.Pin(machine.PA7)
	sensor.Configure(machine.PinConfig{Mode: machine.PinInput})

	// get the value of the sensor
	var value = sensor.Get()

	led := machine.Pin(machine.PA12)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3})

	WriteATCommand("AT+ JOIN\r\n")

	time.Sleep(time.Second * 10)
	SendMessage("Starting the program")
	time.Sleep(time.Second * 10)

	for {
		if value == true { // However we wanted a int value instead of a bool
			led.Low()
			SendMessage("The sensor said true")
		} else {
			led.High()
			SendMessage("The sensor said false")
		}
	}
}
