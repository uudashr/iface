package g

type (
	A interface { // want "^interface 'A' contains identical methods or type constraints with another interface, causing redundancy \\(see: C\\)$"
		Foo()
	}

	B struct{}

	C interface { // want "^interface 'C' contains identical methods or type constraints with another interface, causing redundancy \\(see: A\\)$"
		Foo()
	}
)
