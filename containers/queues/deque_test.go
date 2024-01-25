package queues

import (
	"slices"
	"testing"
)

func buildDeque(l int) *Deque[int] {
	q := NewDeque[int]()
	for i := 1; i <= l; i++ {
		q.PushBack(i)
	}
	return q
}

func getDequeElems(q *Deque[int]) []int {
	res := make([]int, 0, q.Len())
	for node := q.empty.next; node != q.empty; node = node.next {
		res = append(res, node.value)
	}
	return res
}

func TestDeque_Len(t *testing.T) {
	d := buildDeque(3)
	cases := []struct {
		result int
	}{
		{2},
		{1},
		{0},
	}
	for i, c := range cases {
		d.PopBack()
		r := d.Len()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestDeque_Empty(t *testing.T) {
	s := buildDeque(1)
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

func TestDeque_Front(t *testing.T) {
	q := buildDeque(5)
	cases := []struct {
		result int
	}{
		{2},
		{3},
		{4},
	}
	for i, c := range cases {
		q.PopFront()
		r := q.Front()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestDeque_Back(t *testing.T) {
	q := Deque[int]{}
	cases := []struct {
		e      int
		result int
	}{
		{2, 2},
		{1, 1},
		{3, 3},
	}
	for i, c := range cases {
		q.PushBack(c.e)
		r := q.Back()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestDeque_PushFront(t *testing.T) {
	q := buildDeque(0)
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
		q.PushFront(c.e)
		if q.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, q.Len())
		}
		elems := getDequeElems(q)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestDeque_PushBack(t *testing.T) {
	q := buildDeque(0)
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
		q.PushBack(c.e)
		if q.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, q.Len())
		}
		elems := getDequeElems(q)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestDeque_PopFront(t *testing.T) {
	q := buildDeque(3)
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
		r := q.PopFront()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if q.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, q.Len())
		}
		elems := getDequeElems(q)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}

func TestDeque_PopBack(t *testing.T) {
	q := buildDeque(3)
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
		r := q.PopBack()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if q.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, q.Len())
		}
		elems := getDequeElems(q)
		if !slices.Equal(elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, elems)
		}
	}
}
