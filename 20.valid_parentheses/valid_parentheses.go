package valid_parentheses

import (
	"fmt"
)

func IsValid(s string) bool {
	stack := Stack{}

	for i := range s {
		switch s[i] {
		case '(':
			stack.Push('(')
		case '{':
			stack.Push('{')
		case '[':
			stack.Push('[')
		case ')':
			if stack.Peek() != '(' {
				return false
			}
			fmt.Println(string(stack.Pop()))
		case '}':
			if stack.Peek() != '{' {
				return false
			}
			fmt.Println(string(stack.Pop()))

		case ']':
			if stack.Peek() != '[' {
				return false
			}
			fmt.Println(string(stack.Pop()))

		}

	}

	return stack.IsEmpty()
}

type Stack struct {
	data []byte
}

func (s *Stack) Push(item byte) {
	s.data = append(s.data, item)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Peek() byte {
	if s.IsEmpty() {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return 0
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item
}
