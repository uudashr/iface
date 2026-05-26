package typedefalias

type StringResult interface {
	Result() <-chan string
	Error() error
}

type NameResult StringResult

type MessageResult = StringResult
