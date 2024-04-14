package queue

import "errors"

type Queue[T any] struct {
	Data []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		Data: make([]T, 0),
	}
}

func (q *Queue[T]) EnQueue(data T) {
	q.Data = append(q.Data, data)
}

func (q *Queue[T]) DeQueue() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("Queue is empty")
	}

	removed := q.Data[0]
	q.Data = q.Data[1:]
	return removed, nil
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.Data) == 0
}
