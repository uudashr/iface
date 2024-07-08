package stream

import "time"

type NameLister interface {
	Name() <-chan string
	Err() error
}

// StreamNameLister is used to iterate over a channel of stream names.
// Err method can be used to check for errors encountered during iteration.
// Name channel is always closed and therefore can be used in a range loop.
type StreamNameLister NameLister

// ConsumerNameLister is used to iterate over a channel of consumer names.
// Err method can be used to check for errors encountered during iteration.
// Name channel is always closed and therefore can be used in a range loop.
type ConsumerNameLister NameLister

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
