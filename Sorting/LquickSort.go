package main

import (
	"fmt"
)

// Node represents a node in a linked list
type Node struct {
	Value int
	Next  *Node
}

// List is a linked list
type List struct {
	Head *Node
	Tail *Node
}

// NewList creates a new linked list
func NewList() *List {
	return &List{}
}

// Append adds a node to the end of the linked list
func (l *List) Append(value int) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	l.Tail = node
}

// Partition partitions the linked list around a pivot
func (l *List) Partition(pivot int) (*List, *List) {
	var lessList, greaterList List
	node := l.Head
	for node != nil {
		if node.Value < pivot {
			lessList.Append(node.Value)
		} else {
			greaterList.Append(node.Value)
		}
		node = node.Next
	}
	return &lessList, &greaterList
}

// Concat concatenates two linked lists
func (l *List) Concat(other *List) {
	if other.Head == nil {
		return
	}
	if l.Head == nil {
		l.Head = other.Head
		l.Tail = other.Tail
		return
	}
	l.Tail.Next = other.Head
	l.Tail = other.Tail
}

// Quicksort sorts a linked list using the quicksort algorithm
func (l *List) Quicksort() *List {
	if l.Head == l.Tail {
		return l
	}
	lessList, greaterList := l.Partition(l.Head.Value)
	lessList.Quicksort()
	greaterList.Quicksort()
	lessList.Concat(greaterList)
	lessList.Concat(&List{Head: l.Head, Tail: l.Tail})
	return lessList
}

func main() {
	list := NewList()
	list.Append(5)
	list.Append(3)
	list.Append(7)
	list.Append(1)
	list.Append(9)
	list = list.Quicksort()
	node := list.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}
