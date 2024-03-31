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

// append new data to the end of the linked list
func (lk *LinkedList[T]) AppendToEnd(data T) {
	NewNode := &Node[T]{Data: data}

	if lk.Head == nil {
		lk.Head = NewNode
	} else {
		lk.Tail.Next = NewNode
	}

	lk.Tail = NewNode
}

// append new data to the beggining of the linked list
func (lk *LinkedList[T]) AppendToBeggining(data T) {
	NewNode := &Node[T]{Data: data}
	NewNode.Next = lk.Head
	lk.Head = NewNode
}

/*
insert data at a certain index starting at zero
if head is nil, appends data to head
if the linked list is not long enough, appends the data to the tail
*/
func (lk *LinkedList[T]) Insert(index uint, data T) {
	NewNode := &Node[T]{Data: data}

	if lk.Head == nil {
		lk.Head = NewNode
		lk.Tail = NewNode
		return
	}

	current := lk.Head
	for i := range index {
		if current == nil {
			current = NewNode
			lk.Tail = NewNode
			break
		}

		if i == index-1 {
			NewNode.Next = current.Next
			current.Next = NewNode
			break
		}

		current = current.Next
	}
}

// delete the first node from the linked list
func (lk *LinkedList[T]) DeleteBeggining() {
	if lk.Head == nil {
		return
	}

	lk.Head = lk.Head.Next
}

// delete the last item from the linked list
func (lk *LinkedList[T]) DeleteEnd() {
	if lk.Head == nil {
		return
	}

	if lk.Head == lk.Tail {
		lk.Head = nil
		lk.Tail = nil
	}

	current := lk.Head
	for current.Next != lk.Tail {
		current = current.Next
	}

	current.Next = nil
	lk.Tail = current
}
