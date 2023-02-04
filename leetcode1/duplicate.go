package main

import "fmt"

func main() {

	data:=[]int{34,34,56,78,23,12,23,34,45,56,56,34}
	k:=removeDuplicates(data)
	fmt.Println(k,data)
	
}
func removeDuplicates(nums []int) int {
    nextInsertionIndex := 1
    
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[nextInsertionIndex] = nums[i]
            nextInsertionIndex++
        }
    }
    
    return nextInsertionIndex
}




// func removeDuplicates(nums []int) int{
// 	keys:=make(map[int] bool)
// 	list:=[]int{}

// 	for _,value:=range nums{
// 		if _, v:=keys[value]; !v{
// 			keys[value]=true
// 			list=append(list,value)

// 		}


// 	}
// 	n:=len(list)
// 	fmt.Println(list)
	
// 	   return n
// 	}