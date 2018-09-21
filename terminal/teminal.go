package terminal

import (
	serial "github.com/albenik/go-serial"
)

type Parity serial.Parity

type StopBits serial.StopBits

// Parity values are defined in serial library
const (
	NoParity    = Parity(serial.NoParity)
	OddParity   = Parity(serial.OddParity)
	EvenParity  = Parity(serial.EvenParity)
	MarkParity  = Parity(serial.MarkParity)
	SpaceParity = Parity(serial.SpaceParity)
)

const (
	OneStopBit           = StopBits(serial.OneStopBit)
	OnePointFiveStopBits = StopBits(serial.OnePointFiveStopBits)
	TwoStopBits          = StopBits(serial.TwoStopBits)
)

type Terminal struct {
	// mu           sync.RWMutex
	port         string
	mode         serial.Mode
	internalPort serial.Port
	StateRTS     bool
	StateDTR     bool
	dtrInitState bool
	rtsInitState bool
	LineEnding   string
}

func New(opts ...Option) (*Terminal, error) {
	terminal := Terminal{
		port: "",
		mode: serial.Mode{
			BaudRate: 115200,
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
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
	return err
}

func (t *Terminal) Close() error {
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
