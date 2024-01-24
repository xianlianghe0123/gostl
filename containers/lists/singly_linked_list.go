package lists

type SinglyLinkedListNode[E any] struct {
	Value E
	next  *SinglyLinkedListNode[E]
	list  *SinglyLinkedList[E]
}

func (l *SinglyLinkedListNode[E]) Next() *SinglyLinkedListNode[E] {
	return l.next
}

type SinglyLinkedList[E any] struct {
	head   *SinglyLinkedListNode[E]
	tail   *SinglyLinkedListNode[E]
	length int
}

func NewSinglyLinkedList[E any]() *SinglyLinkedList[E] {
	s := new(SinglyLinkedList[E])
	s.init()
	return s
}

func (s *SinglyLinkedList[E]) Len() int {
	return s.length
}

func (s *SinglyLinkedList[E]) Empty() bool {
	return s.Len() == 0
}

func (s *SinglyLinkedList[E]) Front() *SinglyLinkedListNode[E] {
	if s.Empty() {
		return nil
	}
	return s.head.next
}

func (s *SinglyLinkedList[E]) Back() *SinglyLinkedListNode[E] {
	if s.Empty() {
		return nil
	}
	return s.tail
}

func (s *SinglyLinkedList[E]) InsertAfter(node *SinglyLinkedListNode[E], e E) *SinglyLinkedListNode[E] {
	s.checkNode(node)
	return s.insertAfter(node, e)
}

func (s *SinglyLinkedList[E]) InsertBefore(node *SinglyLinkedListNode[E], e E) *SinglyLinkedListNode[E] {
	s.checkNode(node)
	prev := s.searchPrev(node)
	return s.insertAfter(prev, e)
}

func (s *SinglyLinkedList[E]) Remove(node *SinglyLinkedListNode[E]) E {
	s.checkNode(node)
	return s.remove(node)
}

func (s *SinglyLinkedList[E]) PushFront(e E) {
	s.init()
	s.insertAfter(s.head, e)
}

func (s *SinglyLinkedList[E]) PushBack(e E) {
	s.init()
	s.insertAfter(s.tail, e)
}

func (s *SinglyLinkedList[E]) PopFront() E {
	if s.Empty() {
		panic("list is root")
	}
	return s.remove(s.head.next)
}

func (s *SinglyLinkedList[E]) PopBack() E {
	if s.Empty() {
		panic("list is root")
	}
	return s.remove(s.tail)
}

func (s *SinglyLinkedList[E]) MoveFront(node *SinglyLinkedListNode[E]) {
	s.checkNode(node)
	if node == s.Front() {
		return
	}

	s.removeNodeOnly(node)

	node.next = s.head.next
	s.head.next = node
}

func (s *SinglyLinkedList[E]) MoveBack(node *SinglyLinkedListNode[E]) {
	s.checkNode(node)
	if node == s.tail {
		return
	}

	s.removeNodeOnly(node)

	node.next = nil
	s.tail.next = node
	s.tail = s.tail.next
}

func (s *SinglyLinkedList[E]) init() {
	if s.head != nil {
		return
	}
	emptyHeadNode := &SinglyLinkedListNode[E]{list: s}
	s.head = emptyHeadNode
	s.tail = emptyHeadNode
	s.length = 0
}

func (s *SinglyLinkedList[E]) insertAfter(node *SinglyLinkedListNode[E], e E) *SinglyLinkedListNode[E] {
	node.next = &SinglyLinkedListNode[E]{
		Value: e,
		next:  node.next,
		list:  s,
	}
	if s.tail == node {
		s.tail = s.tail.next
	}
	s.length++
	return node.next
}

func (s *SinglyLinkedList[E]) remove(node *SinglyLinkedListNode[E]) E {
	s.removeNodeOnly(node)
	node.list = nil
	node.next = nil
	s.length--
	return node.Value
}
func (s *SinglyLinkedList[E]) removeNodeOnly(node *SinglyLinkedListNode[E]) {
	prev := s.searchPrev(node)
	prev.next = prev.next.next
	if s.tail == node {
		s.tail = prev
	}
}

func (s *SinglyLinkedList[E]) searchPrev(node *SinglyLinkedListNode[E]) *SinglyLinkedListNode[E] {
	iter := s.head
	for ; iter != nil; iter = iter.Next() {
		if iter.next == node {
			break
		}
	}
	return iter
}

func (s *SinglyLinkedList[E]) checkNode(node *SinglyLinkedListNode[E]) {
	if node == nil {
		panic("nil node")
	}
	if node.list != s {
		panic("not belong to list")
	}
}
