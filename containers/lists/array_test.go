package lists

import (
	"slices"
	"testing"
)

func buildArray(l int) *Array[int] {
	a := NewArray[int](l)
	for i := 1; i <= l; i++ {
		a.PushBack(i)
	}
	return a
}

func TestArray_Len(t *testing.T) {
	a := buildArray(5)
	cases := []struct {
		result int
	}{
		{4},
		{3},
		{2},
		{1},
		{0},
	}
	for i, c := range cases {
		a.PopBack()
		r := a.Len()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_Empty(t *testing.T) {
	a := buildArray(0)
	cases := []struct {
		f      func()
		result bool
	}{
		{func() { a.PushBack(1) }, false},
		{func() { a.PopBack() }, true},
	}
	for i, c := range cases {
		c.f()
		r := a.Empty()
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_At(t *testing.T) {
	a := buildArray(5)
	cases := []struct {
		i, e int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	for i, c := range cases {
		r := a.At(c.i)
		if r != c.e {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_Front(t *testing.T) {
	a := buildArray(0)
	cases := []struct {
		e int
	}{
		{2},
		{3},
		{5},
	}
	for i, c := range cases {
		a.Insert(0, c.e)
		r := a.Front()
		if r != c.e {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_Back(t *testing.T) {
	a := buildArray(0)
	cases := []struct {
		e int
	}{
		{2},
		{3},
		{5},
	}
	for i, c := range cases {
		a.PushBack(c.e)
		r := a.Back()
		if r != c.e {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_Set(t *testing.T) {
	a := buildArray(5)
	cases := []struct {
		i, e int
	}{
		{2, 22},
		{3, 100},
		{3, 4},
	}
	for i, c := range cases {
		a.Set(c.i, c.e)
		r := a.At(c.i)
		if r != c.e {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_Insert(t *testing.T) {
	a := buildArray(5)
	cases := []struct {
		i, e  int
		elems []int
	}{
		{0, 0, []int{0, 1, 2, 3, 4, 5}},
		{3, 100, []int{0, 1, 2, 100, 3, 4, 5}},
		{6, 6, []int{0, 1, 2, 100, 3, 4, 6, 5}},
		{8, 333, []int{0, 1, 2, 100, 3, 4, 6, 5, 333}},
	}
	for i, c := range cases {
		a.Insert(c.i, c.e)
		if !slices.Equal(a.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, a.elems)
		}
	}
}

func TestArray_Delete(t *testing.T) {
	a := buildArray(5)
	cases := []struct {
		i, e  int
		elems []int
	}{
		{0, 1, []int{2, 3, 4, 5}},
		{3, 5, []int{2, 3, 4}},
		{1, 3, []int{2, 4}},
	}
	for i, c := range cases {
		r := a.Remove(c.i)
		if r != c.e {
			t.Fatalf("case %d, real return:%v", i, r)
		}
		if !slices.Equal(a.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, a.elems)
		}
	}
}

func TestArray_PushBack(t *testing.T) {
	a := buildArray(0)
	cases := []struct {
		e     int
		elems []int
	}{
		{3, []int{3}},
		{2, []int{3, 2}},
		{1, []int{3, 2, 1}},
		{5, []int{3, 2, 1, 5}},
		{4, []int{3, 2, 1, 5, 4}},
	}
	for i, c := range cases {
		a.PushBack(c.e)
		r := a.Back()
		if r != c.e {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if !slices.Equal(a.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, a.elems)
		}
	}
}

func TestArray_PopBack(t *testing.T) {
	a := buildArray(4)
	cases := []struct {
		e     int
		elems []int
	}{
		{4, []int{1, 2, 3}},
		{3, []int{1, 2}},
		{2, []int{1}},
		{1, []int{}},
	}
	for i, c := range cases {
		r := a.PopBack()
		if r != c.e {
			t.Fatalf("case %d, real:%v", i, r)
		}
		if !slices.Equal(a.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, a.elems)
		}
	}
}

func TestArray_Clear(t *testing.T) {
	cases := []struct {
		a *Array[int]
	}{
		{buildArray(0)},
		{buildArray(3)},
	}
	for i, c := range cases {
		a := c.a
		a.Clear()
		r := a.Len()
		if r != 0 {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestArray_Swap(t *testing.T) {
	a := buildArray(5)
	cases := []struct {
		i, j  int
		elems []int
	}{
		{0, 1, []int{2, 1, 3, 4, 5}},
		{3, 1, []int{2, 4, 3, 1, 5}},
		{1, 3, []int{2, 1, 3, 4, 5}},
	}
	for i, c := range cases {
		a.Swap(c.i, c.j)
		if !slices.Equal(a.elems, c.elems) {
			t.Fatalf("case %d, real:%v", i, a.elems)
		}
	}
}
