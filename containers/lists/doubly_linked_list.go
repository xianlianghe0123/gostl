package lists

type DoublyLinkedListNode[E any] struct {
	Value      E
	prev, next *DoublyLinkedListNode[E]
	list       *DoublyLinkedList[E]
}

func (d *DoublyLinkedListNode[E]) Prev() *DoublyLinkedListNode[E] {
	if d.list == nil || d.prev == d.list.root {
		return nil
	}
	return d.prev
}

func (d *DoublyLinkedListNode[E]) Next() *DoublyLinkedListNode[E] {
	if d.list == nil || d.next == d.list.root {
		return nil
	}
	return d.next
}

type DoublyLinkedList[E any] struct {
	root   *DoublyLinkedListNode[E]
	length int
}

func NewDoublyLinkedList[E any]() *DoublyLinkedList[E] {
	d := new(DoublyLinkedList[E])
	d.init()
	return d
}

func (d *DoublyLinkedList[E]) Len() int {
	return d.length
}

func (d *DoublyLinkedList[E]) Empty() bool {
	return d.Len() == 0
}

func (d *DoublyLinkedList[E]) Front() *DoublyLinkedListNode[E] {
	if d.Empty() {
		return nil
	}
	return d.root.next
}

func (d *DoublyLinkedList[E]) Back() *DoublyLinkedListNode[E] {
	if d.Empty() {
		return nil
	}
	return d.root.prev
}

func (d *DoublyLinkedList[E]) InsertAfter(node *DoublyLinkedListNode[E], e E) *DoublyLinkedListNode[E] {
	d.checkNode(node)
	return d.insertAfter(node, e)
}

func (d *DoublyLinkedList[E]) InsertBefore(node *DoublyLinkedListNode[E], e E) *DoublyLinkedListNode[E] {
	d.checkNode(node)
	return d.insertAfter(node.prev, e)
}

func (d *DoublyLinkedList[E]) Remove(node *DoublyLinkedListNode[E]) E {
	d.checkNode(node)
	return d.remove(node)
}

func (d *DoublyLinkedList[E]) PushFront(e E) {
	d.init()
	d.insertAfter(d.root, e)
}

func (d *DoublyLinkedList[E]) PushBack(e E) {
	d.init()
	d.insertAfter(d.root.prev, e)
}

func (d *DoublyLinkedList[E]) PopFront() E {
	if d.Empty() {
		panic("list is root")
	}
	return d.remove(d.root.next)
}

func (d *DoublyLinkedList[E]) PopBack() E {
	if d.Empty() {
		panic("list is root")
	}
	return d.remove(d.root.prev)
}

func (d *DoublyLinkedList[E]) MoveFront(node *DoublyLinkedListNode[E]) {
	d.checkNode(node)
	d.removeNodeOnly(node)
	d.insertNodeOnly(d.root, node)
}

func (d *DoublyLinkedList[E]) MoveBack(node *DoublyLinkedListNode[E]) {
	d.checkNode(node)
	d.removeNodeOnly(node)
	d.insertNodeOnly(d.root.prev, node)
}

func (d *DoublyLinkedList[E]) init() {
	if d.root != nil {
		return
	}
	d.root = &DoublyLinkedListNode[E]{list: d}
	d.root.next = d.root
	d.root.prev = d.root
	d.length = 0
}

func (d *DoublyLinkedList[E]) insertAfter(node *DoublyLinkedListNode[E], e E) *DoublyLinkedListNode[E] {
	d.insertNodeOnly(node, &DoublyLinkedListNode[E]{Value: e})
	d.length++
	return node.next
}

func (d *DoublyLinkedList[E]) insertNodeOnly(node, newNode *DoublyLinkedListNode[E]) {
	newNode.list = d
	newNode.prev, newNode.next = node, node.next
	node.next, node.next.prev = newNode, newNode
}

func (d *DoublyLinkedList[E]) remove(node *DoublyLinkedListNode[E]) E {
	d.removeNodeOnly(node)
	node.list = nil
	d.length--
	return node.Value
}
func (d *DoublyLinkedList[E]) removeNodeOnly(node *DoublyLinkedListNode[E]) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (d *DoublyLinkedList[E]) checkNode(node *DoublyLinkedListNode[E]) {
	if node == nil {
		panic("nil node")
	}
	if node.list != d {
		panic("not belong to list")
	}
}
