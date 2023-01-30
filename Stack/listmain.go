package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	Head *Node
	l    int
}

func main() {
	S := Stack{}
	fmt.Println(S.l, S.Head)
	var n int

	for n < 7 {
		fmt.Println(" \n Enter the option do \n1.for Push \n2. for Pop\n3.for Peek/search/ \n4.for Display ")
		fmt.Scan(&n)
		switch n {
		case 1:
			S.Push()
		case 2:
			S.Pop()
		case 3:
			S.Peek()
		case 4:
			S.Display()
		default:
			fmt.Println(" please enter the correct option")

		}
	}

}
func (S *Stack) Pop() {
	if S.l == 0 {
		fmt.Println(" stack is empty")
		return
	}
	S.Head = S.Head.next
	S.l--

}
func (S *Stack) Push() {
	var value int
	fmt.Println(" enter the value to push into the stack")
	fmt.Scan(&value)
	newnode := new(Node)
	newnode.data = value
	newnode.next = nil
	if S.Head == nil {
		S.Head = newnode
		S.Head.next = nil
		S.l++

	} else {
		newnode.next = S.Head
		S.Head = newnode
		S.l++
	    return
	}

}

func (S *Stack) Peek() {
	fmt.Println(S.Head.data)
	return

}
func (S *Stack) Display() {
	if S.l == 0 {
		fmt.Println(" empty stack no values are stored in stack ")
		return

	}
	fmt.Println(" display the stack values")
	temp := S.Head
	for temp != nil {
		fmt.Print(temp.data, "\t")
		temp = temp.next
	}

}
