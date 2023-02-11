package main

import "fmt"

type maxheap struct {
	arr []int
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return (2 * i) + 1
}
func right(i int) int {
	return (2 * i) + 2
}
func swap(arr []int,a,b int){
	arr[a],arr[b]=arr[b],arr[a]
}

func main() {
	data := []int{45, 67, 23, 12, 66, 89, 3, 54}
	heap := maxheap{}
	heap.buildheap(data)
	fmt.Println(data)

}
func (h *maxheap) buildheap(arr []int) {
	if len(arr)-1<=0{
		return
	}
	h.arr = arr
	arr = h.arr
	for i := parent(len(arr) - 1); i >= 0; i-- {
		heapify(arr,i)
	}
	swap(arr,0,len(arr)-1)
	h.buildheap(arr[:len(arr)-1])
}
func heapify(arr[]int, i int) {
		left:=left(i)
		right :=right(i)
		large :=i
		n:=len(arr)-1
		if left<=n&&arr[left]>arr[large]{
			large =left
		}
		if right<=n&&arr[right]>arr[large]{
			large=right
		}
		if large!=i{
			swap(arr,large,i)
			heapify(arr,large)
			heapify(arr[:n-1],large)

		}
}
