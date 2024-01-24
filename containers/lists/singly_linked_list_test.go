package lists

import (
	"slices"
	"testing"
)

func buildSinglyLinkedList(l int) (*SinglyLinkedList[int], []*SinglyLinkedListNode[int]) {
	s := NewSinglyLinkedList[int]()
	nodes := make([]*SinglyLinkedListNode[int], 0, l)
	for i := 1; i <= l; i++ {
		s.PushBack(i)
		nodes = append(nodes, s.Back())
	}
	return s, nodes
}

func getSinglyListedListElems(s *SinglyLinkedList[int]) []int {
	res := make([]int, 0, s.Len())
	for node := s.Front(); node != nil; node = node.Next() {
		res = append(res, node.Value)
	}
	return res
}

func TestSinglyLinkedList_Len(t *testing.T) {
	s, _ := buildSinglyLinkedList(3)
	cases := []struct {
		result int
	}{
		{2},
		{1},
		{0},
	}
	for i, c := range cases {
		s.PopBack()
		r := s.Len()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSinglyLinkedList_Empty(t *testing.T) {
	s, _ := buildSinglyLinkedList(1)
	cases := []struct {
		f      func()
		result bool
	}{
		{func() { s.PopBack() }, true},
		{func() { s.PushBack(3) }, false},
	}
	for i, c := range cases {
		c.f()
		r := s.Empty()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSinglyLinkedList_Front(t *testing.T) {
	s := SinglyLinkedList[int]{}
	cases := []struct {
		e      int
		result int
	}{
		{2, 2},
		{1, 1},
		{3, 3},
	}
	for i, c := range cases {
		s.PushFront(c.e)
		r := s.Front().Value
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSinglyLinkedList_Back(t *testing.T) {
	s := SinglyLinkedList[int]{}
	cases := []struct {
		e      int
		result int
	}{
		{2, 2},
		{1, 1},
		{3, 3},
	}
	for i, c := range cases {
		s.PushBack(c.e)
		r := s.Back().Value
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSinglyLinkedList_InsertBefore(t *testing.T) {
	s, nodes := buildSinglyLinkedList(3)
	cases := []struct {
		node   *SinglyLinkedListNode[int]
		e      int
		result int
		len    int
		elems  []int
	}{
		{nodes[0], 7, 7, 4, []int{7, 1, 2, 3}},
		{nodes[1], 8, 8, 5, []int{7, 1, 8, 2, 3}},
		{nodes[2], 9, 9, 6, []int{7, 1, 8, 2, 9, 3}},
	}
	for i, c := range cases {
		r := s.InsertBefore(c.node, c.e).Value
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_InsertAfter(t *testing.T) {
	s, nodes := buildSinglyLinkedList(3)
	cases := []struct {
		node   *SinglyLinkedListNode[int]
		e      int
		result int
		len    int
		elems  []int
	}{
		{nodes[0], 7, 7, 4, []int{1, 7, 2, 3}},
		{nodes[1], 8, 8, 5, []int{1, 7, 2, 8, 3}},
		{nodes[2], 9, 9, 6, []int{1, 7, 2, 8, 3, 9}},
	}
	for i, c := range cases {
		r := s.InsertAfter(c.node, c.e).Value
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	s, nodes := buildSinglyLinkedList(5)
	cases := []struct {
		node   *SinglyLinkedListNode[int]
		result int
		len    int
		elems  []int
	}{
		{nodes[0], 1, 4, []int{2, 3, 4, 5}},
		{nodes[2], 3, 3, []int{2, 4, 5}},
		{nodes[4], 5, 2, []int{2, 4}},
	}
	for i, c := range cases {
		r := s.Remove(c.node)
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_PushFront(t *testing.T) {
	s, _ := buildSinglyLinkedList(0)
	cases := []struct {
		e     int
		len   int
		elems []int
	}{
		{1, 1, []int{1}},
		{2, 2, []int{2, 1}},
		{3, 3, []int{3, 2, 1}},
	}
	for i, c := range cases {
		s.PushFront(c.e)
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_PushBack(t *testing.T) {
	s, _ := buildSinglyLinkedList(0)
	cases := []struct {
		e     int
		len   int
		elems []int
	}{
		{1, 1, []int{1}},
		{2, 2, []int{1, 2}},
		{3, 3, []int{1, 2, 3}},
	}
	for i, c := range cases {
		s.PushBack(c.e)
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_PopFront(t *testing.T) {
	s, _ := buildSinglyLinkedList(3)
	cases := []struct {
		result int
		len    int
		elems  []int
	}{
		{1, 2, []int{2, 3}},
		{2, 1, []int{3}},
		{3, 0, []int{}},
	}
	for i, c := range cases {
		r := s.PopFront()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_PopBack(t *testing.T) {
	s, _ := buildSinglyLinkedList(3)
	cases := []struct {
		result int
		len    int
		elems  []int
	}{
		{3, 2, []int{1, 2}},
		{2, 1, []int{1}},
		{1, 0, []int{}},
	}
	for i, c := range cases {
		r := s.PopBack()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_MoveFront(t *testing.T) {
	s, nodes := buildSinglyLinkedList(3)
	cases := []struct {
		node  *SinglyLinkedListNode[int]
		len   int
		elems []int
	}{
		{nodes[0], 3, []int{1, 2, 3}},
		{nodes[1], 3, []int{2, 1, 3}},
		{nodes[2], 3, []int{3, 2, 1}},
	}
	for i, c := range cases {
		s.MoveFront(c.node)
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestSinglyLinkedList_MoveBack(t *testing.T) {
	s, nodes := buildSinglyLinkedList(3)
	cases := []struct {
		node  *SinglyLinkedListNode[int]
		len   int
		elems []int
	}{
		{nodes[2], 3, []int{1, 2, 3}},
		{nodes[1], 3, []int{1, 3, 2}},
		{nodes[0], 3, []int{3, 2, 1}},
	}
	for i, c := range cases {
		s.MoveBack(c.node)
		if s.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, s.Len())
		}
		elems := getSinglyListedListElems(s)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}
