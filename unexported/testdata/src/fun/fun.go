package fun

import (
	"fmt"
	"io"
	"os"
)

type execer interface {
	Exec() (any, error)
}

type Caller interface {
	Call() (any, error)
}

type AddTask struct {
	a int
	b int
}

func (t *AddTask) Exec() (any, error) {
	return t.a + t.b, nil
}

type CallFunc func() (any, error)

func (cf CallFunc) Call() (any, error) {
	return cf()
}

func ExecAdd(a, b int) execer { // want "^unexported interface 'execer' used as return value in exported function 'ExecAdd'$"
	return &AddTask{a, b}
}

func CallAdd(a, b int) Caller {
	return CallFunc(ExecAdd(a, b).Exec)
}

func NewAddTask(a, b int) *AddTask {
	return &AddTask{a, b}
}

func Exec(e execer) (any, error) { // want "^unexported interface 'execer' used as parameter in exported function 'Exec'$"
	return e.Exec()
}

func ExecPtr(e *execer) (any, error) { // want "^unexported interface '\\*execer' used as parameter in exported function 'ExecPtr'$"
	return (*e).Exec()
}

func exe(c execer) (any, error) {
	return Exec(c)
}

func DumpExec(w io.Writer, c execer) { // want "^unexported interface 'execer' used as parameter in exported function 'DumpExec'$"
	out, err := Exec(c)
	if err != nil {
		fmt.Fprintf(w, "err: %v\n", err)
		return
	}
	fmt.Fprintf(w, "out: %v\n", out)
}

func Call(c Caller) (any, error) {
	return c.Call()
}

func ExecAll(execs []execer) (any, error) { // want "^unexported interface '\\[\\]execer' used as parameter in exported function 'ExecAll'$"
	var out []any
	for _, e := range execs {
		res, err := e.Exec()
		if err != nil {
			return nil, err
		}
		out = append(out, res)
	}
	return out, nil
}

func CallAll(calls []Caller) (any, error) {
	var out []any
	for _, c := range calls {
		res, err := c.Call()
		if err != nil {
			return nil, err
		}
		out = append(out, res)
	}
	return out, nil
}

func Execs(a ...execer) (any, error) { // want "^unexported interface '...execer' used as parameter in exported function 'Execs'$"
	return ExecAll(a)
}

func Calls(a ...Caller) (any, error) {
	return CallAll(a)
}

func ExecChan(ch <-chan execer) chan<- any { // want "^unexported interface '<-chan execer' used as parameter in exported function 'ExecChan'$"
	out := make(chan any)
	go func() {
		for v := range ch {
			res, err := v.Exec()
			if err != nil {
				continue
			}

			out <- res
		}
	}()
	return out
}

func LogError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s err: %v\n", msg, err)
}

func LogInfo(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "INFO: %s\n", msg)
}
