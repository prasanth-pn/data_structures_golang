package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

func (head *Node) BubbleSort() *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, current, tail *Node
	tail = nil
	for current = head; current.Next != tail; {
		for prev, current = nil, head; current.Next != tail; {
			if current.Value > current.Next.Value {
				if prev != nil {
					prev.Next = current.Next
				} else {
					head = current.Next
				}
				current.Next, current.Next.Next = current.Next.Next, current
				prev = current.Next
			} else {
				prev = current
				current = current.Next
			}
		}
		tail = current
	}
	return head
}

func main() {
	head := &Node{5, &Node{3, &Node{2, &Node{1, nil}}}}
	for head!=nil{
		fmt.Println(head.Value)
		head=head.Next
	}
	head = head.BubbleSort()

	for curr := head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}
}
