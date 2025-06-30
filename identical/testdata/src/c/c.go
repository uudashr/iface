package stream

type StreamNameLister interface { // want "^interface 'StreamNameLister' contains identical methods or type constraints with another interface, causing redundancy \\(see: ConsumerNameLister\\)$"
	Name() <-chan string
	Err() error
}

type ConsumerNameLister interface { // want "^interface 'ConsumerNameLister' contains identical methods or type constraints with another interface, causing redundancy \\(see: StreamNameLister\\)$"
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
	Metadata() map[string]string
}

type KeyValueStatus interface {
	Bucket() string
	Values() uint64
	Bytes() uint64
}
