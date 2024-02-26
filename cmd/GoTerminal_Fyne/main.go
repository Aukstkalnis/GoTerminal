package main

import (
	"flag"
	"fmt"
	"os"
)

// flag variables
var (
	port     string
	settings string
	logfile  string
)

func init() {
	flag.Parse()
}

// Version
var (
	Name    = "GeTerminal"
	Version = "0.0.1"
)

func main() {
	if len(os.Args) == 1 {
		run_gui_app()
		return
	}
	fmt.Printf("Welcome to %s v%s CLI!", Name, Version)
}
