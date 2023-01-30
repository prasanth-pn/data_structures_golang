package main

import "fmt"

func main() {
	array := []int{23, 12, 56, 45, 22, 90, 11, 34, 14}
	array=MergeSort(array)
	fmt.Println(array)
}

func MergeSort(array[]int)[]int{
	if len(array)<=1{
		return array
	}
	middle:=len(array)/2
	left:=	MergeSort(array[:middle])

	//fmt.Println(array[:middle],"  ",middle)
	right:=MergeSort(array[middle:])
	//fmt.Println(array[middle:])
	
	return Merge(left,right)
}

func Merge(l,r []int)[]int{
	fmt.Println(l,r)
	left:=l
	right:=r
	
	result:=make([]int,0,len(left)+len(right))
	for len(left)>0||len(right)>0{
		if len(left)==0{
			fmt.Println(right,"left value")
			return append(result,right...)
		}
		if len(right)==0{
			return append(result,left...)

		}
		if left[0]<=right[0]{
			result=append(result, left[0])
			left=left[1:]
		}else{
			result=append(result, right[0])
			right=right[1:]
		}
	}
	return result

}
