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
	data := []int{25, 15, 10, 4, 12, 22, 18, 24, 50, 35, 31, 44, 70, 66, 90}
	for n < 8 {
		fmt.Println(" \n select the options \n1.for  insert\n2.inorder\n3.for preorder\n4.for postorder\n5 for remove\n--------------\n ")
		fmt.Scan(&n)
		switch n {
		case 1:
			for _, k := range data {

				tree.insert(k)
			}
			fmt.Println("\n insertionnnnn completed", tree.root)
		case 2:
			fmt.Println("Enter the value to search ")
			value := 0
			fmt.Scan(&value)
			fmt.Println(tree.contains(value))
		case 3:
			fmt.Println("enter the number to remove from the tree ")
			g:=0
			fmt.Scan(&g)
			tree.remove(g)

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
				if data > currentnode.data {
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
	currentnode := tree.root
	for currentnode != nil {
		if data < currentnode.data {
			currentnode = currentnode.left
		} else if data > currentnode.data {
			currentnode = currentnode.right
		} else {
			return true

		}

	}
	return false

}
func (tree *tree) remove(data int) {
	removeHelper(data, tree.root, nil)

}

func removeHelper(data int, currentnode *node, parentnode *node) {

	for currentnode != nil {
		if data < currentnode.data {
			parentnode = currentnode
			currentnode = currentnode.left

		} else if data > currentnode.data {
			parentnode = currentnode
			currentnode = currentnode.right

		} else {
			if currentnode.left != nil && currentnode.right != nil {
				currentnode.data = getmin(currentnode.right)
				removeHelper(currentnode.data, currentnode.right, currentnode)

			} else {
				if parentnode == nil {
					if currentnode.right == nil {
						currentnode = currentnode.left
					} else {
						currentnode = currentnode.right
					}

				} else {
					if parentnode.left == currentnode {
						if currentnode.right == nil {
							parentnode.left = currentnode.left
						} else {
							parentnode.left = currentnode.right
						}
					} else {
						if currentnode.right == nil {
							parentnode.right = currentnode.left
						} else {
							parentnode.right = currentnode.right
						}

					}
				}
			}
			break

		}
	}
}
func getmin(currentnode *node) int {
	if currentnode.left == nil {
		return currentnode.data
	} else {
		return getmin(currentnode.left)

	}
}
