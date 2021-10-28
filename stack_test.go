package stack

import (
	"testing"
)

func TestEmptySize(t *testing.T) {
	s := New()

	if s.Size() != 0 {
		t.FailNow()
	}

}

func TestEmptyPop(t *testing.T) {
	s := New()

	if s.Pop() != nil {
		t.FailNow()
	}
}

func TestEmptyPeek(t *testing.T) {
	s := New()

	if s.Peek() != nil {
		t.FailNow()
	}
}

func TestEmptyIsEmpty(t *testing.T) {
	s := New()

	if !s.IsEmpty() {
		t.FailNow()
	}
}

func TestOneSize(t *testing.T) {
	s := New()
	s.Push(1)

	if s.Size() != 1 {
		t.FailNow()
	}
}

func TestOnePop(t *testing.T) {
	s := New()
	s.Push(1)

	if s.Pop() != 1 {
		t.FailNow()
	}

	if s.Size() != 0 {
		t.FailNow()
	}
}

func TestOnePeek(t *testing.T) {
	s := New()
	s.Push(1)

	if s.Peek() != 1 {
		t.FailNow()
	}

	if s.Size() != 1 {
		t.FailNow()
	}
}

func TestOneIsEmpty(t *testing.T) {
	s := New()
	s.Push(1)

	if s.IsEmpty() {
		t.FailNow()
	}
}

func TestMultipleOperations(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push("fred")

	a := s.Pop()
	if a != "fred" {
		t.FailNow()
	}

	s.Push("barney")
	if s.IsEmpty() {
		t.FailNow()
	}
	if s.Size() != 3 {
		t.FailNow()
	}
	if s.Peek() != "barney" {
		t.FailNow()
	}
	s.Push([]byte{1, 2, 3})
	_ = s.Pop()
	if s.Pop() != "barney" {
		t.FailNow()
	}

	if s.Pop() != 2 {
		t.FailNow()
	}

	if s.IsEmpty() {
		t.FailNow()
	}

	if s.Pop() != 1 {
		t.FailNow()
	}

	if s.Pop() != nil {
		t.FailNow()
	}

	if !s.IsEmpty() {
		t.FailNow()
	}

	s.Push("equals")

	if s.Peek() != "equals" {
		t.FailNow()
	}

	if s.Size() != 1 {
		t.FailNow()
	}

	s.Clear()

	if !s.IsEmpty() {
		t.FailNow()
	}

	s.Push("Mork")

	if s.Peek() != "Mork" {
		t.FailNow()
	}

	if s.Size() != 1 {
		t.FailNow()
	}

}
