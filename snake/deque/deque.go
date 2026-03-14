package deque

import "iter"

type Deque[T any] []T

func New[T any]() *Deque[T] {
	d := Deque[T]{}
	return &d
}

func (d *Deque[T]) PushFront(v T) {
	*d = append([]T{v}, *d...)
}

func (d *Deque[T]) PopBack() (v T, ok bool) {
	if len(*d) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(*d) - 1
	v = (*d)[lastIndex]
	*d = (*d)[:lastIndex]
	return v, true
}

func (d *Deque[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range *d {
			if !yield(v) {
				return
			}
		}
	}
}
