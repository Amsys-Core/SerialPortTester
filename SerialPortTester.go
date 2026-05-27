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

	for true {
		var inputCommand string
		fmt.Scanln(&inputCommand)
		if inputCommand == "SetDTROn" {
			controlPanelPortName.SetDTR(true)
		} else if inputCommand == "SetDTROff" {
			controlPanelPortName.SetDTR(false)
		} else if inputCommand == "ReadBodyValue" {
			_, err := bottleBodyPortName.Write([]byte("GA00\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			for {
				ReadPortOutput(bottleBodyPortName, buff)
			}
		} else if inputCommand == "Help" {
			println("SetDTROn")
			println("SetDTROff")
			println("ReadBodyValue")
			println("ReadBottomValue")
			println("Help")
			println("Break")
		} else if inputCommand == "Break" {
			break
		} else if inputCommand == "ReadBottomValue" {
			_, err := bottleBottomPortName.Write([]byte("GA00\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			for {
				ReadPortOutput(bottleBottomPortName, buff)
			}
		}
	}
}

func ReadPortOutput(bottleBodyPortName serial.Port, buff []byte) {
	n, err := bottleBodyPortName.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	println(string(buff[:n]))
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
