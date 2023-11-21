// stack provides a generic abstract stack (LIFO) data structure. The implementation is pointer to slice and requires
// the items in the stack to implement the comparable interface and all be of the same type. This implementation is
// concurrent safe.
// Author: Ken Pratt

package stack

import "sync"

type Stack[T any] struct {
	lock   sync.RWMutex
	unsafe bool
	data   []T
}

// New creates an empty stack of type T.
func New[T any]() *Stack[T] {
	return &Stack[T]{unsafe: false}
}

// NewUnsafe creates an empty stack of type T that is not thread safe.
func NewUnsafe[T any]() *Stack[T] {
	return &Stack[T]{unsafe: true}
}

// Push an item onto this.
func (s *Stack[T]) Push(item T) {
	if !s.unsafe {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	s.data = append(s.data, item)
}

// Clear this of all items.
func (s *Stack[T]) Clear() {
	if !s.unsafe {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	s.data = s.data[:0] // keep allocated memory - don't pound the GC!
}

// Peek return the top item without popping it off. Return nil if this is empty.
func (s *Stack[T]) Peek() (item T) {
	if !s.unsafe {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	if len(s.data) > 0 {
		item = s.data[len(s.data)-1]
	}
	return
}

// Pop the top item off. Return nil if this is empty.
func (s *Stack[T]) Pop() (item T) {
	if !s.unsafe {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	if len(s.data) > 0 {
		item = s.data[len(s.data)-1]
		s.data = s.data[0 : len(s.data)-1]
	}
	return
}

// Size returns the number of items in this.
func (s *Stack[T]) Size() int {
	if !s.unsafe {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	return len(s.data)
}

// IsEmpty returns true if this is empty.
func (s *Stack[T]) IsEmpty() bool {
	if !s.unsafe {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	return len(s.data) == 0
}
