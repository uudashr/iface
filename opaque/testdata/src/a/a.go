package app

import "errors"

type Server interface {
	Serve() error
}

type server struct {
	addr string
}

func (s server) Serve() error {
	return nil
}

func NewServer(addr string) Server { // want "'NewServer' function return 'Server' interface at the 1st result, abstract a single concrete implementation of '\\*server'"
	return newServer(addr)
}

func NewServer2(addr string) Server { // want "'NewServer2' function return 'Server' interface at the 1st result, abstract a single concrete implementation of 'server'"
	return server{addr: addr}
}

func NewServer3(addr string) Server { // want "'NewServer3' function return 'Server' interface at the 1st result, abstract a single concrete implementation of '\\*server'"
	return &server{addr: addr}
}

func NewServer4(addr string) (Server, error) { // want "'NewServer4' function return 'Server' interface at the 1st result, abstract a single concrete implementation of '\\*server'"
	if addr == "" {
		return nil, errors.New("addr cannot be nil")
	}
	return &server{addr: addr}, nil
}

func newServer(addr string) *server {
	return &server{addr: addr}
}

type secureServer struct {
	addr string
}

func (ss *secureServer) Serve() error {
	return nil
}

func NewSecureServer(addr string) Server { // want "'NewSecureServer' function return 'Server' interface at the 1st result, abstract a single concrete implementation of '\\*secureServer'"
	return &secureServer{addr: addr}
}
