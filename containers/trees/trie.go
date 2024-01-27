package trees

type trieNode[E comparable] struct {
	cnt      int
	children map[E]*trieNode[E]
}

type Trie[E comparable, S ~[]E] struct {
	root *trieNode[E]
	cnt  int
}

func NewTrie[E comparable, S ~[]E]() *Trie[E, S] {
	t := new(Trie[E, S])
	t.init()
	return t
}

func (t *Trie[E, S]) ContainsCount(s S) int {
	node := t.searchNode(s)
	if node == nil {
		return 0
	}
	return node.cnt
}

func (t *Trie[E, S]) Add(s S) {
	t.init()
	node := t.root
	for _, e := range s {
		if node.children[e] == nil {
			node.children[e] = &trieNode[E]{children: make(map[E]*trieNode[E])}
		}
		node = node.children[e]
	}
	node.cnt++
	t.cnt++
}

func (t *Trie[E, S]) HasPrefix(prefix S) bool {
	return t.searchNode(prefix) != nil
}

func (t *Trie[E, S]) searchNode(s S) *trieNode[E] {
	node := t.root
	for _, e := range s {
		if node.children[e] == nil {
			return nil
		}
		node = node.children[e]
	}
	return node
}

func (t *Trie[E, S]) init() {
	if t.root != nil {
		return
	}
	t.root = &trieNode[E]{children: make(map[E]*trieNode[E])}
}
