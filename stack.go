// stack provides a generic abstract stack (LIFO) data structure. The implementation is pointer to slice.
// Author: Ken Pratt

package stack

type Stack []interface{}

// New create an empty stack.
func New() Stack {
	return Stack{}
}

// Push an item onto the stack.
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// Clear a stack of all items.
func (s *Stack) Clear() {
	*s = Stack{}
}

// Peek return the top item without popping it off the stack. Return nil if the stack is empty.
func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

// Pop the top item off the stack. Return nil if the stack is empty.
func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

// Size returns the number of items on the stack.
func (s *Stack) Size() int {
	return len(*s)
}

// IsEmpty returns true iff stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
