package terminal

import (
	"context"
	"errors"
	"time"

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

type WorkingMode uint8

// WorkingMode enum
const (
	RO_WorkingMode = iota
	WO_WorkingMode
	RW_WorkingMode
)

type PortConfig struct {
	RTSState bool
	DTRState bool
	DataBits uint8
	WorkingMode
	Baudrate uint32
	Parity
	StopBits
	Port string
}

type UserConfig struct {
	LogToFile bool
	LogFile   string
}

type Terminal struct {
	// mu           sync.RWMutex
	UserConfig
	PortConfig
	mode         serial.Mode
	internalPort serial.Port
	StateRTS     bool
	StateDTR     bool
	dtrInitState bool
	rtsInitState bool
	LineEnding
	err      error
	opened   bool
	readBuf  []byte
	writeBuf []byte
}

var (
	// PortClosedErr shows that ports is closed or not initialized
	PortClosedErr   = errors.New("port is closed")
	DTRSetErr       = errors.New("failed to set DTR")
	RTSSetErr       = errors.New("failed to set RTS")
	PortNotFoundErr = errors.New("port not found")
)

func New(opts ...Option) (*Terminal, error) {
	terminal := Terminal{
		PortConfig: PortConfig{
			Port: "",
		},
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
	// Load Config
	t.mode.BaudRate = int(t.Baudrate)
	t.mode.DataBits = int(t.DataBits)
	t.mode.Parity = serial.Parity(t.Parity)
	t.mode.StopBits = serial.StopBits(t.StopBits)
	// Open Port
	t.internalPort, err = serial.Open(t.Port, &t.mode)
	if err != nil {
		return err
	}
	if err = t.SetDTR(t.dtrInitState); err != nil {
		t.internalPort.Close()
		return DTRSetErr
	}
	if err = t.SetRTS(t.rtsInitState); err != nil {
		t.internalPort.Close()
		return DTRSetErr
	}

	return nil
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

func (t *Terminal) WaitResponse(response string, tmo time.Duration) error {
	return nil
}

func (t *Terminal) readRoutine(ctx *context.Context) {
	for {
		select {}
	}
	if t.LogToFile && t.LogFile != "" {

	}

}
