package main

import (
	"fmt"
)

type TrieNode struct {
	data map[rune]*TrieNode
}
type Trie struct {
	root *TrieNode
}

var end = '*'

func main() {
	trie := &Trie{root: &TrieNode{data: make(map[rune]*TrieNode)}}

	trie.Insert("prasanth")
	fmt.Println(trie.contains("santh"))

}
func (t *Trie) Insert(str string) {
	for i := 0; i < len(str); i++ {
		t.insertSubstring(i, str)
	}
}
func (t *Trie) insertSubstring(index int, str string) {
	tempNode := t.root
	for i := index; i < len(str); i++ {
		character := rune(str[i])
		if _, exists := tempNode.data[(character)]; !exists {

			tempNode.data[(character)] = &TrieNode{data: make(map[rune]*TrieNode)}
		}
		tempNode = tempNode.data[(character)]
	}
	tempNode.data[end] = nil
}
func (t *Trie) contains(str string) bool {
	tempNode := t.root
	for i := 0; i < len(str); i++ {
		character := rune(str[i])
		 _, exists := tempNode.data[(character)]
		if !exists {
			return false

		}
		tempNode = tempNode.data[(character)]
	}
	_, exists := tempNode.data[rune(end)]
	return exists
}
