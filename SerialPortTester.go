package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
	} else {
		for _, port := range ports {
			fmt.Printf("Found port: %v\n", port)
		}
	}

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	buff := make([]byte, 100)

	bottleBodyPortName, err := serial.Open(setCOMPort(ports, "Bottle Body"), mode)
	bottleBottomPortName, err := serial.Open(setCOMPort(ports, "Bottle Bottom"), mode)
	controlPanelPortName, err := serial.Open(setCOMPort(ports, "Control Panel"), mode)

	if err != nil {
		log.Fatal(err)
	}


func setCOMPort(ports []string, section string) string {
	println("Please set the " + section + " com port")
	var inputPortName string
	_, err := fmt.Scanln(&inputPortName)
	if err != nil {
		log.Fatal(err)
	}
	for _, port := range ports {
		if port == inputPortName {
			fmt.Printf("Found port: %v\n", port)
			return port
		}
	}
	return ""
}
