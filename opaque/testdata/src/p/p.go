package p

import "fmt"

type Server interface {
	Serve() error
}

type server struct{}

func (s *server) Serve() error { return nil }

func NewServer(ok bool) (s Server, err error) { // want "'NewServer' function return 'Server' interface at the 1st result, abstract a single concrete implementation of '\\*server'"
	s = &server{}
	if !ok {
		return nil, fmt.Errorf("fail")
	}
	return
}