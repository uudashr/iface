package actor

// Doer performs some work.
//
//iface:ignore
type Doer interface {
	Do() error
}

type Greeter interface { // want "interface 'Greeter' is declared but not used within the package"
	Greet() error
}

//iface:ignore=unused
type Runner interface {
	Run() error
}

//iface:ignore=other
type Executor interface { // want "interface 'Executor' is declared but not used within the package"
	Execute() error
}

//iface:ignore=other,unused
type Server interface {
	Serve() error
}
