package excludepkg

type Sender interface {
	Send(msg string) error
}
