package main

import "fmt"

func main() {
	data := []int{3, 2, 4, 1, 5}
	//InsertionSort(data, len(data))
	Insertion(data)
	fmt.Println(data)

}

func Insertion(a []int){
	for i:=1;i<len(a);i++{
		temp:=a[i]
		j:=i-1

		for j>=0&&temp<a[j]{
			a[j+1]=a[j]
			j--
		}
		a[j+1]=temp
	}
}


























func InsertionSort(data []int, n int) { //3,2,4,1,5

	for i := 1; i < n; i++ {
		temp := data[i] //2
		j := i - 1      //3

		for j >= 0 && data[j] > temp { //3>2 true
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp

	}

}
