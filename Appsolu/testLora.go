package main

import (
	"machine"
	"time"
)

// I want to blink the lek a certain amount of time.
func BlinkLedXTimes(times int) {
	led := machine.Pin(12)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for i := 0; i < times; i++ {
		led.High()
		time.Sleep(100 * time.Millisecond)
		led.Low()
		time.Sleep(100 * time.Millisecond)
	}
}

func WriteATCommand(command string) {
	machine.UART1.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A3, RX: machine.A2}) // Configure the UART for a baud rate of 9600, with TX on pin A3 and RX on pin A2
	_, err := machine.UART1.Write([]byte(command))                                              // Write the command to the bluepill
	if err != nil {                                                                             // If an error occurs we print it
		println("Error: " + err.Error())
	}
}

func ReadDataSerial() {
	for {
		if machine.UART1.Buffered() > 0 { // If data is received
			data := make([]byte, machine.UART1.Buffered()) // Create a slice of bytes with the size of the data received
			machine.UART1.Read(data)                       // Read the data received and put it in the slice
			println(string(data))                          // Print the data received                            // Blink the led 3 times
			break
		} else { // If no data is received
			BlinkLedXTimes(1) // Blink the led 1 time
			break
		}
	}
}

// Printing data received by the bluepill on the serial port. If data is received, the led is blinking 3 times. If not, the led is blinking 1 time.
func main() {
	time.Sleep(5000)
	WriteATCommand("AT+VER\r\n")
	for {
		ReadDataSerial()
	}
}
