package basic

// -- used as argument

type Sender interface {
	Send(msg string) error
}

func Ping(sender Sender) error {
	return sender.Send("ping")
}

// -- used as struct field

type MessageHandler interface {
	HandleMessage(msg string) error
}

type Consumer struct {
	Handler MessageHandler
}

// -- used as type cast

type Commitable interface {
	Commit() error
}

func CommitAll(msgs []any) error {
	for _, msg := range msgs {
		c, ok := msg.(Commitable)
		if !ok {
			continue
		}

		if err := c.Commit(); err != nil {
			return err
		}
	}

	return nil
}

// grouped type declaration

type (
	Executor interface { // want "^interface 'Executor' is declared but not used within the package$"
		Execute() error
	}

	Job struct {
		Name string
	}

	Runner interface { // want "^interface 'Runner' is declared but not used within the package$"
		Run() error
	}
)

// -- used as return value

type Calculator interface {
	Calc() int
}

type Task struct {
	A int
	B int
}

func (t *Task) Calc() int {
	return t.A + t.B
}

func Add(a, b int) Calculator {
	return &Task{A: a, B: b}
}

// -- single type declaration

type Writer interface { // want "^interface 'Writer' is declared but not used within the package$"
	Write(b []byte) error
}
