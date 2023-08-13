package lists

import (
	"github.com/xianlianghe0123/gostl/containers/internal"
)

type Queue[E any] struct {
	list *internal.List[E]
}

func NewQueue[E any]() *Queue[E] {
	return &Queue[E]{
		list: internal.NewList[E](),
	}
}

func (q *Queue[E]) Len() int {
	return q.list.Len()
}

func (q *Queue[E]) Empty() bool {
	return q.list.Len() == 0
}

func (q *Queue[E]) Push(elem E) {
	q.list.PushBack(elem)
}

func (q *Queue[E]) Front() E {
	return q.list.Get(0)
}
