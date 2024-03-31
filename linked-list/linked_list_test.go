package linkedlist

import "testing"

func TestAppendToEnd(t *testing.T) {
	lk := NewLinkedList[int]()
	lk.AppendToEnd(1)
	lk.AppendToEnd(2)
	lk.AppendToEnd(3)

	/*testing to check if the head data is the first number to be appended
	and the tail is the last appended number*/
	if lk.Head.Data != 1 && lk.Tail.Data != 3 {
		t.Fail()
	}
}

func TestAppendToBeggining(t *testing.T) {
	lk := NewLinkedList[string]()
	lk.AppendToBeggining("John")
	lk.AppendToBeggining("Smith")
	lk.AppendToBeggining("Doe")

	//check if the head value is the last one to be inserted
	if lk.Head.Data != "Doe" {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	lk := NewLinkedList[float32]()
	lk.Insert(10, 10.23)

	//test if data was inserted to head if head is nil
	if lk.Head.Data != 10.23 && lk.Tail.Data != 10.23 {
		t.Fail()
	}

	lk.AppendToEnd(9.5)
	lk.AppendToEnd(54.12)
	lk.AppendToEnd(87.2134)
	lk.Insert(1, 54.0)

	// check if data was inserted at index 1
	if lk.Head.Next.Data != 54.0 {
		current := lk.Head
		count := 0
		for current != nil {
			t.Logf("%d: %f\n", count, current.Data)
			current = current.Next
			count++
		}

		t.Fail()
	}

	//check if former index 1 was moved to index 2
	if lk.Head.Next.Next.Data != 9.5 {
		t.Fail()
	}

	lk.Insert(100, 8.7)

	//check if data was appended to tail if index is bigger than list length
	if lk.Tail.Data != 8.7 {
		t.Fail()
	}
}

func TestDeleteBeggining(t *testing.T) {
	lk := NewLinkedList[string]()

	lk.AppendToEnd("John")
	lk.AppendToEnd("Doe")
	lk.AppendToEnd("foo")
	lk.AppendToEnd("bar")

	lk.DeleteBeggining()

	if lk.Head.Data != "Doe" {
		t.Fail()
	}

	lk.DeleteBeggining()
	lk.DeleteBeggining()

	if lk.Head.Data != "bar" {
		t.Fail()
	}
}

func TestDeleteEnd(t *testing.T) {
	lk := NewLinkedList[int]()

	lk.AppendToEnd(10)
	lk.AppendToEnd(20)
	lk.AppendToEnd(30)
	lk.AppendToEnd(40)

	lk.DeleteEnd()

	if lk.Tail.Data != 30 {
		t.Fail()
	}

	lk.DeleteEnd()
	lk.DeleteEnd()

	if lk.Tail.Data != 10 {
		t.Fail()
	}
}
