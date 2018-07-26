package main

import (
	"go.bug.st/serial.v1"
)

func UI_SerialInit() {

}

func UI_SetDTR(state bool) {

}

func UI_SetRTS(state bool) {

}

func UI_ListPorts() []string {
	list, _ := serial.GetPortsList()
	return list
}
