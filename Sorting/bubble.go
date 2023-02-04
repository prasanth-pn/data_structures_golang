package main

import "fmt"

func main() {
	fmt.Println("enter the ")
	data := []int{23, 45, 12, 23, 22, 67, 56}
	Bubble(data)
	//gh:=*&data[1]
	fmt.Println(data)

}
func Bubble(a []int){
	n:=len(a)
	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
}






























// func Bubble(data []int, l int) {
// 	if l == 1 {
// 		return
// 	}
// 	for i := 0; i < l-1; i++ {
// 		if data[i] > data[i+1] {
// 			data[i+1], data[i] = data[i], data[i+1]
// 		}
// 		fmt.Println(data)
// 		Bubble(data, l-1)
// 	}
// }



















