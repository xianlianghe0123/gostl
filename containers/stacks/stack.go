package stacks

import (
	"github.com/xianlianghe0123/gostl/containers/internal"
)

type

type Stack[E any] struct {
	list *internal.List[E]
}

func NewStack[E any]() *Stack[E] {
	return &Stack[E]{
		list: internal.NewList[E](),
	}
}

func (s *Stack[E]) Len() int {
	return s.list.Len()
}

func (s *Stack[E]) Empty() bool {
	return s.list.Len() == 0
}

func (s *Stack[E]) Push(elem E) {
	s.list.PushBack(elem)
}

func (s *Stack[E]) Pop() E {
	return s.list.Delete(s.list.Len() - 1)
}

func (s *Stack[E]) Top() E {
	return s.list.Get(s.list.Len() - 1)
}
