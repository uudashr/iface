package protocol

type Reader interface {
	Read() string
}

func readMessage(b string, d string) (attributes int8, baseOffset, timestamp int64, key Reader, err error) {
	return
}
