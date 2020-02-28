package terminal_test

import (
	"testing"

	"github.com/Aukstkalnis/GoTerminal/terminal"
)

func TestTerminal(t *testing.T) {
	com, err := terminal.New(
		terminal.SetPort("COM3"),
		terminal.SetBaudrate(115200),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = com.Open(); err != nil {
		t.Error(err)
	}
	com.Write([]byte(".info\n"))

}
