package app

type Server interface {
	Serve() error
}

type server struct {
	addr string
}

func (s *server) Serve() error {
	return nil
}

func NewServer(addr string) Server { // want "NewServer function return Server at the 1st result, abstract a single concrete implementation of \\*a\\.server"
	return newServer(addr)
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

func NewSecureServer(addr string) Server { // want "NewSecureServer function return Server at the 1st result, abstract a single concrete implementation of \\*a\\.secureServer"
	return &secureServer{addr: addr}
}
