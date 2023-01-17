package main

import "fmt"

func main() {
	data := []int{ 10, 6, 4, 3}
	target := 10

	//fmt.Println(twosum(data, target))
	fmt.Println("--------------------")
	fmt.Println(TwoSum(data, target))
}

func twosum(data []int, target int) []int {
	//iterating both slices
	//checking for the target
	//o(n2) bruteforce solution
	for i, left := range data {
		for k, right := range data {
			if left+right == target && i != k {
				return []int{i, k}
			}
		}
	}
	return nil
}
func TwoSum(data []int, target int) []int {
	hm := make(map[int]int)
	//fmt.Println(hm)
	for i, num := range data {
		_, ok := hm[num]
			fmt.Println(ok, hm[num])
		if ok {
			return []int{i, hm[num]}
		}
		fmt.Println(i, "value of i")
		//fmt.Println(hm,hm[target-num], "value of hm")
		fmt.Println(hm)
		hm[target-num] = i
		fmt.Println(hm)

	}
	return nil

}
