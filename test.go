package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
    data int
    left *Node
    right *Node
}

func main() {
    a := &Node{data: 26}
    a.left = &Node{data: 10}
    a.right = &Node{data: 3}
    a.right.left = &Node{data: 3}
    a.right.right = &Node{data: 3}
    a.left.left = &Node{data: 4}
    a.left.right = &Node{data: 6}

    b := &Node{data:6}
    b.left = &Node{data: 46}
    b.left.right = &Node{data: 10}
var resulta string
var resultb string
    inorder(a,&resulta)
    fmt.Println("\n-----------------")
    inorder(b,&resultb)

   fmt.Println(resulta,resultb)
  fmt.Println( strings.Contains(resulta,resultb))

 
}
func inorder(root *Node,result *string){
    if root==nil{
        return 
    }
    inorder(root.left,result)
   *result+=strconv.Itoa(root.data)
    fmt.Println(root.data)
    inorder(root.right,result)
   // return result

}
