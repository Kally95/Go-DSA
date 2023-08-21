package data_structures

type Stack[T any] struct {
	StackNodes []T
	Length     int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.StackNodes = append(s.StackNodes[:s.Length], v)
	s.Length++
}

func (s *Stack[T]) Pop() (t T) {
	if s.Length == 0 {
		return t
	}
	s.Length--
	return s.StackNodes[s.Length]
}
