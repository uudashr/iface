package comm

//iface:ignore=unusedmethod
type MessageSender interface {
	SendMessage(username, msg string) error
	MustSendMessage(username, msg string)

	// Close the resource.
	Close() error
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
