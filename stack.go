package calculator

import "os"

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(obj interface{}) {
	*s = append(*s, obj)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return os.DevNull, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return os.DevNull
	} else {
		index := len(*s) - 1
		return (*s)[index]
	}
}
