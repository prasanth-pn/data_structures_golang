package main

import (
	"fmt"
)

func main() {
	data := []int{25, 15, 10, 4, 12, 22, 18, 24, 50, 35, 31, 44, 70, 66, 90}

	 sum := 0
	 i:=0
	// num := 5
	// sum = fact(num, sum)
	// fmt.Println(sum, "nksjdkfafjk")
	total(data,i,sum)

}
func total(data []int,i,sum int){
	if len(data)==i{
		return
	}
	sum+=data[i]
	fmt.Println(sum,i,data[i])
	total(data,i+1,sum)
	
}
func fact(num int, sum int) int {

	if num == 1 {
		fmt.Println(sum, "return ")
		return sum

	}
	//fmt.Println(num)
	sum = sum * num

	fmt.Println(sum)
	return fact(num-1, sum)
}
