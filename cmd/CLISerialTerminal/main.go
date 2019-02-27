package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Aukstkalnis/GoTerminal/terminal"
)

func main() {
	term, err := terminal.New(
		terminal.SetPort("COM12"),
	)
	if err != nil {
		log.Fatal("failed to create terminal:", err)
	}
	if err = term.Open(); err != nil {
		log.Fatal("failed to open port:", err)
	}
	for {
		<-time.After(10 * time.Second)
		fmt.Println("tick")
	}
}
