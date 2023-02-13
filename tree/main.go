package main

import (
	"fmt"
	"math"
)

type node struct {
	data        int
	left, right *node
}
type bst struct {
	root *node
}

func main() {
	var visited = make(map[int]bool)
	var n int
	data := []int{25, 10, 4, 22, 22, 18, 24, 50, 25, 31, 44, 70, 66, 90}
	list := bst{}
	for n < 9 {
		fmt.Println("\nenter the option to do the tree operation\n 1.for insert \n 2.for inorder\n.3 for preorder \n 4.postorder \n 5 delete \n 6 for delete duplicates\n 7 for closest value \n--------------------------------- ")
		fmt.Scan(&n)
		switch n {
		case 1:
			for _, value := range data {
				fmt.Println(value)
				list.insert(value)
			}
		case 2:
			fmt.Println("in order list is ")
			inorder(list.root)
		case 3:
			fmt.Println("preorder list is ")
			preorder(list.root)
			fmt.Println()
			fmt.Println(list.root.left.data)
		case 4:
			fmt.Println("The post order list is ")
			postorder(list.root)
		case 5:
			fmt.Println("enter tht number to delete")
			v := 0
			fmt.Scan(&v)
			deleteNode(list.root, v)
		case 6:
			list.deleteDuplicate(list.root, visited)
		case 7:
			fmt.Println("Enter the number to find the closest value")

			key := 26
			diff := 0
			fmt.Println(&diff)
			val := 0
			fmt.Scan(&key)
			closestValue(list.root, key, &diff, &val)
			fmt.Println(val,diff)
		}
	}
	//list.insert()

}
func (list *bst) insert(value int) {

	// fmt.Println(" enter the value to store in bst")
	// fmt.Scan(&value)
	newnode := new(node)
	newnode.data = value
	if list.root == nil {
		list.root = newnode
	} else {
		temp := list.root
		for {
			if value < temp.data {
				if temp.left == nil {
					temp.left = newnode
					break

				} else {
					temp = temp.left
				}
			} else {
				if value >= temp.data {
					if temp.right == nil {
						temp.right = newnode
						break
					} else {
						temp = temp.right

					}
				}
			}
		}
	}

}

func inorder(temp *node) {

	if temp != nil {
		inorder(temp.left)
		fmt.Print(temp.data, "\t")
		inorder(temp.right)
	}
}

func preorder(tree *node) {
	if tree == nil {
		return
	}
	fmt.Print(tree.data, "\t")
	preorder(tree.left)
	preorder(tree.right)

}

func postorder(tree *node) {
	if tree != nil {
		postorder(tree.left)
		postorder(tree.right)
		fmt.Print(tree.data, "\t")
	}
}
func deleteNode(root *node, key int) *node {
	if root == nil {
		return root
	}
	if key < root.data {
		root.left = deleteNode(root.left, key)

	} else if key > root.data {
		root.right = deleteNode(root.right, key)

	} else {
		if root.left == nil {
			fmt.Println("left")
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		minNode := findMinNode(root.right)
		root.data = minNode.data
		root.right = deleteNode(root.right, minNode.data)

	}
	return root
}

func findMinNode(node *node) *node {
	for node.left != nil {
		node = node.left
	}
	return node
}
func (list *bst) deleteDuplicate(root *node, visited map[int]bool) {
	if root == nil {
		return
	}
	list.deleteDuplicate(root.left, visited)
	if !visited[root.data] {
		visited[root.data] = true
	} else {
		list.delete(root.data)
	}
	list.deleteDuplicate(root.right, visited)

}
func (tree *bst) delete(key int) {
	deleteNode(tree.root, key)

}
func closestValue(root *node, key int, dif, val *int) {
	if root == nil {
		return 

	}
	//*dif=7999999

	
	closestValue(root.left, key, dif, val)
	if *dif > root.data-key {
		*dif = int(math.Abs(float64(root.data - key)))
		*val = root.data
	}
	closestValue(root.right, key, dif, val)

}
