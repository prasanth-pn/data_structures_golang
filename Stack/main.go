package main

import (
	"fmt"
)

var array [5]int
var top int = -1

func main() {
	var n int
	for n < 7 {
		fmt.Println(" \n enter the  option to use the operation \n 1 for push \n 2 for pop \n 3 for peek .\n 4 for display\n 7 for exit \n -----------------------------\n ")
		fmt.Scan(&n)
		switch n {
		case 1:
			Push()
		case 2:
			Pop()
		case 3:
			Peek()
		case 4:
			Display()
		default:
			fmt.Println("this option is not available ")

		}
	}
	fmt.Println(" program exited successfully")

}
func Push() {

	if top == len(array)-1 {
		fmt.Println(" stack is in overflow condition")

	} else {
		top++
		fmt.Println("enter the value to insert to stack ")
		value := 0
		fmt.Scan(&value)
		array[top] = value
	}
}

func Pop() {
	var n int

	if top == -1 {
		fmt.Println(" the stack is in empty condition \n ")
	} else {
		n = array[top]
		top--
		fmt.Printf(" the deleted number is %d in the stack\n", n)
	}
}

func Peek() {
	if top == -1 {
		fmt.Println(" the stack is empty")
	} else {
		fmt.Println(" the value of the stack: ", array[top])
	}
}

func Display() {
	if top == -1 {
		fmt.Println(" the stack is empty")
	}
	for i := top; i >= 0; i-- {
		fmt.Print(array[i], "\t")
	}
}
