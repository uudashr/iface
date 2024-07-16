package stream

type NameLister interface {
	Name() <-chan string
	Err() error
}

type StreamNameLister NameLister

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
	Metadata() map[string]string
}

type KeyValueStatus interface {
	Bucket() string
	Values() uint64
	Bytes() uint64
}
