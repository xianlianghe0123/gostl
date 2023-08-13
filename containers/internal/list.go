package internal

var blockCap = 64

type ListNode[E any] struct {
	*Block[E]
	prev, next *ListNode[E]
}

func NewListNode[E any](prev, next *ListNode[E]) *ListNode[E] {
	return &ListNode[E]{
		Block: NewBlock[E](blockCap),
		prev:  prev,
		next:  next,
	}
}

type List[E any] struct {
	head, tail *ListNode[E]
	len        int
}

func NewList[E any]() *List[E] {
	node := NewListNode[E](nil, nil)
	return &List[E]{
		head: node,
		tail: node,
		len:  0,
	}
}

func (l *List[E]) Len() int {
	return l.len
}

func (l *List[E]) Get(idx int) E {
	node, idx := l.find(idx)
	return node.Get(idx)
}

func (l *List[E]) Set(idx int, elem E) {
	node, idx := l.find(idx)
	node.Set(idx, elem)
}

func (l *List[E]) Delete(idx int) E {
	node, idx := l.find(idx)
	e := node.Delete(idx)
	l.len--
	// remove the empty node
	if node.Len() == 0 && l.head != l.tail {
		l.deleteNode(node)
	}
	return e
}

func (l *List[E]) Insert(idx int, elem E) {
	if idx < 0 || idx > l.len {
		panic("out of range")
	}
	switch idx {
	case 0:
		l.PushFront(elem)
	case l.len:
		l.PushBack(elem)
	default:
		node, idx := l.find(idx)
		if node.Full() {
			l.divide(node)
			if idx > node.Len() {
				node = node.next
			}
		}
		node.Insert(idx, elem)
		l.len++
	}
}

func (l *List[E]) PushBack(elem E) {
	if l.tail.Full() {
		l.insertNodeAfter(l.tail)
	}
	l.tail.Insert(l.tail.Len(), elem)
	l.len++
}

func (l *List[E]) PushFront(elem E) {
	if l.head.Full() {
		l.insertNodeBefore(l.head)
	}
	l.head.Insert(0, elem)
	l.len++
}

func (l *List[E]) Begin() *ListIterator[E] {
	return &ListIterator[E]{
		listNode:      l.head,
		blockIterator: l.head.Begin(),
	}
}

func (l *List[E]) End() *ListIterator[E] {
	return &ListIterator[E]{
		listNode:      l.tail,
		blockIterator: l.tail.End(),
	}
}

func (l *List[E]) rangeCheck(idx int) {
	if idx < 0 || idx >= l.len {
		panic("out of range")
	}
}

func (l *List[E]) find(idx int) (*ListNode[E], int) {
	l.rangeCheck(idx)
	if idx < l.len/2 {
		cur := l.head
		for idx >= cur.Len() {
			idx -= cur.Len()
			cur = cur.next
		}
		return cur, idx
	}
	idx = l.len - idx
	cur := l.tail
	for idx > cur.Len() {
		idx -= cur.Len()
		cur = cur.prev
	}
	return cur, cur.Len() - idx
}

func (l *List[E]) divide(node *ListNode[E]) {
	l.insertNodeAfter(node)
	next := node.next
	half := node.Len() / 2
	copy(next.elems[:], node.elems[half:])
	next.len = node.Len() - half
	node.len = half
}

func (l *List[E]) insertNodeAfter(node *ListNode[E]) {
	newNode := NewListNode[E](node, node.next)
	if node.next != nil {
		node.next.prev = newNode
	} else {
		l.tail = newNode
	}
	node.next = newNode
}

func (l *List[E]) insertNodeBefore(node *ListNode[E]) {
	newNode := NewListNode[E](node.prev, node)
	if node.prev != nil {
		node.prev.next = newNode
	} else {
		l.head = newNode
	}
	node.prev = newNode
}

func (l *List[E]) deleteNode(node *ListNode[E]) {
	if node.prev == nil {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
}

type ListIterator[E any] struct {
	listNode      *ListNode[E]
	blockIterator *BlockIterator[E]
}

func NewListIterator[E any]() *ListIterator[E] {
	return &ListIterator[E]{}
}

func (l *ListIterator[E]) HasNext() bool {
	if l.blockIterator.HasNext() {
		return true
	}
	if l.listNode.next == nil {
		return false
	}
	l.listNode = l.listNode.next
	l.blockIterator = l.listNode.Begin()
	return l.blockIterator.HasNext()
}

func (l *ListIterator[E]) Next() E {
	if !l.HasNext() {
		var e E
		return e
	}
	return l.blockIterator.Next()
}

func (l *ListIterator[E]) HasPrev() bool {
	if l.blockIterator.HasPrev() {
		return true
	}
	if l.listNode.prev == nil {
		return false
	}
	l.listNode = l.listNode.prev
	l.blockIterator = l.listNode.End()
	return l.blockIterator.HasPrev()
}

func (l *ListIterator[E]) Prev() E {
	if !l.HasPrev() {
		var e E
		return e
	}
	return l.blockIterator.Prev()
}
