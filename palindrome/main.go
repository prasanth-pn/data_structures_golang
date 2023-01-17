package main

import "fmt"

func main() {
	result := palindromeEasy(121)
	fmt.Println(result)

}
func palindromeHard(x int) bool {

	return false
}
func palindromeEasy(x int) bool {
	//time complexity o(n)
	//space complexity o(1)
	var reverse int

	num := x
	if x < 0 {
		return false
	}
	for num != 0 {
		rem := num % 10
		reverse = reverse*10 + rem
		num /= 10 //num=num/10
	}
	if x == reverse {
		return true
	}
	return false

}
