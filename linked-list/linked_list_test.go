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
