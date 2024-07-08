package db

type Entity interface { // want "interface Entity is empty, providing no meaningful behavior"
}

type Finder interface {
	Find(id string) (Entity, error)
}

type Saver interface {
	Save(e Entity) error
}

type SaveFinder interface {
	Saver
	Finder
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
