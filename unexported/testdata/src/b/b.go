package b

import "io"

// unexportedReader is an unexported interface.
type unexportedReader interface {
	Read([]byte) (int, error)
}

// ExportedWriter is an exported interface.
type ExportedWriter interface {
	Write([]byte) (int, error)
}

// ExportedType is an exported struct type.
type ExportedType struct{}

// 1. Method with unexported arg type interface
func (e *ExportedType) ReadAll(r unexportedReader) ([]byte, error) { // want "unexported interface 'unexportedReader' used as parameter in exported method 'ExportedType.ReadAll'"
	buf := make([]byte, 1024)
	_, err := r.Read(buf)
	return buf, err
}

// 2. Method with exported arg type interface
func (e *ExportedType) WriteHello(w ExportedWriter) error {
	_, err := w.Write([]byte("hello"))
	return err
}

// 3. Method with unexported return type interface
func (e *ExportedType) NewUnexportedReader() unexportedReader { // want "unexported interface 'unexportedReader' used as return value in exported method 'ExportedType.NewUnexportedReader'"
	return nil // stub
}

// 4. Method with exported return type interface
func (e *ExportedType) NewWriter() ExportedWriter {
	return nil // stub
}

// 5. Method with both unexported arg and return type interface
func (e *ExportedType) WrapReader(r unexportedReader) unexportedReader { // want "unexported interface 'unexportedReader' used as parameter in exported method 'ExportedType.WrapReader'" "unexported interface 'unexportedReader' used as return value in exported method 'ExportedType.WrapReader'"
	return r
}

// 6. Method with both exported arg and return type interface
func (e *ExportedType) WrapWriter(w ExportedWriter) ExportedWriter {
	return w
}

// 7. Method with unexported arg type interface and exported return type interface
func (e *ExportedType) PromoteReader(r unexportedReader) ExportedWriter { // want "unexported interface 'unexportedReader' used as parameter in exported method 'ExportedType.PromoteReader'"
	return nil // stub
}

// 8. Method with exported arg type interface and unexported return type interface
func (e *ExportedType) DemoteWriter(w ExportedWriter) unexportedReader { // want "unexported interface 'unexportedReader' used as return value in exported method 'ExportedType.DemoteWriter'"
	return nil // stub
}

// 9. Method with external arg type interface (io.Writer)
func (e *ExportedType) WriteToExternal(w io.Writer) error {
	_, err := w.Write([]byte("external"))
	return err
}

// 10. Method with external return type interface (io.Writer)
func (e *ExportedType) NewExternalWriter() io.Writer {
	return nil // stub
}

// 11. Method with both external arg and return type interface (io.Writer)
func (e *ExportedType) WrapExternalWriter(w io.Writer) io.Writer {
	return w
}

// 13. Method with any arg type
func (e *ExportedType) Print(a any) {
}

// 14. Method with interface{} arg type
func (e *ExportedType) Write(i interface{}) {
}

// 15. Method with error arg type
func (e *ExportedType) Capture(err error) {}

// 16. Method with no arg and no return type interface
func (e *ExportedType) Ping() {}

// 17. Method with primitive arg type
func (e *ExportedType) AddOne(x int) int {
	return x + 1
}

// 18. Method with primitive return type
func (e *ExportedType) GetZero() int {
	return 0
}

// 19. Method with both primitive arg and return type
func (e *ExportedType) Double(x int) int {
	return x * 2
}
