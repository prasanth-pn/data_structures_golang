package main

import "fmt"

func main() {
	list := link{}
	list.head = &Node{data: 34}
	list.head.next = &Node{data: 56}
	list.head.next.next = &Node{data: 23}
	//list.Sort()
	list.BubbleSort()
	for list.head != nil {
		fmt.Println(list.head.data)
		list.head=list.head.next
	}

}

type Node struct {
	data int
	next *Node
}
type link struct {
	head *Node
}
func (list *link)BubbleSort(){
	temp:=list.head
	 var prev *Node
	 prev=nil
	// prev.next=nil
	for temp!=nil{
		temp2:=list.head
	for temp2.next!=prev{
		//fmt.Println(temp2.data)

prev=temp2
		temp2=temp2.next


		}
		fmt.Println(prev)
		temp=temp.next
	}
}


func (list *link)Sort(){
temp:=list.head

for temp!=nil{
	temp2:=list.head
	for temp2!=nil{
		if temp.data<temp2.data{
			t:=temp.data
			temp.data=temp2.data
			temp2.data=t
		}
		temp2=temp2.next
	}
	temp=temp.next
}

}
