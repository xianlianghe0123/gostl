package stacks

import (
	"github.com/xianlianghe0123/gostl/containers/internal"
)

type Stack[E any] interface {
	Size() int
	Empty() bool
	Push(elem E)
	Pop() E
	Top() E
}

type stack[E any] struct {
	list *internal.QuickList[E]
}

func NewStack[E any]() Stack[E] {
	return &stack[E]{
		list: internal.NewQuickList[E](),
	}
}

func (s *stack[E]) Size() int {
	return s.list.Len()
}

func (s *stack[E]) Empty() bool {
	return s.list.Len() == 0
}

func (s *stack[E]) Push(elem E) {
	s.list.PushBack(elem)
}

func (s *stack[E]) Pop() E {
	return s.list.Remove(s.list.Len() - 1)
}

func (s *stack[E]) Top() E {
	return s.list.Get(s.list.Len() - 1)
}
