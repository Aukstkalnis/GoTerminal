package terminal

import serial "github.com/albenik/go-serial"

type Option func(*Terminal) error

func SetPort(port string) Option {
	return func(opt *Terminal) error {
		opt.port = port
		return nil
	}
}

func SetBaudrate(baudrate uint32) Option {
	return func(opt *Terminal) error {
		opt.mode.BaudRate = int(baudrate)
		return nil
	}
}

func SetDataBits(dataBits int) Option {
	return func(opt *Terminal) error {
		opt.mode.DataBits = dataBits
		return nil
	}
}

func SetParity(parity Parity) Option {
	return func(opt *Terminal) error {
		opt.mode.Parity = serial.Parity(parity)
		return nil
	}
}

func SetStopBits(stopBits StopBits) Option {
	return func(opt *Terminal) error {
		opt.mode.StopBits = serial.StopBits(stopBits)
		return nil
	}
}

func SetStartupDTR(state bool) Option {
	return func(opt *Terminal) error {
		opt.dtrInitState = state
		return nil
	}
}

func SetStartupRTS(state bool) Option {
	return func(opt *Terminal) error {
		opt.rtsInitState = state
		return nil
	}

}
