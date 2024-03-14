package reversepolis

import "strconv"

func EvalRPN(tokens []string) int {

	// * / - +

	stack := Stack{}

	for _, token := range tokens {
		switch token {
		case "+":
			a := stack.Pop()
			b := stack.Pop()
			b += a
			stack.Push(b)
		case "-":
			a := stack.Pop()
			b := stack.Pop()
			b -= a
			stack.Push(b)
		case "/":
			a := stack.Pop()
			b := stack.Pop()
			b /= a
			stack.Push(b)
		case "*":
			a := stack.Pop()
			b := stack.Pop()
			b *= a
			stack.Push(b)
		default:
			val, _ := strconv.Atoi(token)
			stack.Push(val)
		}
	}

	return stack.Peek()
}

type Stack struct {
	Val []int
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	temp := s.Peek()
	s.Val = s.Val[:len(s.Val)-1]
	return temp
}

func (s *Stack) IsEmpty() bool {
	return len(s.Val) == 0
}

func (s *Stack) Push(v int) {
	s.Val = append(s.Val, v)
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return 0
	}

	return s.Val[len(s.Val)-1]
}
