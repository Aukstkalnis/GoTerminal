package terminal

type Option func(*Terminal) error

func SetPort(port string) Option {
	return func(opt *Terminal) error {
		opt.Port = port
		return nil
	}
}

func SetBaudrate(baudrate uint32) Option {
	return func(opt *Terminal) error {
		opt.Baudrate = baudrate
		return nil
	}
}

func SetDataBits(dataBits uint8) Option {
	return func(opt *Terminal) error {
		opt.DataBits = dataBits
		return nil
	}
}

func SetParity(parity Parity) Option {
	return func(opt *Terminal) error {
		opt.Parity = parity
		return nil
	}
}

func SetStopBits(stopBits StopBits) Option {
	return func(opt *Terminal) error {
		opt.StopBits = stopBits
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

func SetLineEnding(lineEnding LineEnding) Option {
	return func(opt *Terminal) error {
		opt.LineEnding = lineEnding
		return nil
	}
}

func SetLogFile(fileName string) Option {
	return func(opt *Terminal) error {
		opt.LogFile = fileName
		return nil
	}
}

func SetParser(parseFunction func(b []byte) (n int, err error)) Option {
	return func(opt *Terminal) error {
		opt.packetCapurer = parseFunction
		return nil
	}
}

func SetHandler(handlerFunction func(b []byte) (n int, err error)) Option {
	return func(opt *Terminal) error {
		opt.packetHandler = handlerFunction
		return nil
	}
}
