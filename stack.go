/*
Package goalgo implements standard CS data structures and algorithms.
*/
package goalgo

import "errors"

// Stack is a FIFO data structure.
type Stack struct {
	top  *element
	size int
}

type element struct {
	next *element
	data interface{}
}

// Push adds an element to the top of the Stack.
func (s *Stack) Push(data interface{}) {
	s.top = &element{next: s.top, data: data}
	s.size++
}

// Pop removes and returns the top element in the Stack. If
// the Stack is empty, returns a non-nil error.
func (s *Stack) Pop() (interface{}, error) {
	if s.size > 0 {
		data := s.top.data
		s.top = s.top.next
		s.size--
		return data, nil
	}
	return 0, errors.New("Stack is empty")
}

// Count returns the number of elements in the Stack.
func (s *Stack) Count() int {
	return s.size
}
