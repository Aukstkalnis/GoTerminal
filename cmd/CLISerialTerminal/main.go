package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Aukstkalnis/GoTerminal/terminal"
	"github.com/containerd/console"
)

// RTS control
// DTR control

type config struct {
	port string
	rts  bool
	dtr  bool
}

const traceFileName = "log.txt"

var cfg config

func init() {
	flag.StringVar(&cfg.port, "port", "", "select port")
	flag.BoolVar(&cfg.rts, "rts", false, "set RTS")
	flag.BoolVar(&cfg.rts, "dtr", false, "set DTR")
	flag.Parse()
}

func main() {
	traceFile, err := os.OpenFile(traceFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("failed to open trace file:", err)
	}
	defer traceFile.Close()
	log.SetOutput(traceFile)

	current := console.Current()
	defer current.Reset()

	if err := current.SetRaw(); err != nil {
		log.Fatal(err)
	}
	ws, err := current.Size()
	current.Resize(ws)

	log.Printf("Args: port:%s, rts:%t, dtr:%t\n", cfg.port, cfg.rts, cfg.dtr)
	term, err := terminal.New(
		terminal.SetPort(cfg.port),
		terminal.SetStartupDTR(cfg.dtr),
		terminal.SetStartupRTS(cfg.rts),
		terminal.SetLogFile("log.txt"),
	)
	if err != nil {
		log.Fatal("failed to create:", err)
	}
	if err = term.Open(); err != nil {
		log.Fatal("failed to open port:", err)
	}
	var (
		buf = make([]byte, 2048)
		// wbuf        = make([]byte, 2048)
		n           int
		shiftBuffer int
	)
	logfile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal("failed to open log file:", err)
	}
	defer logfile.Close()
	go func() {
		var (
			n   int
			err error
		)
		for {
			n, err = current.Read(buf)
			if err != nil {
				log.Fatal("Failed to scan:", err, n)
			}
			_, err = fmt.Fprintf(logfile, "input %d chars:%+v\n", n, buf[:n])
			if err != nil {
				log.Fatal("failed to write to file:", err)
			}
			// if n > 0 {
			// 	for int i = 0; i < o
			// 	log.Println("got command:\r\n", wbuf[:n])
			// }
		}
	}()
	for {
		shiftBuffer = 0
		if n, err = term.Read(buf); err != nil {
			log.Fatal("failed to read:", err)
		}
		// _, err = logfile.Write(buf[:n])
		// if err != nil {
		// 	log.Fatal("failed to write to file:", err)
		// }
		if n > 0 {
			for i := 0; i < n; i++ {
				if i > 0 {
					if buf[i-1] == '\r' && buf[i] == '\n' {
						shiftBuffer++
					} else if buf[i-1] == '\r' {

					}

					if buf[i] == '\r' {
						buf[i] = '\n'
					}
				} else {
					if buf[i] == '\r' {
						buf[i] = '\n'
					}
				}
			}
			if n, err = current.Write(buf[:n]); err != nil {
				log.Fatal("failed to write ro stdout:", err)
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}
