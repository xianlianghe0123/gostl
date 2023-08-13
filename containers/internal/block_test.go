package internal

import (
	"reflect"
	"slices"
	"testing"
)

const BlockCap = 8

func initBlock() *Block[int] {
	b := NewBlock[int](BlockCap)
	b.Insert(0, 0)
	b.Insert(1, 1)
	return b
}

func (b *Block[E]) same(elems []E) bool {
	if b.Len() != len(elems) {
		return false
	}
	for i := 0; i < b.Len(); i++ {
		if !reflect.DeepEqual(b.elems[i], elems[i]) {
			return false
		}
	}
	return true
}

func TestBlock_Add(t *testing.T) {
	b := initBlock()
	type Case struct {
		Idx   int
		Value int
		Exp   []int
	}
	for i, c := range []*Case{
		{2, 2, []int{0, 1, 2}},
		{0, -1, []int{-1, 0, 1, 2}},
		{2, 100, []int{-1, 0, 100, 1, 2}},
	} {
		b.Insert(c.Idx, c.Value)
		if !slices.Equal(b.elems[:b.Len()], c.Exp) {
			t.Fatalf("case %d failed, real: %+v", i, b.elems)
		}
	}
	// full
	full := NewBlock[int](BlockCap)
	for i := 0; i < full.Cap(); i++ {
		full.Insert(i, i)
	}
	func() {
		defer func() {
			r := recover()
			if r != "block is full" {
				t.Fatalf("block is full")
			}
		}()
		full.Insert(1, 1)
		t.Fatalf("succeed")
	}()
}

func TestBlock_Get(t *testing.T) {
	b := initBlock()
	for i, c := range [][2]int{
		{0, 0},
		{1, 1},
	} {
		if r := b.Get(c[0]); r != c[1] {
			t.Fatalf("case %d failed, real: %+v", i, r)
		}
	}

	// out of range
	for i, c := range []int{-1, b.Len()} {
		func() {
			defer func() {
				r := recover()
				if r != "out of range" {
					t.Fatalf("case %d failed, real: %+v", i, r)
				}
			}()
			b.Get(c)
			t.Fatalf("case %d succeed", i)
		}()
	}
}

func TestBlock_Del(t *testing.T) {
	b := initBlock()
	for i, c := range [][][]int{
		{{1, 1}, {1}},
		{{0, 0}, {}},
	} {
		r := b.Delete(c[0][0])
		if r != c[0][1] && slices.Equal(b.elems, c[1]) {
			t.Fatalf("case %d failed, real: %+v", i, r)
		}
	}
	b.Insert(0, 2)

	// out of range
	for i, c := range []int{-1, b.Len()} {
		func() {
			defer func() {
				r := recover()
				if r != "out of range" {
					t.Fatalf("case %d failed, real: %+v", i, r)
				}
			}()
			b.Delete(c)
			t.Fatalf("case %d succeed", i)
		}()
	}
}

func TestBlock_Set(t *testing.T) {
	b := initBlock()
	for i, c := range [][][]int{
		{{1, 100}, {0, 100}},
		{{0, -100}, {-100, 100}},
	} {
		r := b.Delete(c[0][0])
		if r != c[0][1] && slices.Equal(b.elems, c[1]) {
			t.Fatalf("case %d failed, real: %+v", i, r)
		}
	}

	// out of range
	for i, c := range []int{-1, b.Len()} {
		func() {
			defer func() {
				r := recover()
				if r != "out of range" {
					t.Fatalf("case %d failed, real: %+v", i, r)
				}
			}()
			b.Set(c, 10)
			t.Fatalf("case %d succeed", i)
		}()
	}
}

func TestBlock_Full(t *testing.T) {
	b := NewBlock[int](BlockCap)
	for i := 0; i < b.Cap()-1; i++ {
		b.Insert(i, i)
		if b.Full() {
			t.Fatalf("full")
		}
	}
	b.Insert(0, 1)
	if !b.Full() {
		t.Fatalf("not full")
	}
}

func TestBlock_Len(t *testing.T) {
	b := NewBlock[int](BlockCap)
	for i := 0; i < b.Cap(); i++ {
		b.Insert(i, i)
		if b.Len() != i+1 {
			t.Fatalf("len err")
		}
	}
	for i := b.Cap() - 1; i >= 0; i-- {
		b.Delete(i)
		if b.Len() != i {
			t.Fatalf("len err")
		}
	}
}
