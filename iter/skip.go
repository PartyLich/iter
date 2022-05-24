package iter

// Skipped is an iterator that skips over n elements.
type Skipped[T any] struct {
	iter Iterable[T]
	n    int
}

// Skip creates an iterator that skips the first n elements.
func Skip[T any](iter Iterable[T], n int) *Skipped[T] {
	if n < 0 {
		panic("Skip requires n >= 0")
	}

	return &Skipped[T]{iter, n}
}

func (s *Skipped[T]) Next() *T {
	for s.n != 0 {
		s.n -= 1
		s.iter.Next()
	}

	return s.iter.Next()
}

func (iter *Skipped[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

//go:generate go run ./cmd/gen/ -name Skipped -output skip_ext_gen.go
