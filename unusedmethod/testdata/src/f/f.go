package comm

type Callback interface {
	Call(msg string)
}

type entry struct {
	cb   Callback
	call func(cb Callback, msg string)
}
type Service struct {
	callbacks []entry
}

func (s *Service) RegisterCallaback(cb Callback) {
	s.callbacks = append(s.callbacks, entry{
		cb:   cb,
		call: Callback.Call,
	})
}
