package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	data := []int{25, 15, 10, 4, 12, 22, 18, 25, 50, 35, 31, 44, 70, 66, 90}
	var result string
	for _,k:=range data{
		result+=strconv.Itoa(k)

	}
fmt.Println(result)

fmt.Println(	strings.Contains(result,"2515104"))
}
