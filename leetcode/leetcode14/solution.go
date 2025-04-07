package leetcode14

type node struct {
	children map[int32]*node
	// end of word
	eow bool
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{root: &node{children: make(map[int32]*node)}}
}

func (t *trie) insert(word string) {
	cur := t.root

	for _, char := range word {
		if cur.children[char] == nil {
			cur.children[char] = &node{children: make(map[int32]*node)}
		}

		cur = cur.children[char]
	}

	cur.eow = true
}

func find(n *node, prefix string) string {
	if len(n.children) > 1 || n.eow {
		return prefix
	}

	for char, children := range n.children {
		if children != nil {
			return find(children, prefix+string(char))
		}
	}

	return ""
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	trie := newTrie()

	for _, str := range strs {
		trie.insert(str)
	}

	return find(trie.root, "")
}
