package shadow

type Notifier interface {
	Notify() error
}

func NotifyAll[n any](items []n) error {
	for _, item := range items {
		if notifier, ok := any(item).(Notifier); ok {
			_ = notifier.Notify()
		}
	}
	return nil
}