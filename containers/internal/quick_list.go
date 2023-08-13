package internal

var blockCap = 64

type QuickListNode[E any] struct {
	*Block[E]
	prev, next *QuickListNode[E]
}

func NewListNode[E any](prev, next *QuickListNode[E]) *QuickListNode[E] {
	return &QuickListNode[E]{
		Block: NewBlock[E](blockCap),
		prev:  prev,
		next:  next,
	}
}

type QuickList[E any] struct {
	head, tail *QuickListNode[E]
	len        int
}

func NewQuickList[E any]() *QuickList[E] {
	node := NewListNode[E](nil, nil)
	return &QuickList[E]{
		head: node,
		tail: node,
		len:  0,
	}
}

func (l *QuickList[E]) Len() int {
	return l.len
}

func (l *QuickList[E]) Get(idx int) E {
	node, idx := l.find(idx)
	return node.Get(idx)
}

func (l *QuickList[E]) Set(idx int, elem E) {
	node, idx := l.find(idx)
	node.Set(idx, elem)
}

func (l *QuickList[E]) Remove(idx int) E {
	node, idx := l.find(idx)
	e := node.Remove(idx)
	l.len--
	// remove the empty node
	if node.Len() == 0 && l.head != l.tail {
		l.deleteNode(node)
	}
	return e
}

func (l *QuickList[E]) Insert(idx int, elem E) {
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

func (l *QuickList[E]) PushBack(elem E) {
	if l.tail.Full() {
		l.insertNodeAfter(l.tail)
	}
	l.tail.Insert(l.tail.Len(), elem)
	l.len++
}

func (l *QuickList[E]) PushFront(elem E) {
	if l.head.Full() {
		l.insertNodeBefore(l.head)
	}
	l.head.Insert(0, elem)
	l.len++
}

func (l *QuickList[E]) Begin() *QuickListIterator[E] {
	return &QuickListIterator[E]{
		listNode:      l.head,
		blockIterator: l.head.Begin(),
	}
}

func (l *QuickList[E]) End() *QuickListIterator[E] {
	return &QuickListIterator[E]{
		listNode:      l.tail,
		blockIterator: l.tail.End(),
	}
}

func (l *QuickList[E]) rangeCheck(idx int) {
	if idx < 0 || idx >= l.len {
		panic("out of range")
	}
}

func (l *QuickList[E]) find(idx int) (*QuickListNode[E], int) {
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

func (l *QuickList[E]) divide(node *QuickListNode[E]) {
	l.insertNodeAfter(node)
	next := node.next
	half := node.Len() / 2
	copy(next.elems[:], node.elems[half:])
	next.len = node.Len() - half
	node.len = half
}

func (l *QuickList[E]) insertNodeAfter(node *QuickListNode[E]) {
	newNode := NewListNode[E](node, node.next)
	if node.next != nil {
		node.next.prev = newNode
	} else {
		l.tail = newNode
	}
	node.next = newNode
}

func (l *QuickList[E]) insertNodeBefore(node *QuickListNode[E]) {
	newNode := NewListNode[E](node.prev, node)
	if node.prev != nil {
		node.prev.next = newNode
	} else {
		l.head = newNode
	}
	node.prev = newNode
}

func (l *QuickList[E]) deleteNode(node *QuickListNode[E]) {
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

type QuickListIterator[E any] struct {
	listNode      *QuickListNode[E]
	blockIterator *BlockIterator[E]
}

func NewListQuickIterator[E any]() *QuickListIterator[E] {
	return &QuickListIterator[E]{}
}

func (l *QuickListIterator[E]) HasNext() bool {
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

func (l *QuickListIterator[E]) Next() E {
	if !l.HasNext() {
		var e E
		return e
	}
	return l.blockIterator.Next()
}

func (l *QuickListIterator[E]) HasPrev() bool {
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

func (l *QuickListIterator[E]) Prev() E {
	if !l.HasPrev() {
		var e E
		return e
	}
	return l.blockIterator.Prev()
}
