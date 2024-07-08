package stream

import "time"

// StreamNameLister is used to iterate over a channel of stream names.
// Err method can be used to check for errors encountered during iteration.
// Name channel is always closed and therefore can be used in a range loop.
type StreamNameLister interface { // want "interface StreamNameLister contains duplicate methods or type constraints from another interface, causing redundancy"
	Name() <-chan string
	Err() error
}

// ConsumerNameLister is used to iterate over a channel of consumer names.
// Err method can be used to check for errors encountered during iteration.
// Name channel is always closed and therefore can be used in a range loop.
type ConsumerNameLister interface { // want "interface ConsumerNameLister contains duplicate methods or type constraints from another interface, causing redundancy"
	Name() <-chan string
	Err() error
}

type ObjectStoresLister interface {
	Status() <-chan ObjectStoreStatus
	Error() error
}

type KeyValueLister interface {
	Status() <-chan KeyValueStatus
	Error() error
}

type ObjectStoreStatus interface {
	Bucket() string
	Description() string
	TTL() time.Duration
	Replicas() int
	Sealed() bool
	Size() uint64
	BackingStore() string
	Metadata() map[string]string
	IsCompressed() bool
}

type KeyValueStatus interface {
	Bucket() string
	Values() uint64
	History() int64
	TTL() time.Duration
	BackingStore() string
	Bytes() uint64
	IsCompressed() bool
}
