package number

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Int interface { // want "^interface 'Int' contains identical methods or type constraints with another interface, causing redundancy \\(see: Int32\\)$"
	int32
}

type Int32 interface { // want "^interface 'Int32' contains identical methods or type constraints with another interface, causing redundancy \\(see: Int\\)$"
	int32
}
