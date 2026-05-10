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

// 20. Method with pointer to unexported interface parameter
func (e *ExportedType) ReadPtr(r *unexportedReader) ([]byte, error) { // want `unexported interface '\*unexportedReader' used as parameter in exported method 'ExportedType.ReadPtr'`
	return nil, nil
}

// 21. Method with pointer to unexported interface return value
func (e *ExportedType) NewReaderPtr() *unexportedReader { // want `unexported interface '\*unexportedReader' used as return value in exported method 'ExportedType.NewReaderPtr'`
	return nil
}

// 22. Method with pointer to unexported interface both param and return
func (e *ExportedType) WrapReaderPtr(r *unexportedReader) *unexportedReader { // want `unexported interface '\*unexportedReader' used as parameter in exported method 'ExportedType.WrapReaderPtr'` `unexported interface '\*unexportedReader' used as return value in exported method 'ExportedType.WrapReaderPtr'`
	return r
}

// 23. Method with pointer to pointer to unexported interface
func (e *ExportedType) ReadPtrPtr(r **unexportedReader) { // want `unexported interface '\*\*unexportedReader' used as parameter in exported method 'ExportedType.ReadPtrPtr'`
}

// 24. Method with pointer to exported interface — no diagnostic
func (e *ExportedType) ReadExportedPtr(r *ExportedWriter) {}

// 25. Method with pointer to external interface — no diagnostic
func (e *ExportedType) ReadExtPtr(r *io.Writer) {}

// 26. Method with variadic unexported interface parameter
func (e *ExportedType) ServeClient(h ...unexportedReader) {} // want `unexported interface '...unexportedReader' used as parameter in exported method 'ExportedType.ServeClient'`

// 27. Method with variadic exported interface — no diagnostic
func (e *ExportedType) ServeExportedVariadic(w ...ExportedWriter) {}

// 28. Method with variadic pointer to unexported interface (nested wrapping)
func (e *ExportedType) ServeVariadicPtr(h ...*unexportedReader) {} // want `unexported interface '...\*unexportedReader' used as parameter in exported method 'ExportedType.ServeVariadicPtr'`

// 29. Method with slice of unexported interface parameter
func (e *ExportedType) ReadMany(r []unexportedReader) {} // want `unexported interface '\[\]unexportedReader' used as parameter in exported method 'ExportedType.ReadMany'`

// 30. Method with slice of unexported interface return value
func (e *ExportedType) NewReaderSlice() []unexportedReader { return nil } // want `unexported interface '\[\]unexportedReader' used as return value in exported method 'ExportedType.NewReaderSlice'`

// 31. Method with slice of unexported interface both param and return
func (e *ExportedType) TransformReaders(r []unexportedReader) []unexportedReader { return r } // want `unexported interface '\[\]unexportedReader' used as parameter in exported method 'ExportedType.TransformReaders'` `unexported interface '\[\]unexportedReader' used as return value in exported method 'ExportedType.TransformReaders'`

// 32. Method with slice of exported interface — no diagnostic
func (e *ExportedType) ReadSliceExported(r []ExportedWriter) {}

// 33. Method with slice of external interface — no diagnostic
func (e *ExportedType) ReadSliceExt(r []io.Writer) {}

// 34. Method with array of unexported interface parameter
func (e *ExportedType) ReadFixed(r [4]unexportedReader) {} // want `unexported interface '\[4\]unexportedReader' used as parameter in exported method 'ExportedType.ReadFixed'`

// 35. Method with array of unexported interface return value
func (e *ExportedType) NewReaderArray() [2]unexportedReader { return [2]unexportedReader{} } // want `unexported interface '\[2\]unexportedReader' used as return value in exported method 'ExportedType.NewReaderArray'`

// 36. Method with array of exported interface — no diagnostic
func (e *ExportedType) ReadArrayExported(r [4]ExportedWriter) {}

// 37. Method with array of external interface — no diagnostic
func (e *ExportedType) ReadArrayExt(r [4]io.Writer) {}

// 38. Method with receive-only channel of unexported interface parameter
func (e *ExportedType) ReadChan(r <-chan unexportedReader) {} // want `unexported interface '<-chan unexportedReader' used as parameter in exported method 'ExportedType.ReadChan'`

// 39. Method with send-only channel of unexported interface parameter
func (e *ExportedType) WriteChan(r chan<- unexportedReader) {} // want `unexported interface 'chan<- unexportedReader' used as parameter in exported method 'ExportedType.WriteChan'`

// 40. Method with bidirectional channel of unexported interface parameter
func (e *ExportedType) BidiChan(r chan unexportedReader) {} // want `unexported interface 'chan unexportedReader' used as parameter in exported method 'ExportedType.BidiChan'`

// 41. Method with receive-only channel of unexported interface return value
func (e *ExportedType) NewReaderChan() <-chan unexportedReader { return nil } // want `unexported interface '<-chan unexportedReader' used as return value in exported method 'ExportedType.NewReaderChan'`

// 42. Method with send-only channel of unexported interface return value
func (e *ExportedType) NewReaderChanSend() chan<- unexportedReader { return nil } // want `unexported interface 'chan<- unexportedReader' used as return value in exported method 'ExportedType.NewReaderChanSend'`

// 43. Method with channel of exported interface — no diagnostic
func (e *ExportedType) ReadChanExported(r <-chan ExportedWriter) {}

// 44. Method with channel of external interface — no diagnostic
func (e *ExportedType) ReadChanExt(r <-chan io.Writer) {}

// 45. Method with map value is unexported interface parameter
func (e *ExportedType) ReadMapVal(r map[string]unexportedReader) {} // want `unexported interface 'map\[string\]unexportedReader' used as parameter in exported method 'ExportedType.ReadMapVal'`

// 46. Method with map value is unexported interface return value
func (e *ExportedType) NewReaderMapVal() map[string]unexportedReader { return nil } // want `unexported interface 'map\[string\]unexportedReader' used as return value in exported method 'ExportedType.NewReaderMapVal'`

// 47. Method with map of exported interface — no diagnostic
func (e *ExportedType) ReadMapExported(r map[string]ExportedWriter) {}

// 48. Method with map of external interface — no diagnostic
func (e *ExportedType) ReadMapExt(r map[string]io.Writer) {}

// 49. Method with pointer to slice of unexported interface (nested wrapping)
func (e *ExportedType) ReadPtrSlice(r *[]unexportedReader) {} // want `unexported interface '\*\[\]unexportedReader' used as parameter in exported method 'ExportedType.ReadPtrSlice'`

// 50. Method with pointer to pointer to slice of unexported interface (nested wrapping)
func (e *ExportedType) ReadPtrPtrSlice(r **[]unexportedReader) {} // want `unexported interface '\*\*\[\]unexportedReader' used as parameter in exported method 'ExportedType.ReadPtrPtrSlice'`

// 51. Method with pointer to slice of unexported interface as return value
func (e *ExportedType) NewReaderPtrSlice() *[]unexportedReader { return nil } // want `unexported interface '\*\[\]unexportedReader' used as return value in exported method 'ExportedType.NewReaderPtrSlice'`

// 52. Method with slice of pointer to unexported interface
func (e *ExportedType) ReadSlicePtr(r []*unexportedReader) {} // want `unexported interface '\[\]\*unexportedReader' used as parameter in exported method 'ExportedType.ReadSlicePtr'`

// 53. Method with pointer to slice of external interface — no diagnostic
func (e *ExportedType) ReadPtrSliceExt(r *[]io.Writer) {}

// 54. Method with pointer to slice of exported interface — no diagnostic
func (e *ExportedType) ReadPtrSliceExported(r *[]ExportedWriter) {}

// 55. Method with slice of pointer to external interface — no diagnostic
func (e *ExportedType) ReadSlicePtrExt(r []*io.Writer) {}

// 56. Method with slice of pointer to exported interface — no diagnostic
func (e *ExportedType) ReadSlicePtrExported(r []*ExportedWriter) {}
