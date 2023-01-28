package trie

import (
	"errors"
)

// Insert returns true when is a new element,
// false if it replaces an existing value.
func (trie *Trie) Insert(key string, value interface{}) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("key must not be empty")
	}
	if value == nil {
		return false, errors.New("value must not be empty")
	}

	if trie.Root == nil {
		trie.Root = &Node{Children: make(map[rune]*Node)}
	}

	node := trie.Root
	for _, v := range key {
		if node.Children[v] == nil {
			node.Children[v] = &Node{Children: make(map[rune]*Node)}
		}
		node = node.Children[v]
	}

	isNewVal := node.Value == nil
	node.Value = value

	if isNewVal {
		trie.WordsCount++
	}

	return isNewVal, nil
}

func (trie *Trie) Find(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("key must not be empty")
	}
	if trie.Root == nil {
		return nil, errors.New("trie is empty")
	}

	nodo := trie.Root
	for _, v := range key {
		nodo = nodo.Children[v]
		if nodo == nil {
			return nil, errors.New("element not found")
		}
	}

	if nodo.Value == nil {
		return nil, errors.New("element not found")
	}
	return nodo.Value, nil
}

func (trie *Trie) Delete(key string) error {
	if len(key) == 0 {
		return errors.New("key must not be empty")
	}
	if trie.Root == nil {
		return errors.New("trie is empty")
	}

	nodo := trie.Root
	path := make([]PathNode, len(key))
	for i, v := range key {
		path[i] = PathNode{Rune: v, Node: nodo}
		nodo = nodo.Children[v]
		if nodo == nil {
			return errors.New("element not found")
		}
	}

	if nodo.Value == nil {
		return errors.New("element not found")
	}
	nodo.Value = nil

	if len(nodo.Children) == 0 {
		// iterate backwards over path
		for i := len(key) - 1; i >= 0; i-- {
			parent := path[i].Node
			r := path[i].Rune
			delete(parent.Children, r)
			if len(parent.Children) > 0 {
				// parent has other children, stop
				break
			}
			if parent.Value != nil {
				// parent has a value, stop
				break
			}
		}

		if len(trie.Root.Children) == 0 {
			trie.Root = nil
		}

	}

	trie.WordsCount--
	return nil
}
