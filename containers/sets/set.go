package sets

type a int

type Set[E comparable] struct {
	elems map[E]struct{}
}

func NewSet[E comparable]() *Set[E] {
	s := new(Set[E])
	s.init()
	return s
}

func (s *Set[E]) Size() int {
	return len(s.elems)
}

func (s *Set[E]) Empty() bool {
	return s.Size() == 0
}

func (s *Set[E]) Contains(e E) bool {
	_, exist := s.elems[e]
	return exist
}

func (s *Set[E]) Add(es ...E) {
	s.init()
	for _, e := range es {
		s.elems[e] = struct{}{}
	}
}

func (s *Set[E]) Remove(es ...E) {
	for _, e := range es {
		delete(s.elems, e)
	}
}

func (s *Set[E]) Clear() {
	clear(s.elems)
}

func (s *Set[E]) Elems() []E {
	slice := make([]E, 0, len(s.elems))
	for k := range s.elems {
		slice = append(slice, k)
	}
	return slice
}

func (s *Set[E]) init() {
	if s.elems != nil {
		return
	}
	s.elems = make(map[E]struct{})
}
