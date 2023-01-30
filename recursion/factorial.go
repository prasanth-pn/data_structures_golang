package main

import "fmt"

func main() {
	n:=6
	data:=Factorialrecursion(n)
fmt.Println(data)	

}

func Factorialrecursion(num int)int {
	if num==1||num==0{
		return 1
	}
		
	return	num*Factorialrecursion(num - 1)   //recursion used here for iteration 

}
