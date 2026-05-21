package comm

type Callback interface {
	Call(msg string)
}

type Service struct {
	callbacks []func(msg string)
}

func (s *Service) RegisterCallaback(cb Callback) {
	s.callbacks = append(s.callbacks, cb.Call)
}
