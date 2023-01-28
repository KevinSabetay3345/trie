package trie

import (
	"testing"
)

func TestEmptyTrie(t *testing.T) {
	var trie Trie

	checkNoElems(trie, t)

	val, err := trie.Find("test")
	if val != nil {
		t.Errorf("Expected %v, Got %v", nil, val)
	}
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "trie is empty", nil)
	}
	if err.Error() != "trie is empty" {
		t.Errorf("Expected %s, Got %s", "trie is empty", err.Error())
	}

	err = trie.Delete("test")
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "trie is empty", err)
	}
	if err.Error() != "trie is empty" {
		t.Errorf("Expected %s, Got %s", "trie is empty", err.Error())
	}
	err = trie.Delete("")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "key must not be empty", err)
	}
}

func TestInsertTrie(t *testing.T) {
	var trie Trie

	_, err := trie.Insert("", nil)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "key must not be empty", err)
	}
	_, err = trie.Insert("test", nil)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "value must not be empty", err)
	}

	exists, err := trie.Insert("test", 4)
	if !exists {
		t.Errorf("Expected %t, Got %t", true, exists)
	}
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkElem(trie, 1, "test", 4, t)

	exists, err = trie.Insert("testt", 2)
	if !exists {
		t.Errorf("Expected %t, Got %t", true, exists)
	}
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkElem(trie, 2, "test", 4, t)
	checkElem(trie, 2, "testt", 2, t)

	exists, err = trie.Insert("test", 40)
	if exists {
		t.Errorf("Expected %t, Got %t", false, exists)
	}
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkElem(trie, 2, "test", 40, t)

	exists, err = trie.Insert("te", false)
	if !exists {
		t.Errorf("Expected %t, Got %t", true, exists)
	}
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkElem(trie, 3, "te", false, t)

}

func TestDeleteTrie(t *testing.T) {
	var prefixTrie Trie

	exists, err := prefixTrie.Insert("test", 4)
	if !exists {
		t.Errorf("Expected %t, Got %t", true, exists)
	}
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	err = prefixTrie.Delete("test")
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkNoElems(prefixTrie, t)

	err = prefixTrie.Delete("testes")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}

	prefixTrie.Insert("te", 1)
	prefixTrie.Insert("tes", 2)
	prefixTrie.Insert("test", 3)
	prefixTrie.Insert("tesa", 4)
	prefixTrie.Insert("testtest", 5)
	checkElem(prefixTrie, 5, "te", 1, t)
	checkElem(prefixTrie, 5, "tes", 2, t)
	checkElem(prefixTrie, 5, "test", 3, t)
	checkElem(prefixTrie, 5, "tesa", 4, t)
	checkElem(prefixTrie, 5, "testtest", 5, t)

	err = prefixTrie.Delete("testtesttest")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}
	err = prefixTrie.Delete("testte")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}

	err = prefixTrie.Delete("testtest")
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkElem(prefixTrie, 4, "te", 1, t)
	err = prefixTrie.Delete("te")
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkElem(prefixTrie, 3, "test", 3, t)

	var noPrefixTrie Trie
	noPrefixTrie.Insert("test", 1)
	noPrefixTrie.Insert("mate", 2)
	noPrefixTrie.Insert("cafe", 3)
	noPrefixTrie.Insert("chocolatada", 4)

	if err = noPrefixTrie.Delete("cafe"); err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if err = noPrefixTrie.Delete("test"); err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkElem(noPrefixTrie, 2, "mate", 2, t)
	checkElem(noPrefixTrie, 2, "chocolatada", 4, t)

	if err = noPrefixTrie.Delete("mate"); err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if err = noPrefixTrie.Delete("chocolatada"); err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkNoElems(noPrefixTrie, t)
}

func TestFindTrie(t *testing.T) {
	var prefixTrie Trie

	_, err := prefixTrie.Find("")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "key must not be empty", err)
	}
	_, err = prefixTrie.Find("test")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}

	prefixTrie.Insert("a", 1)
	prefixTrie.Insert("ab", 2)
	prefixTrie.Insert("abc", 3)
	prefixTrie.Insert("abcde", 4)

	_, err = prefixTrie.Find("abcdef")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}
	_, err = prefixTrie.Find("abcd")
	if err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}

	val, err := prefixTrie.Find("ab")
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if val != 2 {
		t.Errorf("Expected %d, Got %v", 2, val)
	}

	if val, _ = prefixTrie.Find("a"); val != 1 {
		t.Errorf("Expected %d, Got %v", 1, val)
	}
	if val, _ = prefixTrie.Find("abc"); val != 3 {
		t.Errorf("Expected %d, Got %v", 3, val)
	}
	if val, _ = prefixTrie.Find("abcde"); val != 4 {
		t.Errorf("Expected %d, Got %v", 4, val)
	}

	prefixTrie.Insert("a", 5)
	prefixTrie.Insert("ab", 6)
	prefixTrie.Insert("abc", 7)
	prefixTrie.Insert("abcde", 8)

	if val, _ = prefixTrie.Find("a"); val != 5 {
		t.Errorf("Expected %d, Got %v", 5, val)
	}
	if val, _ = prefixTrie.Find("ab"); val != 6 {
		t.Errorf("Expected %d, Got %v", 6, val)
	}
	if val, _ = prefixTrie.Find("abc"); val != 7 {
		t.Errorf("Expected %d, Got %v", 7, val)
	}
	if val, _ = prefixTrie.Find("abcde"); val != 8 {
		t.Errorf("Expected %d, Got %v", 8, val)
	}

	if err = prefixTrie.Delete("abc"); err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if _, err = prefixTrie.Find("abc"); err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}
	if val, _ = prefixTrie.Find("ab"); val != 6 {
		t.Errorf("Expected %d, Got %v", 6, val)
	}
	if val, _ = prefixTrie.Find("abcde"); val != 8 {
		t.Errorf("Expected %d, Got %v", 8, val)
	}

	var noPrefixTrie Trie
	noPrefixTrie.Insert("abc", 1)
	noPrefixTrie.Insert("def", 2)
	noPrefixTrie.Insert("gh", 3)
	noPrefixTrie.Insert("ijklmn", 4)

	if val, _ = noPrefixTrie.Find("abc"); val != 1 {
		t.Errorf("Expected %d, Got %v", 1, val)
	}
	if val, _ = noPrefixTrie.Find("def"); val != 2 {
		t.Errorf("Expected %d, Got %v", 2, val)
	}
	if val, _ = noPrefixTrie.Find("gh"); val != 3 {
		t.Errorf("Expected %d, Got %v", 3, val)
	}
	if val, _ = noPrefixTrie.Find("ijklmn"); val != 4 {
		t.Errorf("Expected %d, Got %v", 4, val)
	}

	if err = noPrefixTrie.Delete("gh"); err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if _, err = noPrefixTrie.Find("gh"); err == nil {
		t.Errorf("Expected %s, Got %v", "element not found", err)
	}
	if val, _ = noPrefixTrie.Find("ijklmn"); val != 4 {
		t.Errorf("Expected %d, Got %v", 4, val)
	}
}

func TestTrieOfTries(t *testing.T) {
	var triesOfTries Trie
	var bebidas Trie
	var comidas Trie

	bebidas.Insert("mate", 1)
	bebidas.Insert("cafe", 2)
	bebidas.Insert("chocolatada", 3)
	bebidas.Insert("fernet", 4)

	comidas.Insert("empanada", 1)
	comidas.Insert("tarta", 2)
	comidas.Insert("quinoa", 3)
	comidas.Insert("acelga", 4)

	triesOfTries.Insert("bebidas", bebidas)
	triesOfTries.Insert("comidas", comidas)

	val, err := triesOfTries.Find("bebidas")
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	bebidasTrie := val.(Trie)

	if val, _ = bebidasTrie.Find("mate"); val != 1 {
		t.Errorf("Expected %d, Got %v", 1, val)
	}
	if val, _ = bebidasTrie.Find("cafe"); val != 2 {
		t.Errorf("Expected %d, Got %v", 2, val)
	}
	if val, _ = bebidasTrie.Find("chocolatada"); val != 3 {
		t.Errorf("Expected %d, Got %v", 3, val)
	}
	if val, _ = bebidasTrie.Find("fernet"); val != 4 {
		t.Errorf("Expected %d, Got %v", 4, val)
	}

	val, err = triesOfTries.Find("comidas")
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	comidasTrie := val.(Trie)

	if val, _ = comidasTrie.Find("empanada"); val != 1 {
		t.Errorf("Expected %d, Got %v", 1, val)
	}
	if val, _ = comidasTrie.Find("tarta"); val != 2 {
		t.Errorf("Expected %d, Got %v", 2, val)
	}
	if val, _ = comidasTrie.Find("quinoa"); val != 3 {
		t.Errorf("Expected %d, Got %v", 3, val)
	}
	if val, _ = comidasTrie.Find("acelga"); val != 4 {
		t.Errorf("Expected %d, Got %v", 4, val)
	}

}

func checkNoElems(trie Trie, t *testing.T) {
	if trie.WordsCount != 0 {
		t.Errorf("Expected %d, Got %d", 0, trie.WordsCount)
	}
	if trie.Root != nil {
		t.Errorf("Expected %v, Got %+v", nil, trie.Root)
	}
}

func checkElem(trie Trie, count int, key string, value interface{}, t *testing.T) {
	if trie.WordsCount != count {
		t.Errorf("Expected %d, Got %d", count, trie.WordsCount)
	}
	if trie.Root.Value != nil {
		t.Errorf("Expected %v, Got %+v", nil, trie.Root.Value)
	}

	node := trie.Root
	for _, v := range key {
		if node.Children[v] == nil {
			t.Fatalf("Expected %s, Got %v", "exists", nil)
		}
		node = node.Children[v]
	}

	if node.Value != value {
		t.Errorf("Expected %v, Got %v", value, node.Value)
	}
}
