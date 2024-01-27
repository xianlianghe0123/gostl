package sets

import (
	"maps"
	"slices"
	"testing"
)

func TestSet_Size(t *testing.T) {
	var s Set[int]
	s.Add(1, 2, 3)
	cases := []struct {
		e      int
		result int
	}{
		{1, 2},
		{2, 1},
		{3, 0},
	}
	for i, c := range cases {
		s.Remove(c.e)
		r := s.Size()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSet_Empty(t *testing.T) {
	var s Set[int]
	cases := []struct {
		f      func()
		result bool
	}{
		{func() { s.Add(1) }, false},
		{func() { s.Remove(1) }, true},
	}
	for i, c := range cases {
		c.f()
		r := s.Empty()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSet_Contains(t *testing.T) {
	s := NewSet[int]()
	s.Add(7, 9)
	cases := []struct {
		e      int
		result bool
	}{
		{1, false},
		{2, false},
		{7, true},
		{9, true},
	}
	for i, c := range cases {
		r := s.Contains(c.e)
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSet_Elems(t *testing.T) {
	s := NewSet[int]()
	s.Add(7, 9)
	cases := []struct {
		e      int
		result []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{7, []int{1, 2, 7}},
		{9, []int{1, 2, 7, 9}},
	}
	for i, c := range cases {
		s.Add(c.e)
		r := s.Elems()
		if slices.Equal(r, c.result) {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestSet_Add(t *testing.T) {
	s := NewSet[int]()
	cases := []struct {
		e      int
		size   int
		result map[int]struct{}
	}{
		{1, 1, map[int]struct{}{1: {}}},
		{2, 2, map[int]struct{}{1: {}, 2: {}}},
		{7, 3, map[int]struct{}{1: {}, 2: {}, 7: {}}},
		{9, 4, map[int]struct{}{1: {}, 2: {}, 7: {}, 9: {}}},
	}
	for i, c := range cases {
		s.Add(c.e)
		r := s.Size()
		if r != c.size {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if !maps.Equal(s.elems, c.result) {
			t.Fatalf("case %d, real:%v", i, s.elems)
		}
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet[int]()
	s.Add(3, 4, 5)
	cases := []struct {
		e      int
		size   int
		result map[int]struct{}
	}{
		{4, 2, map[int]struct{}{3: {}, 5: {}}},
		{9, 2, map[int]struct{}{3: {}, 5: {}}},
		{3, 1, map[int]struct{}{5: {}}},
		{5, 0, nil},
	}
	for i, c := range cases {
		s.Remove(c.e)
		r := s.Size()
		if r != c.size {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if !maps.Equal(s.elems, c.result) {
			t.Fatalf("case %d, real:%v", i, s.elems)
		}
	}
}

func TestSet_Clear(t *testing.T) {
	cases := []struct {
		elems []int
	}{
		{[]int{1, 2, 3, 4}},
		{[]int{5, 6, 7, 8, 9, 10}},
		{[]int{123, 35, 235, 25, 34}},
	}
	for i, c := range cases {
		s := NewSet[int]()
		s.Add(c.elems...)
		s.Clear()
		r := s.Size()
		if r != 0 {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if !maps.Equal(s.elems, (map[int]struct{})(nil)) {
			t.Fatalf("case %d, real:%v", i, s.elems)
		}
	}
}
