package stack

type Stack[T any] struct {
	Data []T
	Top  int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{Top: -1}
}
