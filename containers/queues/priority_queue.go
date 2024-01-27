package queues

type PriorityQueue[E any] struct {
	elems          []E
	priorityHigher func(E, E) bool
}

func NewPriorityQueue[E any](capability int, priorityHigher func(E, E) bool) *PriorityQueue[E] {
	if priorityHigher == nil {
		panic("")
	}
	return &PriorityQueue[E]{
		elems:          make([]E, 0, capability),
		priorityHigher: priorityHigher,
	}
}

func (p *PriorityQueue[E]) Len() int {
	return len(p.elems)
}

func (p *PriorityQueue[E]) Empty() bool {
	return p.Len() == 0
}

func (p *PriorityQueue[E]) Push(e E) {
	p.elems = append(p.elems, e)
	p.up(p.Len() - 1)
}

func (p *PriorityQueue[E]) Pop() E {
	e := p.elems[0]
	p.swap(0, p.Len()-1)
	p.elems = p.elems[:p.Len()-1]
	p.down(0)
	return e
}

func (p *PriorityQueue[E]) Build(elems []E) {
	p.elems = elems
	p.init()
}

func (p *PriorityQueue[E]) init() {
	for i := p.Len()/2 - 1; i >= 0; i-- {
		p.down(i)
	}
}

func (p *PriorityQueue[E]) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if p.priorityHigher(p.elems[parent], p.elems[i]) {
			break
		}
		p.swap(i, parent)
		i = parent
	}
}

func (p *PriorityQueue[E]) down(i int) {
	for {
		left, right := 2*i+1, 2*i+2
		if left >= p.Len() {
			break
		}
		t := left
		if right < p.Len() && p.priorityHigher(p.elems[right], p.elems[left]) {
			t = right
		}
		if p.priorityHigher(p.elems[i], p.elems[t]) {
			break
		}
		p.swap(i, t)
		i = t
	}
}

func (p *PriorityQueue[E]) swap(i, j int) {
	p.elems[i], p.elems[j] = p.elems[j], p.elems[i]
}
