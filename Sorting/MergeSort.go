package main

import "fmt"

func main() {
	array := []int{23, 12, 56, 45, 22, 90, 11, 34, 14}
	array = MergeSort(array)
	fmt.Println(array)
}

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	fmt.Println("hhhhhhhhhhhhhhhhhhhh")
	return Merge(left, right)
}
func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	fmt.Println("jjjjjjjjjjjjjjjjj", left, right)

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	fmt.Println("kkkkkkkkkkkkkkkkkkkkkkk")
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// func MergeSort(array[]int)[]int{
// 	if len(array)<=1{
// 		return array
// 	}
// 	middle:=len(array)/2
// 	left:=	MergeSort(array[:middle])

// 	//fmt.Println(array[:middle],"  ",middle)
// 	right:=MergeSort(array[middle:])
// 	//fmt.Println(array[middle:])

// 	return Merge(left,right)
// }

// func Merge(l,r []int)[]int{
// 	fmt.Println(l,r)
// 	left:=l
// 	right:=r

// 	result:=make([]int,0,len(left)+len(right))
// 	for len(left)>0||len(right)>0{
// 		if len(left)==0{
// 			fmt.Println(right,"left value")
// 			return append(result,right...)
// 		}
// 		if len(right)==0{
// 			return append(result,left...)

// 		}
// 		if left[0]<=right[0]{
// 			result=append(result, left[0])
// 			left=left[1:]
// 		}else{
// 			result=append(result, right[0])
// 			right=right[1:]
// 		}
// 	}
// 	return result

// }
