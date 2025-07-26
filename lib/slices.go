package lib

import "sync"

type SafeSlice[T any] struct {
	mux   *sync.Mutex
	items []T
}

func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{
		mux:   &sync.Mutex{},
		items: make([]T, 2),
	}
}

func (s *SafeSlice[T]) Append(item T) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.items = append(s.items, item)
}

func (s *SafeSlice[T]) GetAll() []T {
	s.mux.Lock()
	defer s.mux.Unlock()
	return append([]T(nil), s.items...)
}
