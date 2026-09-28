package typeparams

// Generic interface method used through a concrete instantiation.
type Optional[T any] interface {
	Get() (T, bool)
}

func Concrete(o Optional[int]) (int, bool) {
	return o.Get()
}

// Generic interface method used through a type parameter.
type Iter[T any] interface {
	Next() (T, bool)
}

func ThroughParam[T any](it Iter[T]) (T, bool) {
	return it.Next()
}

// Generic interface method used as a method expression.
type Callback[T any] interface {
	Call(T)
}

func Invoke[T any](cb Callback[T], v T) {
	Callback[T].Call(cb, v)
}

// Generic interface method used as a method reference.
type Handler[T any] interface {
	Handle(T)
}

func Register[T any](h Handler[T]) func(T) {
	return h.Handle
}

// A genuinely unused generic interface method must still be reported.
type Resource[T any] interface {
	Acquire() (T, error) // want "^method 'Acquire\\(\\)' is declared on interface 'Resource' but not used within the package$"
	Release() (T, error)
}

func UseResource[T any](r Resource[T]) (T, error) {
	return r.Release()
}
