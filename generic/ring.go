package generic

import "container/ring"

// Ring puts some type safety around container/ring.Ring
type Ring[T any] ring.Ring

func NewRing[T any](n int) *Ring[T] {
	return (*Ring[T])(ring.New(n))
}

func (r *Ring[T]) unsafe() *ring.Ring {
	if r == nil {
		return nil
	}
	return (*ring.Ring)(r)
}

func (r *Ring[T]) Do(f func(T)) {
	r.unsafe().Do(func(a any) { f(a.(T)) })
}

func (r *Ring[T]) Len() int {
	return r.unsafe().Len()
}

func (r *Ring[T]) Link(s *Ring[T]) *Ring[T] {
	if r == nil {
		return nil
	}
	return (*Ring[T])(r.unsafe().Link(s.unsafe()))
}

func (r *Ring[T]) Move(n int) *Ring[T] {
	if r == nil {
		return nil
	}
	return (*Ring[T])(r.unsafe().Move(n))
}

func (r *Ring[T]) Next() *Ring[T] {
	if r == nil {
		return nil
	}
	return (*Ring[T])(r.unsafe().Next())
}

func (r *Ring[T]) Prev() *Ring[T] {
	return (*Ring[T])(r.unsafe().Prev())
}

func (r *Ring[T]) Unlink(n int) *Ring[T] {
	return (*Ring[T])(r.unsafe().Unlink(n))
}

func (r *Ring[T]) Set(val T) {
	r.unsafe().Value = val
}

func (r *Ring[T]) Get() T {
	return r.unsafe().Value.(T)
}
