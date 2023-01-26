package main

import "fmt"

var data = []int{12, 23, 45, 21, 1, 567, 345, 23}

func main() {
	var i int
	fmt.Print(i)
	LinearSearch(i, 23)

}
func LinearSearch(i, v int) {
	k := len(data)
	if i == k {
		return
	}
	if data[i] == v {
		fmt.Println("the data is found at a position", i)

	}
	i++
	LinearSearch(i, v)

}
