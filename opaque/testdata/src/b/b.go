package app

import "strings"

type Server interface {
	Serve() error
}

type server struct {
	addr string
}

func (s *server) Serve() error {
	return nil
}

type secureServer struct {
	addr string
}

func (ss *secureServer) Serve() error {
	return nil
}

func NewServer(addr string) (Server, bool, error) {
	if strings.HasPrefix(addr, "secure:") {
		return &secureServer{addr: addr}, true, nil
	}

	return &server{addr: addr}, true, nil
}
