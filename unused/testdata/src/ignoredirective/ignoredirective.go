package ignoredirective

type Sender interface { // want "^interface 'Sender' is declared but not used within the package$"
	Send(msg string) error
}

//iface:ignore=unused
type MessageHandler interface {
	HandleMessage(msg string) error
}

//iface:ignore=identical
type Commitable interface { // want "^interface 'Commitable' is declared but not used within the package$"
	Commit() error
}

type (
	//iface:ignore=unused
	Executor interface {
		Execute() error
	}

	Job struct {
		Name string
	}

	Runner interface { // want "^interface 'Runner' is declared but not used within the package$"
		Run() error
	}

	//iface:ignore=identical
	Calculator interface { // want "^interface 'Calculator' is declared but not used within the package$"
		Calc() int
	}
)

//iface:ignore=unused
type (
	Writer interface {
		Write(b []byte) error
	}

	Reader interface {
		Read() ([]byte, error)
	}
)

//iface:ignore=identical
type (
	Caller interface { // want "^interface 'Caller' is declared but not used within the package$"
		Call() error
	}
)
