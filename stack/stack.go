package stack

type Stack struct {
	Data []any
}

// Push adds an element to the stack
func (s *Stack) Push(element any) {
	s.Data = append(s.Data, element)
}

// Pop removes and returns the top element, or nil if empty
func (s *Stack) Pop() any {
	if len(s.Data) == 0 {
		return nil
	}
	last := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return last
}

// Peek returns the top element without removing it
func (s *Stack) Peek() any {
	if len(s.Data) == 0 {
		return nil
	}
	return s.Data[len(s.Data)-1]
}

// Clear removes all elements
func (s *Stack) Clear() {
	s.Data = nil
}

// Length returns the number of elements
func (s *Stack) Length() int {
	return len(s.Data)
}
