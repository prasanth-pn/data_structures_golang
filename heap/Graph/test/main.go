package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}
type root struct {
	root *node
}

func main() {
	tree := root{}

	data := []int{25, 15, 10, 4, 12, 22, 18, 25, 50, 35, 31, 44, 70, 66, 90}
	for _, val := range data {
		tree.insert(val)
	}

	// //	fmt.Println(tree.root.right.right.data)
	// inorder(tree.root)
	// fmt.Println("\n---------------------")

	// inorder(tree.root)
	// fmt.Println("\n----------------------")

	// preorder(tree.root)

	fmt.Println("\n---------------------")
	deleteNode(tree.root, 25)
	postorder(tree.root)

}

func (list *root) insert(value int) {

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
func preorder(temp *node) {
	if temp != nil {
		fmt.Print(temp.data, "\t")
		preorder(temp.left)
		preorder(temp.right)
	}

}
func postorder(temp *node) {
	if temp != nil {
		postorder(temp.left)
		postorder(temp.right)
		fmt.Print(temp.data, "\t")
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
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		minNode := findMin(root.right)
		root.data = minNode.data
		root.right = deleteNode(root.right, minNode.data)

	}

	return root
}

func findMin(root *node) *node {
	for root.left != nil {
		root = root.left
	}
	return root
}
