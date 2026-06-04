package ignoredirective

type execer interface {
	Exec() (any, error)
}

//iface:ignore=unexported
func Exec(e execer) (any, error) {
	return e.Exec()
}
