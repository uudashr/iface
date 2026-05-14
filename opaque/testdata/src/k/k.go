package k

type Doer interface {
	Do()
}

type doerImpl struct{}

func (d *doerImpl) Do() {}

func NilOnly() Doer {
	return nil
}

type container struct {
	Impl *doerImpl
}

var c container

func FromStructField() Doer { // want "'FromStructField' function return 'Doer' interface at the 1st result, abstract a single concrete implementation of '\\*doerImpl'"
	return c.Impl
}

func FromTypeAssertion(i interface{}) Doer { // want "'FromTypeAssertion' function return 'Doer' interface at the 1st result, abstract a single concrete implementation of '\\*doerImpl'"
	return i.(*doerImpl)
}

func FromInterfaceAssertion(i Doer) Doer {
	return i.(Doer)
}