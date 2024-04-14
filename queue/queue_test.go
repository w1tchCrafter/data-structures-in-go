package queue

import "testing"

func TestEnQueue(t *testing.T) {
	q := NewQueue[int]()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	if q.Data[0] != 1 || q.Data[1] != 2 || q.Data[2] != 3 || q.IsEmpty() {
		t.Fail()
	}
}

func TestDeQueue(t *testing.T) {
	q := NewQueue[int]()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	first, err := q.DeQueue()

	if first != 1 || err != nil {
		t.Fail()
	}

	second, err := q.DeQueue()

	if second != 2 || err != nil {
		t.Fail()
	}
}
