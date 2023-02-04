package main

import "fmt"

func main() {
	data:=[]int{4,12,46,78,34}
	QuickSort(data)
	
}
func QuickSort(a []int){
	n:=len(a)-1
	if n==1{
		return
	}
	pivot:=a[n]

	i:=0
	for j:=0;j<n;j++{
		if pivot>a[j]{
       a[i],a[j]=a[j],a[i]
	   i++
		}

	}

	a[i],a[n]=a[n],a[i]
	//fmt.Println(a)
	QuickSort(a[:i])
	QuickSort(a[i:])
}