package user

type (
	Doer interface { // want "interface 'Doer' is declared but not used within the package"
		Do() error
	}
)