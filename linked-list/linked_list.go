package linkedlist

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (lk *LinkedList[T]) AppendToEnd(data T) {
	NewNode := &Node[T]{Data: data}

	if lk.Head == nil {
		lk.Head = NewNode
	} else {
		lk.Tail.Next = NewNode
	}

	lk.Tail = NewNode
}

func (lk *LinkedList[T]) AppendToBeggining(data T) {
	NewNode := &Node[T]{Data: data}

	if lk.Head == nil {
		lk.Head = NewNode
	} else {
		NewNode.Next = lk.Head
	}

	lk.Head = NewNode
}
