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

type (
	Reader interface { // want "^interface 'Reader' contains identical methods or type constraints with another interface, causing redundancy \\(see: Scanner\\)$"
		Read() error
	}

	Scanner interface { // want "^interface 'Scanner' contains identical methods or type constraints with another interface, causing redundancy \\(see: Reader\\)$"
		Read() error
	}
)

type (
	TaskRunner interface {
		Run() error
	}

	//iface:ignore
	JobRunner interface {
		Run() error
	}
)

//iface:ignore
type (
	Starter interface {
		Start() error
	}

	Launcher interface {
		Start() error
	}
)

//iface:ignore=identical
type (
	Logger interface {
		Log() error
	}

	AuditLogger interface {
		Log() error
	}
)

//iface:ignore=unused
type (
	Encoder interface { // want "^interface 'Encoder' contains identical methods or type constraints with another interface, causing redundancy \\(see: Serializer\\)$"
		Encode() error
	}

	Serializer interface { // want "^interface 'Serializer' contains identical methods or type constraints with another interface, causing redundancy \\(see: Encoder\\)$"
		Encode() error
	}
)