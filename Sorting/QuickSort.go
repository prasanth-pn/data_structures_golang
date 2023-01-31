package main

import "fmt"

func main() {
	//data := []int{34, 23, 123, 678, 4, 34, 55}
	data := []int{17, 28, 3, 7, 5, 642, 24}
	//n:=len(data)-1
	QuickSort(data)
	fmt.Println(data)

}

func QuickSort(a []int) []int {//17,3,7,5
	var n = len(a) - 1
	if n < 1 {
		return a
	}

	pivot := a[n]
	//fmt.Println(pivot)
	i := 0
	for j := 0; j < n; j++ {
		if pivot > a[j] {
			//fmt.Println(a[j],pivot,a[i])
			a[i], a[j] = a[j], a[i]
			i++
			//fmt.Println(i,j)
		}
	}

	a[i], a[n] = a[n], a[i]
	QuickSort(a[:i])	
	QuickSort(a[i+1:])
	return a

}
