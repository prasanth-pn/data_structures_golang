package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}
type tree struct {
	root *node
}

func main() {
	tree := tree{}
	var n int
	data := []int{25, 25, 10, 4, 12, 22, 18, 24, 50, 25, 31, 44, 70, 66, 90}
	for n < 8 {
		fmt.Println(" \n select the options \n1.for  insert\n2.Search\n3.for remove\n4.for postorder\n5 for preorder\n6.for Inorder\n--------------\n ")
		fmt.Scan(&n)
		switch n {
		case 1:
			for _, k := range data {

				tree.insert(k)
			}
			fmt.Println("\n insertionnnnn completed", tree.root)
		case 2:
			fmt.Println("Enter the value to search ")
			value := 012
			fmt.Scan(&value)
			fmt.Println(tree.contains(value))
		case 3:
			fmt.Println("enter the number to remove from the tree ")
			g:=0
			fmt.Scan(&g)
			tree.remove(g)
		case 4:
			postorder(tree.root)

		case 5:
			preorder(tree.root)
		case 6:
			inorder(tree.root)

		}
	}

}
func (tree *tree) insert(data int) {
	fmt.Println(data)
	newnode := new(node)
	newnode.data = data
	currentnode := tree.root
	if currentnode == nil {
		tree.root = newnode

	} else {
		currentnode = tree.root
		for {
			if data < currentnode.data {
				if currentnode.left == nil {
					currentnode.left = newnode
					break
				} else {
					currentnode = currentnode.left
				}
			} else {
				if data >= currentnode.data {
					if currentnode.right == nil {
						currentnode.right = newnode
						break

					} else {
						currentnode = currentnode.right
					}
				}
			}
		}
	}

}
func (tree *tree) contains(data int) bool {
	temp:=tree.root
	for temp!=nil{
		if data<temp.data{
			temp=temp.left
		}else if data>temp.data {
			temp=temp.right
		}else{
			return true
		}
	}
	return false

}
func (tree *tree) remove(data int) {
	removeHelper(tree.root,data)

}

func removeHelper(currentnode *node,key int) *node{
	

	if key <currentnode.data{
		currentnode.left=removeHelper(currentnode.left,key)
		
	}else if key>currentnode.data{
		currentnode.right=removeHelper(currentnode.right,key)
	}else{
		if currentnode.left==nil{
			return currentnode.right
		}
		if currentnode.right==nil{
			return currentnode.left
		}
		
			minNode:=getmin(currentnode.right)
			currentnode.data=minNode.data
			currentnode.right=removeHelper(currentnode.right,minNode.data)
		
	}
	return currentnode


	
}

func getmin(currentnode *node)*node  {
	for currentnode.left!=nil{
		currentnode=currentnode.left
		
	}
	return currentnode
	
}


func postorder(temp *node){
	if temp==nil{
		return
	}
	postorder(temp.left)
	postorder(temp.right)
	fmt.Print(temp.data," ")

}
func preorder(temp *node){
	if temp==nil{
		return
	}
	fmt.Print(temp.data," ")
	preorder(temp.left)
	preorder(temp.right)
}
func inorder(temp *node){
	if temp==nil{
		return
	}
	inorder(temp.left)
	fmt.Print(temp.data," ")
	inorder(temp.right)
}
// func delteDuplicates(temp *node,){
// 	if temp==nil{
// 		return 
// 	}

// }
