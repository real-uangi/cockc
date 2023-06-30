package stack

import "errors"

// Stack 栈
type Stack struct {
	items []interface{}
	size  int
}

func New() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
		size:  0,
	}
}

func (s *Stack) IsEmpty() bool {
	return 0 == s.size
}

func (s *Stack) GetSize() int {
	return s.size
}

func (s *Stack) Pop() (data interface{}, err error) {
	if s.IsEmpty() {
		err = errors.New("empty stack")
	}
	data = s.items[s.size]
	s.items = s.items[:s.size]
	s.size--
	return data, err
}

func (s *Stack) Push(data interface{}) {
	s.items = append(s.items, data)
	s.size++
}

func (s *Stack) PeekNext() (data interface{}, err error) {
	if s.IsEmpty() {
		err = errors.New("empty stack")
	}
	return s.items[0], err
}
