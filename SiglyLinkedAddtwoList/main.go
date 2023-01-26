// this program on linked list
package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}
type singlelinkedlist struct {
	head *node
	tail *node
}

func main() {
	var n int
    list := singlelinkedlist{}
	fmt.Println("\nstart..................")
loop:
	for {
		fmt.Print("\n\n1--:Linked List and its operations\n\n2--:Convert Array to single linked list\n\n0--:EXIT\n\nChoose option  :-")
		fmt.Scan(&n)
		switch n {
		case 1:
			
			list.Start()
		case 2:
			list.convert()
		case 0:
			break loop
		default:
			fmt.Println("Choose the correct option")

		}
	}

}

func (list *singlelinkedlist) Start() {
	var n int
	fmt.Print("\n\n                         SINGLE LINKED LIST\n                       ----------------------\n\n\n")
loop:
	for {
		fmt.Print("Enter the options:-\n\n1Insert\n2Display\n3Delete Last\n4-->Delete Specific\n5Insertbeg\n6sort\n7Delete Duplicates\n0Exit\n\n")
		fmt.Print("\n\nEnter a option :-  ")
		fmt.Scan(&n)
		switch n {
		case 1:
			list.insert(0)
		case 2:
			list.display()
		case 3:
			list.deletelast()
		case 4:
			list.deletespecific()
		case 5:
			list.insertbeginnig()
		case 6:
			list.sort()
		case 7:
			list.deleteDuplicate()
		case 0:
			break loop
		default:
			fmt.Println("\nChoose the correct option................!")

		}
	}
}

func (list *singlelinkedlist) insert(a int) {
	var data int
	if a == 0 {
		fmt.Print("\n\n\nEnter the number:-  ")
		fmt.Scan(&data)
	} else {
		data = a
	}
	newNode := new(node)
	newNode.data = data
	newNode.next = nil

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *singlelinkedlist) insertbeginnig() {
	var data int
	fmt.Print("\n\n\nEnter the number:-  ")
	fmt.Scan(&data)
	newNode := new(node)
	newNode.data = data
	newNode.next = list.head

	list.head = newNode

	if list.tail == nil {
		list.tail = list.head
	}
}

func (list *singlelinkedlist) display() {
	if list.head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		fmt.Print("The list is  :-  ")
		temp := list.head
		for temp != nil {
			fmt.Print("  ", temp.data)
			temp = temp.next
		}
		fmt.Println()
	}
}
func (list *singlelinkedlist) deletelast() {
	if list.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {
		temp := list.head
		for temp != nil {
			if temp.next.next == nil {
				temp.next = nil
				list.tail = temp
				break
			}
			temp = temp.next
		}
	}
}

func (list *singlelinkedlist) deletespecific() {
	if list.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {

		var data int
		fmt.Print("\n\n\nEnter the number:-  ")
		fmt.Scan(&data)
		temp := list.head
		if temp.data == data {
			list.head = temp.next
		} else {

			for temp != nil {
				if temp.next.data == data {
					temp.next = temp.next.next
					break
				}
				temp = temp.next
			}
		}
	}

}

func (list *singlelinkedlist) sort() {
	temp := list.head

	for temp != nil {
		temp2 := list.head
		for temp2 != nil {
			if temp.data < temp2.data {
				t := temp.data
				temp.data = temp2.data
				temp2.data = t
			}
			temp2 = temp2.next
		}
		temp = temp.next
	}

}

func (list *singlelinkedlist) deleteDuplicate() {
	key := list.head
	for key != nil {
		prev := key
		temp := key.next

		for temp != nil {
			if temp.data == key.data {
				prev.next = temp.next
				if temp == list.tail {
					list.tail = prev
				}

			}
			prev = temp
			temp = temp.next

		}
		key = key.next

	}

}

func (list *singlelinkedlist) convert() {
	var n int
	fmt.Print("\n\nEnter the the size of array  :-")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("\n\nEnter elaements  :-")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Println("\nThe array you enterd  :  ", arr, "\n\nLet convert to Linked List   :-  ")
	for _, val := range arr {
		list.insert(val)
	}
	fmt.Println("After the convertion  :  ")
	list.display()
}
