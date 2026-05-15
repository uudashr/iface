package o

type Server interface {
	Serve() error
}

func BrokenServer() (s Server) {
	return
}