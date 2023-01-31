package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int64{34, 23, 234, 12, 5}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println("median is : ",arr[len(arr)/2])
}
