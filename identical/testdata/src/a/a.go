package comm

type Pinger interface { // want "interface 'Pinger' contains identical methods or type constraints with another interface, causing redundancy"
	Ping() error
}

type Healthcheck interface { // want "interface 'Healthcheck' contains identical methods or type constraints with another interface, causing redundancy"
	Ping() error
}

type Checker interface { // want "interface 'Checker' contains identical methods or type constraints with another interface, causing redundancy"
	Pinger
}

type PingPonger interface {
	Pinger
	Pong() error
}

type Submitter interface {
	Submit(msg string) error
}

type PingSubmitter interface {
	Pinger
	Submitter
}

type SubmitPingPonger interface {
	Submitter
	PingPonger
}
