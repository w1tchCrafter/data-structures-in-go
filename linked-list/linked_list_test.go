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
}
