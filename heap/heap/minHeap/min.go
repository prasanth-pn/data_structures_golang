package main

import "fmt"

type MinHeap struct {
	arr []int
}

func left(i int)int{
	return (2*i)+1
}
func right(i int)int{
	return (2*i)+2
}
func parent (i int)int{
	return (i-1)/2
}
func swap(arr []int,a,b int){
	arr[a],arr[b]=arr[b],arr[a]
}
func (h *MinHeap) minvalue(i int)int{
	arr:=h.arr
	left:=left(i)
	right:=right(i)

	if arr[left]<arr[right]{
		return left

	}else if arr[left]>arr[right]{
		return right
	}
	return left
}

func main() {
	heap:=MinHeap{}
var n int
data:=[]int{45,23,12,567,343,21,12,1,789,23}
fmt.Println(data)
for n<8 {
	fmt.Println("\n enter the option \n1.insert\n2.delete\n 3.peek\n4.for heapify\n--------------------")
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Println("Enter the value to heap ")
		k:=0
		fmt.Scan(&k)
		heap.insert(k)
	case 2:
		heap.rootDelete()
	case 3:
		fmt.Println("The value in heap")
		heap.peek()
	case 4:
		heap.buildheap(data)

	}
}

}
func (h *MinHeap)buildheap(arr []int){
	h.arr=arr
	for i:=parent(len(arr)-1);i>=0;i--{
		h.heapify(i)
	}
}
func (h *MinHeap)heapify(i int){
	left:=left(i)
	right:=right(i)
	n:=len(h.arr)-1
	large :=i
	heap:=h.arr
	if left<=n&&heap[left]<heap[large]{
		large=left
	}
	if right<=n&&heap[right]<heap[large]{
		large=right
	}
	if large!=i{
		swap(heap,large,i)
		h.heapify(large)
	}


}
func (h *MinHeap)insert(k int){
	h.arr=append(h.arr, k)
	h.insertHelper(len(h.arr)-1)


}
func (h *MinHeap)insertHelper(position int){
	if h.arr[position]<h.arr[parent(position)]{
		swap(h.arr,position,parent(position))
		h.insertHelper(parent(position))
	}
	
}
func (h *MinHeap)peek(){
	for i:=0;i<=len(h.arr)-1;i++{
		fmt.Println(h.arr[i])
	}
}
func (h *MinHeap)rootDelete(){
	h.arr[0]=h.arr[len(h.arr)-1]
	h.arr=h.arr[:len(h.arr)-1]
	h.rootdeleteHelper(0)
}
func (h *MinHeap)rootdeleteHelper(i int){
	if left(i)>=len(h.arr){
		return 
	}
	arr:=h.arr
	min:=h.minvalue(i)

	if arr[i]>arr[min]{
		swap(arr,i,min)
		h.rootdeleteHelper(min)

	}

}
