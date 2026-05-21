package comm

type MessageSender interface {
	SendMessage(username, msg string) error
	MustSendMessage(username, msg string) //iface:ignore=unusedmethod

	// Close the resource.
	Close() error // want "method 'Close\\(\\)' is declared on interface 'MessageSender' but not used within the package"
}

type Service struct {
	sender MessageSender
}

func NewService(sender MessageSender) *Service {
	return &Service{
		sender: sender,
	}
}

func (s *Service) GreetMorning(username string) error {
	return s.sender.SendMessage(username, "Good morning")
}
