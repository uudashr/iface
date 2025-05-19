package a

import (
	"io"
)

// unexportedReader is an unexported interface.
type unexportedReader interface {
	Read([]byte) (int, error)
}

// ExportedWriter is an exported interface.
type ExportedWriter interface {
	Write([]byte) (int, error)
}

// 1. Function with unexported arg type interface
func ReadAll(r unexportedReader) ([]byte, error) { // want "unexported interface 'unexportedReader' used as parameter in exported function 'ReadAll'"
	buf := make([]byte, 1024)
	_, err := r.Read(buf)
	return buf, err
}

// 2. Function with exported arg type interface
func WriteHello(w ExportedWriter) error {
	_, err := w.Write([]byte("hello"))
	return err
}

// 3. Function with unexported return type interface
func NewUnexportedReader() unexportedReader { // want "unexported interface 'unexportedReader' used as return value in exported function 'NewUnexportedReader'"
	return nil // stub
}

// 4. Function with exported return type interface
func NewWriter() ExportedWriter {
	return nil // stub
}

// 5. Function with both unexported arg and return type interface
func WrapReader(r unexportedReader) unexportedReader { // want "unexported interface 'unexportedReader' used as parameter in exported function 'WrapReader'" "unexported interface 'unexportedReader' used as return value in exported function 'WrapReader'"
	return r
}

// 6. Function with both exported arg and return type interface
func WrapWriter(w ExportedWriter) ExportedWriter {
	return w
}

// 7. Function with unexported arg type interface and exported return type interface
func PromoteReader(r unexportedReader) ExportedWriter { // want "unexported interface 'unexportedReader' used as parameter in exported function 'PromoteReader'"
	return nil // stub
}

// 8. Function with exported arg type interface and unexported return type interface
func DemoteWriter(w ExportedWriter) unexportedReader { // want "unexported interface 'unexportedReader' used as return value in exported function 'DemoteWriter'"
	return nil // stub
}

// 9. Function with external arg type interface (io.Writer)
func WriteToExternal(w io.Writer) error {
	_, err := w.Write([]byte("external"))
	return err
}

// 10. Function with external return type interface (io.Writer)
func NewExternalWriter() io.Writer {
	return nil // stub
}

// 11. Function with both external arg and return type interface (io.Writer)
func WrapExternalWriter(w io.Writer) io.Writer {
	return w
}

// 12. Function with any arg type
func Print(a any) {
}

// 13. Function witn interface{} arg type
func Write(i interface{}) {
}

// 14. Function with error arg type
func Capture(err error) {}

// 15. Function with no arg and no return type interface
func Ping() {}

// 16. Function with primitive arg type
func AddOne(x int) int {
	return x + 1
}

// 17. Function with primitive return type
func GetZero() int {
	return 0
}

// 18. Function with both primitive arg and return type
func Double(x int) int {
	return x * 2
}
