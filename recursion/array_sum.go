package main

import (
	"fmt"
)

func main() {
	var i int
	array := []int{20, 3, 2, 100, 50}
	sum := 0
	j:=0
	data := Sort(array, i,j, sum)
	fmt.Println(data,array)

}
func Sort(array []int, i int, sum int) int {
	if i >=len(array)-1 {
		return sum

	}
	if array[i]>array[i+1]{
		array[i],array[i+1]=array[i+1],array[i]
	}
	sum += array[i]
	//fmt.Println(sum)
	return Sort(array, i+1, sum) 
	
}
