package main

import (
	"fmt"
	"io"
	"os"
)



type BinaryNode struct {
	right *BinaryNode
	left *BinaryNode
	data float64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree)insert(data float64)*BinaryTree{
	if t.root==nil{
		t.root=&BinaryNode{data :data,left:nil,right:nil}
		
	}else{
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode)insert(data float64){
	if n==nil{
		return
	}else if data<=n.data{
		if n.left==nil{
			n.left=&BinaryNode{data:data,left:nil,right:nil}
		}else{
			n.left.insert(data)
		}
	}else{
		if n.right==nil{
			n.right=&BinaryNode{data:data,left:nil,right:nil}

		}else{
			n.right.insert(data)
		}
	}
}
//tree datastructure
func main() {
	tree := &BinaryTree{}
    tree.insert(10).
        insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
        insert(50).
		insert(60).
		insert(55).
        insert(85).
		insert(15).
		insert(5).
        insert(-10)
    print(os.Stdout, tree.root, 0, 'M')
//	slice :=[]int{45,23,22,12,34,13,33,24,56}
	// for _,v:=range slice{

	// }


	

}
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
    if node == nil {
        return
    }
 
    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.data)
    print(w, node.left, ns+2, 'L')
    print(w, node.right, ns+2, 'R')
}