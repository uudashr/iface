package ignoredirective

//iface:ignore=identical
type (
	Sender interface {
		Send(msg string) error
	}

	Messenger interface {
		Send(msg string) error
	}
)

type (
	//iface:ignore=identical
	Ping interface {
		Ping() error
	}

	Healthcheck interface {
		Ping() error
	}
)
