package comm

type Pinger interface {
	Ping() error
}

type Healthcheck = Pinger

type Checker = Pinger

type Submitter interface {
	Submit(msg string) error
}
