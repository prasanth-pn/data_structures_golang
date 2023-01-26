package main

import (
	"fmt"
	"sort"
)

type Node struct {
	data int
	next *Node
}
type Single struct {
	head   *Node
	length int
}

// -----------prepend--------
func (list *Single) Prepend(value int) {
	//fmt.Println("enter the value to add in linked list")
	//fmt.Scan(&value)
	newnode := new(Node)
	newnode.data = value
	newnode.next = nil
	if list.head != nil {
		newnode.next = list.head
		list.head = newnode
		list.length++

	} else {
		list.head = newnode
		list.length++
	}
	return

}
func (list *Single) Display() {
	temp := list.head
	if temp == nil {
		fmt.Println("list contains no data")
		return

	}
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func (list *Single) DeleteValue(v int) {
	if list.length == 0 { //checking length of list  0 or not if has 0  node return
		return
	}
	if list.head.data == v { // if head have the value then
		list.head = list.head.next
		list.length--
		return

	}
	temp := list.head
	for temp.next.data != v {
		if temp.next.next == nil {
			return

		}
		temp = temp.next

	}
	temp.next = temp.next.next
	list.length--

}

func main() {
	data := []int{345, 23, 23, 345, 23, 345, 23}
	sort.Ints(data)

	fmt.Println(data)
	list := Single{}
	for _, v := range data {
		list.Prepend(v)
	}
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("the sum of linked list")
	list.Sum()

	//list.DeleteValue(23)
	//list.Display()
	//list.Sum()
	fmt.Println("the sorted linked list")
	//list.SortAscend()
	list.Display()

	list.DuplicateDelete()
	list.reverseList()
	list.Display()
	//list.Sum()
}
func (l *Single) SortAscend() {
	i := l.head
	j := l.head.next

	for i != nil {

		for j != nil {
			if i.data > j.data {
				temp := i.data
				i.data = j.data
				//fmt.Println(i.data,j.data)
				j.data = temp
				fmt.Println(i.data, j.data)
			}
			j = j.next
		}
		i = i.next

	}
}
func (l *Single) Sum() {
	var sum = 0
	temp := l.head
	for temp != nil {
		sum = sum + temp.data
		temp = temp.next
	}
	fmt.Println(sum)
}

func (l *Single) DuplicateDelete() *Node {
	check := make(map[int]bool)
	current := l.head
	prev := &Node{}      // initialize dummy node
	for current != nil {
		if _, ok := check[current.data]; ok {
			prev.next = current.next
		} else {
			check[current.data] = true
			prev = current
		}
		current = current.next
	}
	return l.head

}

func (l *Single)reverseList() *Node {
    var prev *Node
    current :=l. head
    var next *Node

    for current != nil {
        next = current.next
        current.next = prev
        prev = current
        current = next
    }

    return prev
}
