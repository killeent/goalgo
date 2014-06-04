/*
Package goalgo implements standard CS data structures and algorithms.
*/
package goalgo

// Stack is a LIFO data structure.
type Stack struct {
	top  *stackElement
	size int
}

type stackElement struct {
	next *stackElement
	data interface{}
}

// Push adds an stackElement to the top of the Stack.
func (s *Stack) Push(data interface{}) {
	s.top = &stackElement{next: s.top, data: data}
	s.size++
}

// Pop removes and returns the top stackElement in the Stack. If
// the Stack is empty, returns nil
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		data := s.top.data
		s.top = s.top.next
		s.size--
		return data
	}
	return nil
}

// Count returns the number of stackElements in the Stack.
func (s *Stack) Count() int {
	return s.size
}
