package main

import "fmt"

func main() {
	var array []int

	array = insert(0, 3, array)
	fmt.Println(array)
	pos := 1
	dat := 23
	update(pos, dat, array)
	i := 0
	n := len(array)
	fmt.Println(array)
	fmt.Println(" raversing array")
	display(i, n, array)
	var after []int
	after = delete(i, 78, n, array, after)
	after = nil
	after = delete(i, 23, n, array, after)
	fmt.Println(after, "this is the last value")

}

func delete(i, data int, n int, array []int, after []int) []int {
	for i := 0; i < n; i++ {
		if array[i] != data {
			after = append(after, array[i])
		}
	}
	return after
}

// -------------------------------------display
func display(i, n int, array []int) {
	if i == n {
		return
	}
	fmt.Println(array[i])
	display(i+1, n, array)
}
func update(pos int, data int, array []int) {
	array[pos] = data
}

func insert(i, size int, array []int) []int {
	var data int

	for i := 0; i < size; i++ {
		fmt.Scan(&data)
		array = append(array, data)
	}
	return array

}
