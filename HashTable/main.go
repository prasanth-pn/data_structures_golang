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
	var value string
	var key string
	var n int
	for n < 7 {
		fmt.Println("\nEnter the Options \n1.Inser the value to the hash table \n2.search the value\n 3.delete the value")
		fmt.Scan(&n)
		switch n {
		case 1:

			fmt.Println(" enter the key  to store in the hash table")
			fmt.Scan(&key)
			fmt.Println("enter the value to store in the hash table ")
			fmt.Scan(&value)
			insert(key, value)
		case 2:
			fmt.Println("Enter the key value to search")
			var data string
			fmt.Scan(&data)
			ok, lk := get(data)
			fmt.Println("the string value is ", ok, "\n is available :", lk)
		case 3:
			fmt.Println("Enter the key value to delete the value")
			fmt.Println(key)
			fmt.Scan(&key)
			fmt.Println(key)

			Delete(key)
		default:
			fmt.Println("Enter the correct option ")
		}
	}
	fmt.Println("program successfully completed")
}
func hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % size
}
func insert(key string, value string) {
	h := hash(key)
	//fmt.Println(value, h)
	newnode := new(node)
	newnode.key = key
	newnode.value = value
	newnode.next = table[h]
	table[h] = newnode 
	//fmt.Println(table[h])
	//fmt.Println(newnode.key,newnode.value)
}
func get(key string) (string, bool) {
	//fmt.Println(key,"search")
	h := hash(key)

	for n := table[h]; n != nil; n = n.next {
		if n.key == key {
			fmt.Println(n.key,n.value)
			return n.value, true
		}
	}
	return "", false

}
func Delete(key string) {
	h := hash(key)
	if table[h] == nil {
		return
	}
	if table[h].key == key {
		fmt.Println(table[h].value)
		return
	}
	for n := table[h]; n.next != nil; n = n.next {
		if n.key == key {
			n.next = n.next.next
			fmt.Println("availabel", n.value)
			return
		}
	}

}
