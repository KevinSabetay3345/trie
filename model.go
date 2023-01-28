package trie

type Trie struct {
	Root       *Node
	WordsCount int
}

type Node struct {
	Children map[rune]*Node
	Value    interface{}
}

type PathNode struct {
	Node *Node
	Rune rune
}
