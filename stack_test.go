package brainfuck

import "testing"

func TestStack_Push(t *testing.T) {
	s := Stack{}
	s.Push(101)
	s.Push("simple text value ")
	s.Push(1.95)
	if s.Len() != 3 {
		t.Errorf("wrong stack length, got %d", s.Len())
	}
}

func TestStack_Pop(t *testing.T) {
	s := Stack{}
	s.Push(101)
	s.Push("simple text value")
	s.Push(1.95)
	pop1 := s.Pop()
	if pop1 != 1.95 {
		t.Errorf("wrong value, got %d", pop1)
	}

	pop2 := s.Pop()
	if pop2 != "simple text value" {
		t.Errorf("wrong value, got %s", pop2)
	}
}
