package brainfuck

import "sync"

// Default data type for stack
type Item interface{}

// Stack data structure
type Stack struct {
	items []Item
	mutex sync.RWMutex
}

func (s *Stack) Len() int {
	return len(s.items)
}

// Push the new data to the top of the stack
func (s *Stack) Push(t Item) {
	s.mutex.Lock()
	s.items = append(s.items, t)
	s.mutex.Unlock()
}

// Pop the last item from top of the stack
func (s *Stack) Pop() Item {
	s.mutex.Lock()
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.mutex.Unlock()
	return top

}

// Create new instance of stack
func (s *Stack) NewStack() *Stack {
	s.items = []Item{}
	return s
}
