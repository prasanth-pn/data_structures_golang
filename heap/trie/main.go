package main

import "fmt"

type TrieNode struct {
	data map[rune]*TrieNode
}
type trie struct{
root *TrieNode
}
var end='*'

func main() {
	Trie:=new(trie)
	Trie.root=&TrieNode{
		data:map[rune]*TrieNode{},
	}
	Trie.insertToTrie("prasanth")
	fmt.Println(Trie.contains("asanth"))

}
func (t *trie)insertToTrie(str string){
	t.populateSufixTrie(str)
}
func (t *trie)populateSufixTrie(str string){
	for i:=0;i<len(str);i++{
		t.insertSubstring(i,str)
	}
}
func (t *trie)insertSubstring(index int,str string){
	tempNode:=t.root
	for i:=index;i<len(str);i++{
		character:=str[i]
		if _,exist:=tempNode.data[rune(character)];!exist{
			newTrieNode:=&TrieNode{data:make(map[rune]*TrieNode)}
			tempNode.data[rune (character)]=newTrieNode
		}
		
		tempNode=tempNode.data[rune(character)]

	}
	tempNode.data[end]=nil
}
func (t *trie)contains(str string)bool{
	tempNode:=t.root
	var character rune 
	fmt.Println(tempNode)
	for i:=0;i<len(str);i++{
		character=rune(str[i])
		if _,exist:=tempNode.data[rune(character)];!exist{
			return false
		}
		tempNode=tempNode.data[rune(character)]
		
	}
	_,exist:=tempNode.data[rune(end)]
	return exist
}
