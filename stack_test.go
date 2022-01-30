package brainfuck

import "testing"

var testPushPop1 = []int{10, 20, 30, 40, 50}
var testPushPop2 = []int{10, 20, 30, 40, 50, 60, 70, 80}

func TestStack_Push(t *testing.T) {
	for _, tt := range []struct {
		input        []int
		wantedLength int
	}{
		{testPushPop1, 5},
		{testPushPop2, 8},
	} {
		s := Stack{}
		for _, item := range tt.input {
			s.Push(item)
		}
		if s.Len() != tt.wantedLength {
			t.Errorf("wrong stack length, got %d", s.Len())
		}
	}
}
func TestStack_Pop(t *testing.T) {
	for _, tt := range []struct {
		input        []int
		wantedLength int
	}{
		{testPushPop1, 5},
		{testPushPop2, 8},
	} {
		s := Stack{}
		for _, item := range tt.input {
			s.Push(item)
		}
		for i, _ := range tt.input {
			pop := s.Pop()
			if pop != tt.input[len(tt.input)-i-1] {
				t.Errorf("wrong value, want: %d, got %s", tt.input[len(tt.input)-i-1], pop)
			}
		}
	}

}
