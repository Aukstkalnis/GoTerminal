package main

import (
	"log"

	"github.com/Aukstkalnis/GoTerminal/terminal"
)

func main() {
	term, err := terminal.New(
		terminal.SetPort("COM12"),
	)
	if err != nil {
		log.Fatal("feiled to create terminal:", err)
	}
	if err = term.Open(); err != nil {
		log.Fatal("failed to open port:", err)
	}
	term.Log()
}
