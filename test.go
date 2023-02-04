package main

import "fmt"
type Node struct{
	data int
	next *Node
}
type queue struct{
	front *Node
	rear *Node
}
func main() {
	list:=queue{}
	i:=1
	for i<5{
		fmt.Println(" enter the data to store in queue")
		n:=0
		fmt.Scan(&n)
		list.Enqueue(n)
		i++
	}
	list.front=list.front.next



	for list.front!=nil{
		fmt.Println(list.front.data)
		list.front=list.front.next
	}
	
}
func(list *queue) Enqueue(n int){
	newnode:=new(Node)
	newnode.data=n
	newnode.next=nil
	if list.front==nil{
		list.front,list.rear=newnode,newnode

	}else{
		list.rear.next=newnode
		list.rear=newnode

	}
}























// package main

// import "fmt"
// type Node struct{
// 	data int
// 	next *Node
// }
// type stack struct{
// 	top *Node
// }

// func main() {
// 	list:=stack{}
// 	i:=1
// 	for i<=5{
// 		fmt.Println("Enter the number to store in the stack")
// 		n:=0
// 		fmt.Scan(&n)
// 		list.Push(n)
// 		i++

// 	}
// 	for list.top!=nil{
// 		fmt.Println(list.top.data)
// 		list.top=list.top.next

// 	}
	
// }
// func (list *stack)Push(v int){
// 	newnode:=new(Node)
// 	newnode.data=v
// 	newnode.next=nil
// 	if list.top==nil{
// 		list.top=newnode

// 	}else{
// 		newnode.next=list.top
// 		list.top=newnode
// 	}
// }