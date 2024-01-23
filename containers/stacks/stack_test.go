package stacks

import (
	"slices"
	"testing"
)

func buildStack(l int) *Stack[int] {
	a := NewStack[int](l)
	for i := 1; i <= l; i++ {
		a.Push(i)
	}
	return a
}

func TestStack_Size(t *testing.T) {
	s := buildStack(3)
	cases := []struct {
		result int
	}{
		{2},
		{1},
		{0},
	}
	for i, c := range cases {
		s.Pop()
		r := s.Size()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestStack_Empty(t *testing.T) {
	s := buildStack(0)
	cases := []struct {
		f      func()
		result bool
	}{
		{func() { s.Push(1) }, false},
		{func() { s.Pop() }, true},
	}
	for i, c := range cases {
		c.f()
		r := s.Empty()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestStack_Top(t *testing.T) {
	s := buildStack(0)
	cases := []struct {
		result int
	}{
		{1},
		{4},
		{6},
		{3},
	}
	for i, c := range cases {
		s.Push(c.result)
		r := s.Top()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestStack_Push(t *testing.T) {
	s := buildStack(0)
	cases := []struct {
		e      int
		result []int
	}{
		{1, []int{1}},
		{4, []int{1, 4}},
		{6, []int{1, 4, 6}},
		{3, []int{1, 4, 6, 3}},
	}
	for i, c := range cases {
		s.Push(c.e)
		if !slices.Equal(s.elems, c.result) {
			t.Fatalf("case %d, real:%v", i, s.elems)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	s := buildStack(3)
	cases := []struct {
		result int
		elems  []int
	}{
		{3, []int{1, 2}},
		{2, []int{1}},
		{1, []int{}},
	}
	for i, c := range cases {
		r := s.Pop()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if !slices.Equal(s.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, s.elems)
		}
	}
}
