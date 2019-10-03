package main

import (
	"flag"
	"log"
	"os"

	"github.com/Aukstkalnis/GoTerminal/terminal"
)

// RTS control
// DTR control

type config struct {
	port string
	rts  bool
	dtr  bool
}

var cfg config

func init() {
	flag.StringVar(&cfg.port, "port", "", "select port")
	flag.BoolVar(&cfg.rts, "rts", false, "set RTS")
	flag.BoolVar(&cfg.rts, "dtr", false, "set DTR")
	flag.Parse()
}

func main() {
	log.Printf("Args: port:%s, rts:%t, dtr:%t\n", cfg.port, cfg.rts, cfg.dtr)
	term, err := terminal.New(
		terminal.SetPort(cfg.port),
		terminal.SetStartupDTR(cfg.dtr),
		terminal.SetStartupRTS(cfg.rts),
	)
	if err != nil {
		log.Fatal("failed to create :", err)
	}
	if err = term.Open(); err != nil {
		log.Fatal("failed to open port:", err)
	}
	var (
		buf = make([]byte, 2048)
		n   int
	)
	logfile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal("failed to open log file:", err)
	}
	defer logfile.Close()
	for {
		if n, err = term.Read(buf); err != nil {
			log.Fatal("failed to read:", err)
		}
		_, err = logfile.Write(buf[:n])
		if err != nil {
			log.Fatal("failed to write to file:", err)
		}
		if n > 0 {
			for i := 0; i < n; i++ {
				if buf[i] == '\r' {
					buf[i] = '\n'
				}
			}
			if n, err = os.Stdout.Write(buf[:n]); err != nil {
				log.Fatal("failed to write ro stdout:", err)
			}
		}
	}
}
