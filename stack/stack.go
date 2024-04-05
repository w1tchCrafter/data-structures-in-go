package stack

import "errors"

var EmptyStack error = errors.New("Stack is empty")

type Stack[T any] struct {
	Data []T
	Top  int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		Data: make([]T, 0),
		Top:  -1,
	}
}

// insert element into the stack
func (s *Stack[T]) Push(data T) {
	s.Data = append(s.Data, data)
	s.Top++
}

// remove element from the stack then return it
func (s *Stack[T]) Pop() (T, error) {
	if !s.IsEmpty() {
		e := s.Data[s.Top]
		s.Data = s.Data[0:s.Top]
		s.Top--
		return e, nil
	}

	return *new(T), EmptyStack
}

// returns the top element of the stack
func (s *Stack[T]) Peek() (T, error) {
	if !s.IsEmpty() {
		return s.Data[s.Top], nil
	}

	return *new(T), EmptyStack
}

// returns true if top element == -1 (initial value)
func (s *Stack[T]) IsEmpty() bool {
	return s.Top == -1
}

// returns the number of elements currently on the stack
func (s *Stack[T]) Size() int {
	return len(s.Data)
}
