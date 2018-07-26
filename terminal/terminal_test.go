package terminal_test

import (
	"testing"

	"github.com/Aukstkalnis/GoTerminal/terminal"
)

func TestTerminal(t *testing.T) {
	term, err := terminal.New(
		// terminal.SetPort("COM3"),
		terminal.SetBaudrate(115200),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = term.Open(); err != nil {
		t.Error(err)
	}
	term.Write([]byte("This is test string\n"))

}
