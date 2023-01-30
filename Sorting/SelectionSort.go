package main

import "fmt"

func main() {

	data := []int{23, 12, 34, 56, 3, 1, 344, 678, 123}
	n:=len(data)
	fmt.Println(n)
	SelectionSort(data, n)
	fmt.Println(data)
}
func SelectionSort(a []int,n int){



for i:=0;i<n-1;i++{//i=0 valueis 23
	min:=i  //i=0
	for j:=i+1;j<n;j++{ //j=1 traversing to end till nth position 
		if a[j]<a[min]{ // checking the value 12<23,true stores the j to min (1)then checks 34<23 false then 56<23 false
			//3<23 true value to min is 4 a[min]is 3; 1<3 true min=5 344<1 false, 678<1 false 123<1 false
			min=j
		}
	}
	if min!=i{// now the minimum value is not equal to the i then swap
	 temp:=a[min]
	 a[min]=a[i]
	 a[i]=temp
	}
}
}
