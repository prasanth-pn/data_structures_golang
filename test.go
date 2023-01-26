package main

import "fmt"

func main() {
	array:=[]int{23,34,12,234}
	array[0],array[1]=array[1],array[0]
	fmt.Println(array)
}

