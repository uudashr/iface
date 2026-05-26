package typeconstraints

type Int interface { // want "^interface 'Int' contains identical methods or type constraints with another interface, causing redundancy \\(see: Integer\\)$"
	~int | ~int32 | ~int64
}

type Integer interface { // want "^interface 'Integer' contains identical methods or type constraints with another interface, causing redundancy \\(see: Int\\)$"
	~int | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}
