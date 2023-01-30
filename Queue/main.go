package main

import "fmt"

var Queue [3]int
var rear int = -1

var front int = -1

func main() {
	var n int
	for n < 7 {
		fmt.Println("\n Enter the correct option \n1.Enqueue\n2.Dequeue\n3.Peek\n4.Display\n-------------------")
		fmt.Scan(&n)
		switch n {
		case 1:
			Enqueue()
		case 2:
			Dequeue()
		case 3:
			Peek()
		case 4:
			Display()
		default:
			fmt.Println(" enter the correct option")
		}

	}

}
func Enqueue() {
	if rear > len(Queue) {
		fmt.Println("The stack is overflow")
		return
	}
	var value int
	fmt.Println("enter the value to store in the queue")
	fmt.Scan(&value)
	if rear == len(Queue)-1 {
		fmt.Println(" the queue in overflow condition ")
	} else if rear == -1 && front == -1 {
		rear = 0
		front = 0
		Queue[rear] = value

	} else {
		rear++
		Queue[rear] = value
	}

}
func Dequeue() {
	if front == -1 && rear == -1 {
		fmt.Println(" the Queue is empty ")
		return
	} else if front == rear {

		front = -1
		rear = -1
		return
	} else {
		front++
		return
	}

}
func Peek() {
	if front == -1 && rear == -1 {
		fmt.Println(" the Queue is empty")
		return
	}
	fmt.Println(Queue[front])

}
func Display() {
	if rear == -1 || front == -1 {
		fmt.Println("The queue is empty")
		return
	}
	fmt.Println("The data in the queue are below")
	for i := front; i <= rear; i++ {
		fmt.Println(Queue[i])
	}

}
