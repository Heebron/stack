// stack provides a generic abstract stack (LIFO) data structure. The implementation is pointer to slice and requires
// the items in the stack to implement the comparable interface and all be of the same type. This implementation is
// concurrent safe.
// Author: Ken Pratt

package stack

import "sync"

type Stack[T comparable] struct {
	lock sync.RWMutex
	data []T
}

// New create an empty stack of type T.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push an item onto this.
func (s *Stack[T]) Push(item T) {
	s.lock.Lock()
	s.data = append(s.data, item)
	s.lock.Unlock()
}

// Clear this of all items.
func (s *Stack[T]) Clear() {
	s.lock.Lock()
	s.data = s.data[:0] // keep allocated memory - don't pound the GC!
	s.lock.Unlock()
}

// Peek return the top item without popping it off. Return nil if this is empty.
func (s *Stack[T]) Peek() (item T) {
	s.lock.RLock()
	if len(s.data) > 0 {
		item = s.data[len(s.data)-1]
	}
	s.lock.RUnlock()
	return
}

// Pop the top item off. Return nil if this is empty.
func (s *Stack[T]) Pop() (item T) {
	s.lock.RLock()
	if len(s.data) > 0 {
		item = s.data[len(s.data)-1]
		s.data = s.data[0 : len(s.data)-1]
	}
	s.lock.RUnlock()
	return
}

// Size returns the number of items in this.
func (s *Stack[T]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

// IsEmpty returns true if this is empty.
func (s *Stack[T]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data) == 0
}
