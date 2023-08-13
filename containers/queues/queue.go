package queues

import (
	"github.com/xianlianghe0123/gostl/containers/internal"
)

type Queue[E any] interface {
	Len() int
	Empty() bool
	Push(elem E)
	Pop() E
	Front() E
}

type queue[E any] struct {
	list *internal.QuickList[E]
}

func NewQueue[E any]() Queue[E] {
	return &queue[E]{
		list: internal.NewQuickList[E](),
	}
}

func (q *queue[E]) Len() int {
	return q.list.Len()
}

func (q *queue[E]) Empty() bool {
	return q.list.Len() == 0
}

func (q *queue[E]) Push(elem E) {
	q.list.PushBack(elem)
}

func (q *queue[E]) Pop() E {
	return q.list.Remove(0)
}

func (q *queue[E]) Front() E {
	return q.list.Get(0)
}
