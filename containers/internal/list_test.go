package internal

import (
	"testing"
)

func init() {
	blockCap = 2
}

func (l *List[E]) same(t [][]E) bool {
	node := l.head
	for _, elems := range t {
		if !node.same(elems) {
			return false
		}
		node = node.next
	}
	if node != nil {
		return false
	}
	return true
}

func TestList_PushBack(t *testing.T) {
	type Case struct {
		Val int
		Exp [][]int
	}
	l := NewList[int]()
	for i, c := range []*Case{
		{0, [][]int{{0}}},
		{1, [][]int{{0, 1}}},
		{2, [][]int{{0, 1}, {2}}},
	} {
		l.PushBack(c.Val)
		if !l.same(c.Exp) {
			t.Fatalf("case %d failed", i)
		}
	}
}

func TestList_PushFront(t *testing.T) {
	type Case struct {
		Val int
		Exp [][]int
	}
	l := NewList[int]()
	for i, c := range []*Case{
		{0, [][]int{{0}}},
		{1, [][]int{{1, 0}}},
		{2, [][]int{{2}, {1, 0}}},
	} {
		l.PushFront(c.Val)
		if !l.same(c.Exp) {
			t.Fatalf("case %d failed", i)
		}
	}
}

func TestList_Get(t *testing.T) {
	l := NewList[int]()
	for i := 0; i < 7; i++ {
		l.PushBack(i)
	}
	for i := 0; i < 7; i++ {
		if l.Get(i) != i {
			t.Fatalf("case %d failed", i)
		}
	}
	// out of range
	for i, c := range []int{-1, l.Len()} {
		func() {
			defer func() {
				r := recover()
				if r != "out of range" {
					t.Fatalf("case %d failed, real: %+v", i, r)
				}
			}()
			l.Get(c)
			t.Fatalf("case %d succeed", i)
		}()
	}
}

func TestList_Set(t *testing.T) {
	l := NewList[int]()
	for i := 0; i < 7; i++ {
		l.PushBack(i)
	}
	for i := 7; i > 0; i-- {
		l.PushBack(i)
		if r := l.Get(i); r != i {
			t.Fatalf("case %d failed, real: %+v", i, r)
		}
	}
	// out of range
	for i, c := range []int{-1, l.Len()} {
		func() {
			defer func() {
				r := recover()
				if r != "out of range" {
					t.Fatalf("case %d failed, real: %+v", i, r)
				}
			}()
			l.Set(c, 0)
			t.Fatalf("case %d succeed", i)
		}()
	}
}

func TestList_Len(t *testing.T) {
	type Case struct {
		Modify string
		Len    int
	}
	l := NewList[int]()
	for i, c := range []*Case{
		{"PushBack", 1},
		{"PushFront", 2},
		{"Insert", 3},
		{"Delete", 2},
	} {
		switch c.Modify {
		case "PushBack":
			l.PushBack(0)
		case "PushFront":
			l.PushFront(0)
		case "Insert":
			l.Insert(1, 0)
		case "Delete":
			l.Delete(1)
		default:
			t.Fatalf("unknow %s", c.Modify)
		}
		if r := l.Len(); r != c.Len {
			t.Fatalf("case %d failed, real: %+v", i, r)
		}
	}
}

func TestList_Delete(t *testing.T) {
	type Case struct {
		Idx int
		Exp [][]int
	}
	l := NewList[int]()
	for i := 0; i < 7; i++ {
		l.PushBack(i)
	}
	for i, c := range []*Case{
		{0, [][]int{{1}, {2, 3}, {4, 5}, {6}}},
		{2, [][]int{{1}, {2}, {4, 5}, {6}}},
		{2, [][]int{{1}, {2}, {5}, {6}}},
		{3, [][]int{{1}, {2}, {5}}},
		{1, [][]int{{1}, {5}}},
		{0, [][]int{{5}}},
		{0, [][]int{{}}},
	} {
		l.Delete(c.Idx)
		if !l.same(c.Exp) {
			t.Fatalf("case %d failed", i)
		}
	}
}

func TestList_Insert(t *testing.T) {
	type Case struct {
		Idx int
		Val int
		Exp [][]int
	}
	l := NewList[int]()
	for i, c := range []*Case{
		{0, 0, [][]int{{0}}},
		{0, -1, [][]int{{-1, 0}}},
		{0, -2, [][]int{{-2}, {-1, 0}}},
		{3, 2, [][]int{{-2}, {-1, 0}, {2}}},
		{3, 1, [][]int{{-2}, {-1, 0}, {1, 2}}},
		{2, 100, [][]int{{-2}, {-1, 100}, {0}, {1, 2}}},
		{1, 1, [][]int{{-2}, {1, -1}, {100}, {0}, {1, 2}}},
	} {
		l.Insert(c.Idx, c.Val)
		if !l.same(c.Exp) {
			t.Fatalf("case %d failed", i)
		}
	}
}

func TestList_Begin(t *testing.T) {
	l := NewList[int]()
	for i := 0; i < 7; i++ {
		l.PushBack(i)
	}
	iter := l.Begin()
	if iter.HasPrev() {
		t.Fatalf("has prev")
	}
	for i := 0; i < 7; i++ {
		if !iter.HasNext() {
			t.Fatalf("case %d failed, does not has next", i)
		}
		if r := iter.Next(); r != i {
			t.Fatalf("case %d failed, real: %d", i, r)
		}
	}
	if iter.HasNext() {
		t.Fatalf("not end")
	}
}

func TestList_End(t *testing.T) {
	l := NewList[int]()
	for i := 0; i < 7; i++ {
		l.PushBack(i)
	}
	iter := l.End()
	if iter.HasNext() {
		t.Fatalf("has next")
	}
	for i := 6; i >= 0; i-- {
		if !iter.HasPrev() {
			t.Fatalf("case %d failed, does not has prev", i)
		}
		if r := iter.Prev(); r != i {
			t.Fatalf("case %d failed, real: %d", i, r)
		}
	}
	if iter.HasPrev() {
		t.Fatalf("not end")
	}
}
