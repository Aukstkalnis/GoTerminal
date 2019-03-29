package terminal

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	serial "github.com/albenik/go-serial"
	"github.com/sirupsen/logrus"
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
	WorkingMode
	DataBits uint8
	Baudrate uint32
	Parity
	StopBits
	Port string
}

type UserConfig struct {
	LogToFile bool
	LogFile   string
}

type internal struct {
	opened       bool
	dtrInitState bool
	rtsInitState bool
	port         serial.Port
}

type Terminal struct {
	mu sync.RWMutex
	UserConfig
	PortConfig
	StateRTS bool
	StateDTR bool
	LineEnding
	internal
	err      error
	readBuf  []byte
	writeBuf []byte
}

var (
	// PortClosedErr shows that ports is closed or not initialized
	ErrPortClosed   = errors.New("port is closed")
	ErrDTRSet       = errors.New("failed to set DTR")
	ErrRTSSet       = errors.New("failed to set RTS")
	ErrPortNotFound = errors.New("port not found")
	ErrPortNotSet   = errors.New("port not set")
)

func New(opts ...Option) (*Terminal, error) {
	terminal := Terminal{
		PortConfig: PortConfig{
			Port:     "",
			Baudrate: 115200,
			DataBits: 8,
			Parity:   NoParity,
			StopBits: OneStopBit,
		},
		LineEnding: "\r",
		internal: internal{
			opened:       false,
			dtrInitState: false,
			rtsInitState: false,
		},
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
	if t.Port == "" {
		return ErrPortNotSet
	}
	t.internal.port, err = serial.Open(
		t.Port,
		&serial.Mode{int(t.Baudrate),
			int(t.DataBits),
			serial.Parity(t.Parity),
			serial.StopBits(t.StopBits)},
	)
	if err != nil {
		return err
	}
	if err = t.SetDTR(t.dtrInitState); err != nil {
		t.internal.port.Close()
		return ErrDTRSet
	}
	if err = t.SetRTS(t.rtsInitState); err != nil {
		t.internal.port.Close()
		return ErrRTSSet
	}
	t.opened = true
	if t.WorkingMode == RO_WorkingMode || t.WorkingMode == RW_WorkingMode {
		logrus.Println("Start reading goroutine")
		go t.readRoutine(context.Background())
	}
	return nil
}

func (t *Terminal) Close() error {
	t.opened = false
	return t.internal.port.Close()
}

func (t *Terminal) Write(b []byte) (n int, err error) {
	return t.internal.port.Write(b)
}

func (t *Terminal) Read(b []byte) (int, error) {
	return t.internal.port.Read(b)
}

func (t *Terminal) SetDTR(state bool) error {
	err := t.internal.port.SetDTR(state)
	if err != nil {
		t.StateDTR = state
	}
	return err
}

func (t *Terminal) SetRTS(state bool) error {
	err := t.internal.port.SetRTS(state)
	if err != nil {
		t.StateRTS = state
	}
	return err
}

func (t *Terminal) WaitResponse(response string, tmo time.Duration) error {
	t.mu.Lock()
	t.mu.Unlock()
	return nil
}

func (t *Terminal) readRoutine(ctx context.Context) {
	var (
		// n    int
		err  error
		file *os.File
	)
	if len(t.readBuf) == 0 {
		t.readBuf = make([]byte, 1204)
	}
	for {
		if t.opened {
			if _, err = t.internal.port.Read(t.readBuf); err != nil {
				logrus.Println("failed to read:", err)
				t.err = err
				if file != nil {
					if err = file.Close(); err != nil {
						logrus.Println("failed to close log file:", err)
					}
				}
				break
			}
		} else {
			log.Println("port closed, stop reading")
			break
		}
	}
}
