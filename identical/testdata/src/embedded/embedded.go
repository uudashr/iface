package embedded

import "io"

type Sender interface { // want "^interface 'Sender' contains identical methods or type constraints with another interface, causing redundancy \\(see: Messenger\\)$"
	Send(msg string) error
}

type Closer interface {
	Close() error
}

type Messenger interface { // want "^interface 'Messenger' contains identical methods or type constraints with another interface, causing redundancy \\(see: Sender\\)$"
	Sender
}

type SendCloser interface { // want "^interface 'SendCloser' contains identical methods or type constraints with another interface, causing redundancy \\(see: Channel, Comm\\)$"
	Sender
	Closer
}

type Comm interface { // want "^interface 'Comm' contains identical methods or type constraints with another interface, causing redundancy \\(see: Channel, SendCloser\\)$"
	Send(msg string) error
	Close() error
}

type Channel interface { // want "^interface 'Channel' contains identical methods or type constraints with another interface, causing redundancy \\(see: Comm, SendCloser\\)$"
	Sender
	io.Closer
}
