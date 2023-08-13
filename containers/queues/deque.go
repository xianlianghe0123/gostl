package queues

import (
	"container/list"
	"github.com/xianlianghe0123/gostl/containers/internal"
)

type Deque[E any] interface {
	Len() int
	Empty() bool
	PushBack(elem E)
	PushFront(elem E)
	PopFront() E
	PopBack() E
	Front() E
	Back() E
}

type deque[E any] struct {
	list *internal.QuickList[E]
}

func NewDeque[E any]() Deque[E] {
	return &deque[E]{
		list: internal.NewQuickList[E](),
	}
}

func (d *deque[E]) Len() int {
	list.New()
	return d.list.Len()
}

func (d *deque[E]) Empty() bool {
	return d.list.Len() == 0
}

func (d *deque[E]) PushBack(elem E) {
	d.list.PushBack(elem)
}

func (d *deque[E]) PushFront(elem E) {
	d.list.PushFront(elem)
}

func (d *deque[E]) PopFront() E {
	return d.list.Remove(0)
}

func (d *deque[E]) PopBack() E {
	return d.list.Remove(d.list.Len() - 1)
}

func (d *deque[E]) Front() E {
	return d.list.Get(0)
}

func (d *deque[E]) Back() E {
	return d.list.Get(d.list.Len() - 1)
}
