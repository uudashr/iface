package typeparams

import (
	"fmt"
)

type execer[T any] interface {
	Exec() (T, error)
}

func Exec[T any](name string, reg map[string]execer[T]) (T, error) { // want "^unexported interface 'map\\[string\\]execer\\[T\\]' used as parameter in exported function 'Exec'$"
	exe, ok := reg[name]
	if !ok {
		var zero T
		return zero, fmt.Errorf("no %q on reg", name)
	}

	return exe.Exec()
}

type matcher[T any] interface {
	Match(T) bool
}

type Repository[T any] struct {
	entries []T
}

func (r *Repository[T]) Query(m matcher[T]) ([]T, error) { // want "^unexported interface 'matcher\\[T\\]' used as parameter in exported method 'Repository\\[T\\].Query'$"
	var res []T
	for _, e := range r.entries {
		if ok := m.Match(e); !ok {
			continue
		}
		res = append(res, e)
	}
	return res, nil
}
