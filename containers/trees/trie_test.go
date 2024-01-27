package trees

import "testing"

func TestTrie_ContainsCount(t *testing.T) {
	var trie Trie[byte, []byte]
	for _, e := range []string{"aoe", "aoe", "iuv", "bpmf", "a", "a", "1", "22", "22", "333", "333", "333"} {
		trie.Add([]byte(e))
	}

	cases := []struct {
		e      string
		result int
	}{
		{"aoe", 2},
		{"333", 3},
		{"", 0},
		{"216", 0},
	}
	for i, c := range cases {

		r := trie.ContainsCount([]byte(c.e))
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestTrie_HasPrefix(t *testing.T) {
	cases := []struct {
		elems  []string
		prefix string
		result bool
	}{
		{[]string{"aoe", "aoe", "iuv", "bmpf", "a", "a"}, "a", true},
		{[]string{"1", "12", "123"}, "", true},
		{[]string{"1", "22", "22", "333", "333", "333"}, "4", false},
	}
	for i, c := range cases {
		var trie Trie[byte, []byte]
		for _, e := range c.elems {
			trie.Add([]byte(e))
		}
		r := trie.HasPrefix([]byte(c.prefix))
		if r != c.result {
			t.Fatalf("case %d, real:%v", i, r)
		}
	}
}

func TestTrie_Add(t *testing.T) {
	trie := NewTrie[byte, []byte]()
	cases := []struct {
		e   string
		cnt int
	}{
		{"abcd", 1},
		{"lala", 2},
		{"aaaaaa", 3},
	}
	for i, c := range cases {
		trie.Add([]byte(c.e))
		r := trie.cnt
		if r != c.cnt {
			t.Fatalf("case %d, real:%v", i, r)
		}
		cnt := trie.ContainsCount([]byte(c.e))
		if cnt != 1 {
			t.Fatalf("case %d, real:%v", i, cnt)
		}
	}
}
