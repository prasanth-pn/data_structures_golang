package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type queue struct {
	front *Node
	rear  *Node
}

func main() {
	var n int
	q := queue{}

	for n < 7 {
		fmt.Println("\n selct the prefered option\n1.Enqueue\n2.Dequeue\n3.Peek\n4.Display\n ----------------------------")
		fmt.Scan(&n)
		switch n {
		case 1:
			q.Enqueue()
		case 2:
			q.Dequeue()
		case 3:
			q.Peek()
		case 4:
			q.Display()
		default:
			fmt.Println(" enter the correct option b/w 1to 4")
		}
	}
	fmt.Println("program is exited successfully")

}
func (q *queue) Enqueue() {
	var value int
	fmt.Println(" Enter the value to store in queue")
	fmt.Scan(&value)
	newnode := new(Node)
	newnode.next = nil
	newnode.data = value
	if q.front == nil && q.rear == nil {
		q.front = newnode
		q.rear = newnode
	} else {

		q.rear.next = newnode
		q.rear = newnode
		//return
	}

}
func (q *queue) Dequeue() {
	if q.front == nil {
		fmt.Println(" the list empty")
		return
	}
	q.front = q.front.next

}
func (q *queue) Peek() {
	if q.front == nil {
		fmt.Println(" the list is empty ")
		return
	}
	fmt.Println(q.front.data)

}
func (q *queue) Display() {
	if q.front == nil {
		fmt.Println("The queue is empty ")
		return
	}
	fmt.Println("the values are listed")
	temp := q.front
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

}
