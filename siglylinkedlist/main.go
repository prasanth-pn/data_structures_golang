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
	list := &single{}
	list1 := &single{}

	data1 := []int{2, 4, 6, 7, 9}
	data := []int{12, 13, 14, 15, 16}
	list.Addnode(data)
	list1.Addnode(data1)
//sum:=mergeTwoLists(list,list1)
	list1.display()
	// list.delete(12)
	// list.display()

}
//  func mergeTwoLists(list *node ,list1 *node)*node{

//  }
func (list *single) delete(data int) {
	temp := list.head
	if temp.data == data {
		list.head = temp.next
		fmt.Println("\n\n ")
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
	fmt.Println("\n ")

}
func (list *single) display() {
	temp := list.head
	if temp==nil{
		fmt.Println("the list is empty ")
		return 
	}
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

}
