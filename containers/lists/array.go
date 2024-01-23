package lists

type Array[E any] struct {
	elems []E
}

func NewArray[E any](capacity int) *Array[E] {
	return &Array[E]{
		elems: make([]E, 0, capacity),
	}
}

func (a *Array[E]) Len() int {
	return len(a.elems)
}

func (a *Array[E]) Empty() bool {
	return a.Len() == 0
}

func (a *Array[E]) At(i int) E {
	return a.elems[i]
}

func (a *Array[E]) Front() E {
	return a.elems[0]
}

func (a *Array[E]) Back() E {
	return a.elems[a.Len()-1]
}

func (a *Array[E]) Set(i int, e E) {
	a.elems[i] = e
}

func (a *Array[E]) Insert(i int, e E) {
	a.elems = append(a.elems, e)
	copy(a.elems[i+1:], a.elems[i:])
	a.elems[i] = e
}

func (a *Array[E]) Remove(i int) E {
	e := a.elems[i]
	a.elems = append(a.elems[:i], a.elems[i+1:]...)
	return e
}

func (a *Array[E]) PushBack(e E) {
	a.Insert(a.Len(), e)
}

func (a *Array[E]) PopBack() E {
	return a.Remove(a.Len() - 1)
}

func (a *Array[E]) Clear() {
	a.elems = a.elems[:0]
}

func (a *Array[E]) Swap(i, j int) {
	a.elems[i], a.elems[j] = a.elems[j], a.elems[i]
}
