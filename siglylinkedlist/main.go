package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}
type single struct {
	head *node
	tail *node
}

func main() {
	list := single{}

	data := []int{12, 13, 14, 15, 16}
	list.Addnode(data)
	list.display()
	list.delete(12)
	list.display()

}
func (list *single) delete(data int) {
	temp := list.head
	if temp.data == data {
		list.head = temp.next
		fmt.Println("\n\n\n\n ")
	}
	return

}

func (list *single) Addnode(data []int) {
	for _, v := range data {

		newnode := new(node)
		newnode.data = v
		newnode.next = nil
		if list.head == nil {
			list.head = newnode
			list.tail = newnode
		} else {
			list.tail.next = newnode
		}
		list.tail = newnode

	}
	fmt.Println("\n\n\n\n ")

}
func (list *single) display() {
	temp := list.head
	if list.head == nil {
		fmt.Println("list is empty")
	}
	temp = list.head

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

}
