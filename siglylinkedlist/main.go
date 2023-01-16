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

	//list.delete(13)

	//list.display()
	//list.delete(12)
	//list.delete(15)
	//list.delete(16)

	list.insert(150, 16)
	list.display()

}
// search the value and add new node
func (c *single) insert(data int, search int) {
	newnode := new(node)
	newnode.data = data
	temp := c.head
	if search == temp.data {
		temp.data = data
		return

	}
	for temp.next != nil {
		if temp.next.data == search {
			newnode.next = temp.next
			temp.next = newnode
			return

		}
		temp = temp.next
	}
	fmt.Println("the number is not available")
}
//---------------------------delete the node 
func (list *single) delete(data int) {
	temp := list.head
	if temp.data == data {
		list.head = temp.next
		return
	} else {
		temp = list.head
		for temp != nil {
			if temp.next.data == data {
				if temp.next == list.tail {
					list.tail = temp
					list.tail.next = nil

				} else {
					temp.next = temp.next.next
					return
				}
			}
			temp = temp.next
		}
	}

}
//----------------------------- add the node    

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
//-------------------------------------display the values
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
