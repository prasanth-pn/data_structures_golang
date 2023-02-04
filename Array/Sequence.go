package main

import "fmt"

func main() {
	array := []int{15, 1, 22, 25, 6, -1, 8, 10}
	s := []int{1, 6, -1, 10}

	ok := IsValidSubsequence(array, s)
	fmt.Println(ok)
}
func IsValidSubsequence(array []int, sequence []int) bool {

	j,i:=0,0
		for i<len(array)&&j<len(sequence){
			if array[i]==sequence[j]{
				j++
			}
			i++
		}
	   
		return j==len(sequence)
	}
	