package queues

type dequeNode[E any] struct {
	value      E
	prev, next *dequeNode[E]
}

type Deque[E any] struct {
	empty  *dequeNode[E]
	length int
}

func NewDeque[E any]() *Deque[E] {
	d := new(Deque[E])
	d.init()
	return d
}

func (d *Deque[E]) Len() int {
	return d.length
}

func (d *Deque[E]) Empty() bool {
	return d.Len() == 0
}

func (d *Deque[E]) Front() E {
	if d.Empty() {
		panic("empty queue")
	}
	return d.empty.next.value
}

func (d *Deque[E]) Back() E {
	if d.Empty() {
		panic("empty queue")
	}
	return d.empty.prev.value
}

func (d *Deque[E]) PushFront(e E) {
	d.init()
	newNode := &dequeNode[E]{
		value: e,
		prev:  d.empty,
		next:  d.empty.next,
	}
	d.empty.next.prev = newNode
	d.empty.next = newNode
	d.length++
}

func (d *Deque[E]) PushBack(e E) {
	d.init()
	newNode := &dequeNode[E]{
		value: e,
		prev:  d.empty.prev,
		next:  d.empty,
	}
	d.empty.prev.next = newNode
	d.empty.prev = newNode
	d.length++
}

func (d *Deque[E]) PopFront() E {
	if d.Empty() {
		panic("empty queue")
	}
	e := d.empty.next.value
	d.empty.next = d.empty.next.next
	d.empty.next.prev = d.empty
	d.length--
	return e
}

func (d *Deque[E]) PopBack() E {
	if d.Empty() {
		panic("empty queue")
	}
	e := d.empty.prev.value
	d.empty.prev = d.empty.prev.prev
	d.empty.prev.next = d.empty
	d.length--
	return e
}

func (d *Deque[E]) init() {
	if d.empty != nil {
		return
	}
	d.empty = &dequeNode[E]{}
	d.empty.prev = d.empty
	d.empty.next = d.empty
	d.length = 0
}
