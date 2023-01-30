package main

import "fmt"

func main() {
	data := []int{34, 23, 45, 23, 1, 45, 67, 456, 3}
	InsertionSort(data, len(data))
	fmt.Println(data)
}

func InsertionSort(data []int, n int) {

	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j] > temp {
			data[j+1] = data[j]
			j--
		} 
		data[j+1]=temp 

	}

}
