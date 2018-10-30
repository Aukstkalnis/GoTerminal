package terminal

import (
	"errors"
	"sync"

	serial "github.com/albenik/go-serial"
)

type Parity serial.Parity

// Parity values are defined in serial library
const (
	NoParity    = Parity(serial.NoParity)
	OddParity   = Parity(serial.OddParity)
	EvenParity  = Parity(serial.EvenParity)
	MarkParity  = Parity(serial.MarkParity)
	SpaceParity = Parity(serial.SpaceParity)
)

type StopBits serial.StopBits

const (
	OneStopBit           = StopBits(serial.OneStopBit)
	OnePointFiveStopBits = StopBits(serial.OnePointFiveStopBits)
	TwoStopBits          = StopBits(serial.TwoStopBits)
)

type LineEnding string

type Terminal struct {
	mu           sync.RWMutex
	port         string
	mode         serial.Mode
	internalPort serial.Port
	StateRTS     bool
	StateDTR     bool
	dtrInitState bool
	rtsInitState bool
	LineEnding
	opened  bool
	err     error
	readErr error
	readBuf []byte
}

var (
	// PortClosedErr shows that ports is closed or not initialized
	PortClosedErr = errors.New("port is closed")
)

func New(opts ...Option) (*Terminal, error) {
	terminal := Terminal{
		port: "",
		mode: serial.Mode{
			BaudRate: 115200,
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
		LineEnding:   "\r",
		dtrInitState: false,
		rtsInitState: false,
	}
	var err error
	for _, o := range opts {
		if err = o(&terminal); err != nil {
			return nil, err
		}
	}
	return &terminal, err
}

func (t *Terminal) Open() (err error) {
	t.internalPort, err = serial.Open(t.port, &t.mode)
	if err == nil {
		err = t.SetDTR(t.dtrInitState)
		if err == nil {
			err = t.SetRTS(t.rtsInitState)
		}
	}
	if err != nil {
		t.opened = false
	}
	return err
}

func (t *Terminal) Close() error {
	t.opened = false
	return t.internalPort.Close()
}

func (t *Terminal) Write(b []byte) (n int, err error) {
	return t.internalPort.Write(b)
}

func (t *Terminal) Read(b []byte) (int, error) {
	return t.internalPort.Read(b)
}

func (t *Terminal) SetDTR(state bool) error {
	err := t.internalPort.SetDTR(state)
	if err != nil {
		t.StateDTR = state
	}
	return err
}

func (t *Terminal) SetRTS(state bool) error {
	err := t.internalPort.SetRTS(state)
	if err != nil {
		t.StateRTS = state
	}
	return err
}

func (t *Terminal) read() {
	// var n int
	// t.mu.RLock()
	// t.mu.RUnlock()
	// for {
	// 	if n, t.readErr = t.internalPort.Read(t.readBuf); t.readErr != nil {
	// 		return
	// 	}
	// 	n = 0
	// }
}
