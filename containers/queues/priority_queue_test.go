package queues

import (
	"slices"
	"testing"
)

func buildPriorityQueue(elems []int) *PriorityQueue[int] {
	pq := NewPriorityQueue[int](10, func(i int, j int) bool { return i < j })
	for _, e := range elems {
		pq.Push(e)
	}
	return pq
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := buildPriorityQueue([]int{1, 2, 3})
	cases := []struct {
		result int
	}{
		{2},
		{1},
		{0},
	}
	for i, c := range cases {
		pq.Pop()
		r := pq.Len()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestPriorityQueue_Empty(t *testing.T) {
	pq := buildPriorityQueue([]int{1})
	cases := []struct {
		f      func()
		result bool
	}{
		{func() { pq.Pop() }, true},
		{func() { pq.Push(3) }, false},
	}
	for i, c := range cases {
		c.f()
		r := pq.Empty()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	pq := buildPriorityQueue([]int{})
	cases := []struct {
		e     int
		len   int
		elems []int
	}{
		{99, 1, []int{99}},
		{55, 2, []int{55, 99}},
		{23, 3, []int{23, 99, 55}},
		{12, 4, []int{12, 23, 55, 99}},
		{34, 5, []int{12, 23, 55, 99, 34}},
		{77, 6, []int{12, 23, 55, 99, 34, 77}},
		{64, 7, []int{12, 23, 55, 99, 34, 77, 64}},
		{75, 8, []int{12, 23, 55, 75, 34, 77, 64, 99}},
		{91, 9, []int{12, 23, 55, 75, 34, 77, 64, 99, 91}},
		{22, 10, []int{12, 22, 55, 75, 23, 77, 64, 99, 91, 34}},
	}
	for i, c := range cases {
		pq.Push(c.e)
		if pq.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, pq.Len())
		}
		if !slices.Equal(pq.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, pq.elems)
		}
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := buildPriorityQueue([]int{12, 22, 55, 75, 23, 77, 64, 99, 91, 34})
	cases := []struct {
		result int
		len    int
		elems  []int
	}{
		{12, 9, []int{22, 23, 55, 75, 34, 77, 64, 99, 91}},
		{22, 8, []int{23, 34, 55, 75, 91, 77, 64, 99}},
		{23, 7, []int{34, 75, 55, 99, 91, 77, 64}},
		{34, 6, []int{55, 75, 64, 99, 91, 77}},
		{55, 5, []int{64, 75, 77, 99, 91}},
		{64, 4, []int{75, 91, 77, 99}},
		{75, 3, []int{77, 91, 99}},
		{77, 2, []int{91, 99}},
		{91, 1, []int{99}},
		{99, 0, []int{}},
	}
	for i, c := range cases {
		r := pq.Pop()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if pq.Len() != c.len {
			t.Fatalf("case %d, len:%v", i, pq.Len())
		}
		if !slices.Equal(pq.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, pq.elems)
		}
	}
}

func TestPriorityQueue_Build(t *testing.T) {
	pq := buildPriorityQueue(nil)
	cases := []struct {
		elems  []int
		result []int
	}{
		{[]int{11, 89, 48, 17, 86, 85, 78, 9, 12, 97}, []int{9, 11, 48, 12, 86, 85, 78, 17, 89, 97}},
		{[]int{24, 99, 57, 70, 72, 80, 32, 10, 94, 19}, []int{10, 19, 32, 70, 24, 80, 57, 99, 94, 72}},
		{[]int{66, 83, 76, 88, 79, 99, 61, 40, 73, 69}, []int{40, 66, 61, 73, 69, 99, 76, 88, 83, 79}},
		{[]int{61, 28, 66, 30, 97, 90, 15, 27, 76, 33}, []int{15, 27, 61, 28, 33, 90, 66, 30, 76, 97}},
		{[]int{29, 42, 56, 57, 91, 32, 47, 14, 44, 72}, []int{14, 29, 32, 42, 72, 56, 47, 57, 44, 91}},
	}
	for i, c := range cases {
		pq.Build(c.elems)
		if !slices.Equal(pq.elems, c.result) {
			t.Fatalf("case %d, real:%v", i, pq.elems)
		}
	}
}
