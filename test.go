package main

import "fmt"

func main() {
	array :=[]int{17,28,3,7,5,642,44}
	i:=5
	array[0],array[1]=array[1],array[0]
	fmt.Println(array[:i])
}

