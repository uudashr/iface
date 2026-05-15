package m

type Conn interface {
	Read() error
}

type tcpConn struct{}

func (c *tcpConn) Read() error { return nil }

type udpConn struct{}

func (c *udpConn) Read() error { return nil }

func GetConn(useTCP bool) (c Conn) {
	if useTCP {
		c = &tcpConn{}
		return
	}
	return &udpConn{}
}