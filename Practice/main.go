package main

import "fmt"

type node struct {
	key   string
	value string
	next  *node
}

const size int = 100

var table [size]*node

func main() {
	insert("dada", "prasanth")
	insert("mama", "devil")
	insert("my", "dog")
	insert("dada", "vee")
	k, ok := get("dada")
	fmt.Println(k, ok)
	Delete("dada")
	k, ok = get("dada")
	fmt.Println(k, ok)

}
func hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {

		sum += int(key[i])
	}
	return sum % size
}
func insert(key, value string) {
	h := hash(key)
	newnode := new(node)
	newnode.key = key
	newnode.value = value
	newnode.next = table[h]
	table[h] = newnode
}

func get(key string) (string, bool) {
	h := hash(key)
	for n := table[h]; n != nil; n = n.next {
		if n.key == key {
			return n.value, true
		}
	}
	return "", false
}

func Delete(key string) {
	h := hash(key)
	if table[h] == nil {
		fmt.Println("the table is empty")
		return
	}
	if table[h].key == key {
		table[h] = nil
		return
	}

	for n := table[h]; n != nil; n = n.next {
		if n.key == key {
			n.next = n.next.next
		}
	}
}
