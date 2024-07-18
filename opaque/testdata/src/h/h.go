package app

type Server interface {
	Serve() error
}

type server struct {
	addr string
}

func (s server) Serve() error {
	return nil
}

//iface:ignore
func NewServer(addr string) Server {
	return newServer(addr)
}

func NewServer2(addr string) Server { // want "NewServer2 function return Server interface at the 1st result, abstract a single concrete implementation of server"
	return server{addr: addr}
}

func newServer(addr string) *server {
	return &server{addr: addr}
}
