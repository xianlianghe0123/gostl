package queues

type queueNode[E any] struct {
	value E
	next  *queueNode[E]
}

type Queue[E any] struct {
	head   *queueNode[E]
	tail   *queueNode[E]
	length int
}

func NewQueue[E any]() *Queue[E] {
	q := new(Queue[E])
	q.init()
	return q
}

func (q *Queue[E]) Len() int {
	return q.length
}

func (q *Queue[E]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[E]) Front() E {
	if q.Empty() {
		panic("empty queue")
	}
	return q.head.next.value
}

func (q *Queue[E]) Back() E {
	if q.Empty() {
		panic("empty queue")
	}
	return q.tail.value
}

func (q *Queue[E]) Push(e E) {
	q.init()
	q.tail.next = &queueNode[E]{
		value: e,
		next:  nil,
	}
	q.tail = q.tail.next
	q.length++
}

func (q *Queue[E]) Pop() E {
	if q.Empty() {
		panic("empty queue")
	}
	e := q.head.next.value
	q.head.next = q.head.next.next
	q.length--
	return e
}

func (q *Queue[E]) init() {
	if q.head != nil {
		return
	}
	emptyHeadNode := &queueNode[E]{}
	q.head = emptyHeadNode
	q.tail = emptyHeadNode
	q.length = 0
}
