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

// 19. Function with pointer to unexported interface parameter
func ReadPtr(r *unexportedReader) ([]byte, error) { // want `unexported interface '\*unexportedReader' used as parameter in exported function 'ReadPtr'`
	return nil, nil
}

// 20. Function with pointer to unexported interface return value
func NewReaderPtr() *unexportedReader { // want `unexported interface '\*unexportedReader' used as return value in exported function 'NewReaderPtr'`
	return nil
}

// 21. Function with pointer to unexported interface both param and return
func WrapReaderPtr(r *unexportedReader) *unexportedReader { // want `unexported interface '\*unexportedReader' used as parameter in exported function 'WrapReaderPtr'` `unexported interface '\*unexportedReader' used as return value in exported function 'WrapReaderPtr'`
	return r
}

// 22. Function with pointer to pointer to unexported interface
func ReadPtrPtr(r **unexportedReader) { // want `unexported interface '\*\*unexportedReader' used as parameter in exported function 'ReadPtrPtr'`
}

// 23. Function with pointer to exported interface — no diagnostic
func ReadExportedPtr(r *ExportedWriter) {}

// 24. Function with pointer to external interface — no diagnostic
func ReadExtPtr(r *io.Writer) {}

// 25. Function with variadic unexported interface parameter
func ServeClient(h ...unexportedReader) {} // want `unexported interface '...unexportedReader' used as parameter in exported function 'ServeClient'`

// 26. Function with variadic exported interface — no diagnostic
func ServeExportedVariadic(w ...ExportedWriter) {}

// 27. Function with variadic any — no diagnostic
func ServeOK(h ...any) {}

// 28. Function with variadic error — no diagnostic
func ServeErr(h ...error) {}

// 29. Function with variadic pointer to unexported interface (nested wrapping)
func ServeVariadicPtr(h ...*unexportedReader) {} // want `unexported interface '...\*unexportedReader' used as parameter in exported function 'ServeVariadicPtr'`

// 30. Function with slice of unexported interface parameter
func ReadMany(r []unexportedReader) {} // want `unexported interface '\[\]unexportedReader' used as parameter in exported function 'ReadMany'`

// 31. Function with slice of unexported interface return value
func NewReaderSlice() []unexportedReader { return nil } // want `unexported interface '\[\]unexportedReader' used as return value in exported function 'NewReaderSlice'`

// 32. Function with slice of unexported interface both param and return
func TransformReaders(r []unexportedReader) []unexportedReader { return r } // want `unexported interface '\[\]unexportedReader' used as parameter in exported function 'TransformReaders'` `unexported interface '\[\]unexportedReader' used as return value in exported function 'TransformReaders'`

// 33. Function with slice of exported interface — no diagnostic
func ReadSliceExported(r []ExportedWriter) {}

// 34. Function with slice of external interface — no diagnostic
func ReadSliceExt(r []io.Writer) {}

// 35. Function with array of unexported interface parameter
func ReadFixed(r [4]unexportedReader) {} // want `unexported interface '\[4\]unexportedReader' used as parameter in exported function 'ReadFixed'`

// 36. Function with array of unexported interface return value
func NewReaderArray() [2]unexportedReader { return [2]unexportedReader{} } // want `unexported interface '\[2\]unexportedReader' used as return value in exported function 'NewReaderArray'`

// 37. Function with array of exported interface — no diagnostic
func ReadArrayExported(r [4]ExportedWriter) {}

// 38. Function with array of external interface — no diagnostic
func ReadArrayExt(r [4]io.Writer) {}

// 39. Function with receive-only channel of unexported interface parameter
func ReadChan(r <-chan unexportedReader) {} // want `unexported interface '<-chan unexportedReader' used as parameter in exported function 'ReadChan'`

// 40. Function with send-only channel of unexported interface parameter
func WriteChan(r chan<- unexportedReader) {} // want `unexported interface 'chan<- unexportedReader' used as parameter in exported function 'WriteChan'`

// 41. Function with bidirectional channel of unexported interface parameter
func BidiChan(r chan unexportedReader) {} // want `unexported interface 'chan unexportedReader' used as parameter in exported function 'BidiChan'`

// 42. Function with receive-only channel of unexported interface return value
func NewReaderChan() <-chan unexportedReader { return nil } // want `unexported interface '<-chan unexportedReader' used as return value in exported function 'NewReaderChan'`

// 43. Function with send-only channel of unexported interface return value
func NewReaderChanSend() chan<- unexportedReader { return nil } // want `unexported interface 'chan<- unexportedReader' used as return value in exported function 'NewReaderChanSend'`

// 44. Function with channel of exported interface — no diagnostic
func ReadChanExported(r <-chan ExportedWriter) {}

// 45. Function with channel of external interface — no diagnostic
func ReadChanExt(r <-chan io.Writer) {}

// 46. Function with map value is unexported interface parameter
func ReadMapVal(r map[string]unexportedReader) {} // want `unexported interface 'map\[string\]unexportedReader' used as parameter in exported function 'ReadMapVal'`

// 47. Function with map value is unexported interface return value
func NewReaderMapVal() map[string]unexportedReader { return nil } // want `unexported interface 'map\[string\]unexportedReader' used as return value in exported function 'NewReaderMapVal'`

// 48. Function with map of exported interface — no diagnostic
func ReadMapExported(r map[string]ExportedWriter) {}

// 49. Function with map of external interface — no diagnostic
func ReadMapExt(r map[string]io.Writer) {}

// 50. Function with pointer to slice of unexported interface (nested wrapping)
func ReadPtrSlice(r *[]unexportedReader) {} // want `unexported interface '\*\[\]unexportedReader' used as parameter in exported function 'ReadPtrSlice'`

// 51. Function with pointer to pointer to slice of unexported interface (nested wrapping)
func ReadPtrPtrSlice(r **[]unexportedReader) {} // want `unexported interface '\*\*\[\]unexportedReader' used as parameter in exported function 'ReadPtrPtrSlice'`

// 52. Function with pointer to slice of unexported interface as return value
func NewReaderPtrSlice() *[]unexportedReader { return nil } // want `unexported interface '\*\[\]unexportedReader' used as return value in exported function 'NewReaderPtrSlice'`

// 53. Function with slice of pointer to unexported interface
func ReadSlicePtr(r []*unexportedReader) {} // want `unexported interface '\[\]\*unexportedReader' used as parameter in exported function 'ReadSlicePtr'`

// 54. Function with pointer to slice of external interface — no diagnostic
func ReadPtrSliceExt(r *[]io.Writer) {}

// 55. Function with pointer to slice of exported interface — no diagnostic
func ReadPtrSliceExported(r *[]ExportedWriter) {}

// 56. Function with slice of pointer to external interface — no diagnostic
func ReadSlicePtrExt(r []*io.Writer) {}

// 57. Function with slice of pointer to exported interface — no diagnostic
func ReadSlicePtrExported(r []*ExportedWriter) {}
