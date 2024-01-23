package stacks

type Stack[E any] struct {
	elems []E
}

func NewStack[E any](capacity int) *Stack[E] {
	return &Stack[E]{
		elems: make([]E, 0, capacity),
	}
}

func (s *Stack[E]) Size() int {
	return len(s.elems)
}

func (s *Stack[E]) Empty() bool {
	return s.Size() == 0
}

func (s *Stack[E]) Top() E {
	return s.elems[s.Size()-1]
}

func (s *Stack[E]) Push(e E) {
	s.elems = append(s.elems, e)
}

func (s *Stack[E]) Pop() E {
	e := s.elems[s.Size()-1]
	s.elems = s.elems[:s.Size()-1]
	return e
}
