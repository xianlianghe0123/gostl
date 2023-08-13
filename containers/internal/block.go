package internal

import (
	"container/list"
)

type Block[E any] struct {
	elems []E
	len   int
	cap   int
}

func NewBlock[E any](cap int) *Block[E] {
	return &Block[E]{
		elems: make([]E, cap),
		len:   0,
		cap:   cap,
	}
}

func (b *Block[E]) Len() int {
	return b.len
}

func (b *Block[E]) Cap() int {
	return b.cap
}

func (b *Block[E]) Full() bool {
	return b.len == b.cap
}

func (b *Block[E]) Get(idx int) E {
	b.rangeCheck(idx)
	return b.elems[idx]
}

func (b *Block[E]) Set(idx int, elem E) {
	b.rangeCheck(idx)
	b.elems[idx] = elem
}

func (b *Block[E]) Delete(idx int) E {
	list.New()
	b.rangeCheck(idx)
	r := b.elems[idx]
	copy(b.elems[idx:], b.elems[idx+1:b.len])
	b.len--
	return r
}

func (b *Block[E]) Insert(idx int, elem E) {
	if b.Full() {
		panic("block is full")
	}
	if idx < 0 || idx > b.len {
		panic("out of range")
	}
	copy(b.elems[idx+1:], b.elems[idx:b.len])
	b.elems[idx] = elem
	b.len++
}

func (b *Block[E]) Begin() *BlockIterator[E] {
	return NewBlockIterator[E](b, -1)
}

func (b *Block[E]) End() *BlockIterator[E] {
	return NewBlockIterator[E](b, b.Len())
}

func (b *Block[E]) rangeCheck(idx int) {
	if idx < 0 || idx >= b.len {
		panic("out of range")
	}
}

type BlockIterator[E any] struct {
	block *Block[E]
	idx   int
}

func NewBlockIterator[E any](block *Block[E], idx int) *BlockIterator[E] {
	return &BlockIterator[E]{
		block: block,
		idx:   idx,
	}
}

func (b *BlockIterator[E]) HasNext() bool {
	if b.idx+1 < b.block.Len() {
		return true
	}
	return false
}

func (b *BlockIterator[E]) Next() E {
	if !b.HasNext() {
		var e E
		return e
	}
	b.idx++
	e := b.block.elems[b.idx]
	return e
}

func (b *BlockIterator[E]) HasPrev() bool {
	if b.idx-1 >= 0 {
		return true
	}
	return false
}

func (b *BlockIterator[E]) Prev() E {
	if !b.HasPrev() {
		var e E
		return e
	}
	b.idx--
	e := b.block.elems[b.idx]
	return e
}
