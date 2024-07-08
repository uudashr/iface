package db

type Entity interface { // want "interface Entity has no methods, providing no meaningful functionality"
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
