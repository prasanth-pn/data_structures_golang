package main

import (
	"fmt"
)

type node struct {
	data int
	pre  *node
	next *node
}
type double struct {
	head *node
	tail *node
}

func main() {
	list := double{}
	data := []int{12, 13, 14, 15}
	//fmt.Println(data)
	list.addnode(data)

	//list.search(12)
	list.insert(120, 15)
	list.insert(200, 14)
	list.display()
	fmt.Println("\n\n\n the reverse order is \n ")
	list.reverseOrder()
	fmt.Println("the ordered values are")
	list.order()

}
func (c *double) order() {
	temp := c.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (c *double) reverseOrder() {
	temp := c.tail

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.pre
	}

}
func (c *double) insert(value int, search int) {
	newnode := new(node)
	newnode.data = value
	newnode.pre = nil
	newnode.next = nil
	temp := c.head
	if temp.data == search {
		temp.pre = newnode
		newnode.next = temp
		c.head = newnode
		fmt.Println("node added successfully")
		return
	}

	for temp.data != search {
		
		if temp.next.data == search {
			newnode.next = temp.next
			newnode.pre = temp
			temp.next.pre = newnode
			temp.next = newnode

			break
		}

		temp = temp.next
	}

}

// ----------------------------add node
func (c *double) addnode(data []int) {
	for _, v := range data {

		newnode := new(node)
		newnode.data = v
		newnode.next = nil
		newnode.pre = nil
		if c.head == nil {
			c.head = newnode
			c.tail = newnode

		} else {
			newnode.pre = c.tail

			c.tail.next = newnode

		}
		c.tail = newnode
	}
}

// --------------------------------display values
func (c *double) display() {
	temp := c.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

// -------------------------------insert node-----
func (c *double) search(k int) {
	temp := c.head

	if temp.data == k {
		fmt.Println("the data is available in first node")
		return
	}
	for temp.next != nil {
		if temp.next.data == k {
			fmt.Println("the data is available")

			return
		}
		temp = temp.next
	}

	fmt.Println("the data is not available")

}
