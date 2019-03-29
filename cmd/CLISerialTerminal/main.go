package main

import (
	"log"

	"github.com/Aukstkalnis/GoTerminal/terminal"
)

func main() {
	term, err := terminal.New(
		terminal.SetPort("COM14"),
	)
	if err != nil {
		log.Fatal("failed to create terminal:", err)
	}
	if err = term.Open(); err != nil {
		log.Fatal("failed to open port:", err)
	}
}
