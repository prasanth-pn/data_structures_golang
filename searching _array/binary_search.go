package main

import (
	"fmt"
	"sort"
)

func main() {
	var data = []int{45, 34, 23, 54, 767, 21, 1, 342}
	sort.Ints(data)
	fmt.Println(data)


	var i int
	s:=767
h:=	binarySearchRecursive(data , s, i, len(data)-1)
if h<0{
    fmt.Println("the number is not available in the array")
}else{
    fmt.Println(" the number is available at the position  :  ",h)
}

	//BinarySearch(data, i, len(data)-1, s)
}

func binarySearchRecursive(arr []int, x int, low int, high int) int {
    if low > high {
        return -1
    }
    mid :=(low + high) / 2
    if arr[mid] == x {
		//fmt.Println("the number is present in the position ",mid)
        return mid
    } else if arr[mid] < x {
        return binarySearchRecursive(arr, x, mid, high)
    } else {
        return binarySearchRecursive(arr, x, low, mid)
    }
}

