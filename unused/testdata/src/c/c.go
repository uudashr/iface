package log

type Logger interface {
	Log(keyvals ...interface{}) error
}
