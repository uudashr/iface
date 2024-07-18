package actor

type Pinger interface {
	Ping() error
}

//iface:ignore
type Healthcheck interface {
	Ping() error
}

type Submitter interface {
	Submit(msg string) error
}
