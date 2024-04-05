package stack

import "testing"

func TestPush(t *testing.T) {
	s := NewStack[string]()
	s.Push("John")
	s.Push("Doe")
	s.Push("foo")
	s.Push("bar")

	if s.Top != 3 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// due to fifo principle the popped number must be 3
	popped := s.Pop()

	if popped != 3 && s.Top != 1 && s.Data[s.Top] != 2 {
		t.Fail()
	}
}

func TestTopElement(t *testing.T) {
	s := NewStack[int]()
	s.Push(23)
	s.Push(47)

	te := s.TopElement()

	if te != 47 {
		t.Fail()
	}

	s.Pop()
	te = s.TopElement()

	if te != 23 {
		t.Fail()
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack[string]()

	if !s.IsEmpty() {
		t.Fail()
	}

	s.Push("bruh")

	if s.IsEmpty() {
		t.Fail()
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	s := NewStack[float32]()

	if s.Size() != 0 {
		t.Fail()
	}

	s.Push(43.6)
	s.Push(764)
	s.Push(32.54)

	if s.Size() != 3 {
		t.Fail()
	}

	s.Pop()
	s.Pop()

	if s.Size() != 1 {
		t.Fail()
	}
}
