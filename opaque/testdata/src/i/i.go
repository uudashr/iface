package i

func OnlyError() error {
	return nil
}

func OnlyAny() any {
	return nil
}

func OnlyInterfaceEmpty() interface{} {
	return nil
}

type MyInterface interface {
	Do()
}

type myInterfaceImpl struct{}

func (m *myInterfaceImpl) Do() {}

func WithError() (MyInterface, error) { // want "'WithError' function return 'MyInterface' interface at the 1st result, abstract a single concrete implementation of '\\*myInterfaceImpl'"
	return &myInterfaceImpl{}, nil
}

type myError struct{}

func (e *myError) Error() string {
	return "error"
}

type error interface { // this shadows the predeclared error
	Error() string
}

func GetError() error { // want "'GetError' function return 'error' interface at the 1st result, abstract a single concrete implementation of '\\*myError'"
	return &myError{}
}

type MyError interface {
	Error() string
}

type myMyError struct{}

func (e *myMyError) Error() string {
	return "error"
}

func GetMyError() MyError { // want "'GetMyError' function return 'MyError' interface at the 1st result, abstract a single concrete implementation of '\\*myMyError'"
	return &myMyError{}
}