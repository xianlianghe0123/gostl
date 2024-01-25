package queues

import (
	"slices"
	"testing"
)

func buildQueue(l int) *Queue[int] {
	q := NewQueue[int]()
	for i := 1; i <= l; i++ {
		q.Push(i)
	}
	return q
}

func getQueueElems(q *Queue[int]) []int {
	res := make([]int, 0, q.Len())
	for node := q.head.next; node != nil; node = node.next {
		res = append(res, node.value)
	}
	return res
}

func TestQueue_Len(t *testing.T) {
	q := buildQueue(3)
	cases := []struct {
		result int
	}{
		{2},
		{1},
		{0},
	}
	for i, c := range cases {
		q.Pop()
		r := q.Len()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestQueue_Empty(t *testing.T) {
	s := buildQueue(1)
	cases := []struct {
		f      func()
		result bool
	}{
		{func() { s.Pop() }, true},
		{func() { s.Push(3) }, false},
	}
	for i, c := range cases {
		c.f()
		r := s.Empty()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestQueue_Front(t *testing.T) {
	q := buildQueue(5)
	cases := []struct {
		result int
	}{
		{2},
		{3},
		{4},
	}
	for i, c := range cases {
		q.Pop()
		r := q.Front()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestQueue_Back(t *testing.T) {
	q := Queue[int]{}
	cases := []struct {
		e      int
		result int
	}{
		{2, 2},
		{1, 1},
		{3, 3},
	}
	for i, c := range cases {
		q.Push(c.e)
		r := q.Back()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestQueue_Push(t *testing.T) {
	q := buildQueue(0)
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
		q.Push(c.e)
		if q.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, q.Len())
		}
		elems := getQueueElems(q)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestQueue_Pop(t *testing.T) {
	q := buildQueue(3)
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
		r := q.Pop()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if q.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, q.Len())
		}
		elems := getQueueElems(q)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}
