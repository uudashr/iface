package basic

type Healthcheck interface { // want "^interface 'Healthcheck' contains identical methods or type constraints with another interface, causing redundancy \\(see: Pinger\\)$"
	Ping() error
}

type Pinger interface { // want "^interface 'Pinger' contains identical methods or type constraints with another interface, causing redundancy \\(see: Healthcheck\\)$"
	Ping() error
}

type Sender interface {
	Send(msg string) error
}

type IntStreamer interface {
	Stream() <-chan int
	Error() error
}

type StringStreamer interface {
	Stream() <-chan string
	Error() error
}

type (
	Runner interface { // want "^interface 'Runner' contains identical methods or type constraints with another interface, causing redundancy \\(see: Task\\)$"
		Run() error
	}

	Job struct {
		Name string
	}

	Task interface { // want "^interface 'Task' contains identical methods or type constraints with another interface, causing redundancy \\(see: Runner\\)$"
		Run() error
	}
)
