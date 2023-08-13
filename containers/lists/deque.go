package lists

import "github.com/xianlianghe0123/gostl/containers/internal"

type Deque[E any] struct {
	list *internal.List[E]
}

func NewDeque[E any]() *Deque[E] {
	return &Deque[E]{
		list: internal.NewList[E](),
	}
}

func (d *Deque[E]) Len() int {
	return d.list.Len()
}

func (d *Deque[E]) Empty() bool {
	return d.list.Len() == 0
}

func (d *Deque[E]) PushBack(elem E) {
	d.list.PushBack(elem)
}

func (d *Deque[E]) PushFront(elem E) {
	d.list.PushFront(elem)
}

func (d *Deque[E]) PopFront() E {
	return d.list.Delete(0)
}

func (d *Deque[E]) PopBack() E {
	return d.list.Delete(d.list.Len() - 1)
}

func (d *Deque[E]) Front() E {
	return d.list.Get(0)
}

func (d *Deque[E]) Back() E {
	return d.list.Get(d.list.Len() - 1)
}
