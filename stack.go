package goalgo

import "errors"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	next *Element
	data int
}

func (s *Stack) Push(data int) {
	s.top = &Element{next: s.top, data: data}
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size > 0 {
		data := s.top.data
		s.top = s.top.next
		s.size--
		return data, nil
	}
	return 0, errors.New("Stack is empty")
}

func (s *Stack) Count() int {
	return s.size
}
