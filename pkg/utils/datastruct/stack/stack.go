package stack

import "sync"

type (
	Stack[T any] struct {
		top    *node[T]
		mu     sync.Mutex
		length int
	}
	node[T any] struct {
		value T
		prev  *node[T]
	}
)

func New[T any]() *Stack[T] {
	return &Stack[T]{
		top:    nil,
		length: 0,
	}
}

func (s *Stack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.top = &node[T]{
		value: value,
		prev:  s.top,
	}
	s.length++
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var value T
	if s.length == 0 {
		return value, false
	}
	value = s.top.value
	s.top = s.top.prev
	s.length--
	return value, true
}

func (s *Stack[T]) Length() int {
	return s.length
}
