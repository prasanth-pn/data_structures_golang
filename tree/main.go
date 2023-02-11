package main

import "fmt"

type node struct {
	data        int
	left, right *node
}
type bst struct {
	root *node
}

func main() {
	var n int
	data := []int{25, 15, 10, 4, 12, 22, 18, 24, 50, 35, 31, 44, 70, 66, 90}
	list := bst{}
	for n < 8 {
		fmt.Println("\nenter the option to do the tree operation\n 1.for insert \n 2.for inorder\n.3 for preorder \n 4.postorder \n 6 delete \n--------------------------------- ")
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
				if value > temp.data {
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
		} else if root.right == nil {
			return root.left
		}
		minNode := findMinNode(root.right)
		root.data = minNode.data
		root.right = deleteNode(root.right, minNode.data)
		fmt.Println(root.right,"end")
	}
	return root
}

func findMinNode(node *node) *node {
	for node.left != nil {
		node = node.left
	}
	return node
}
