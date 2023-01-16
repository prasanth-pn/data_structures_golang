package main

import "fmt"

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
	list.display()
}

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
			c.tail.pre = c.tail
			c.tail.next = newnode

		}
		c.tail = newnode
	}
}
func (c *double) display() {
	temp := c.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
