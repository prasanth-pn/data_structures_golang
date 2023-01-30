package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello  brother"
	str1 := " how are you"
	cocat(str1, str)
	fmt.Println("\n\nprinting sting to ascii")
	for i, v := range str1 {
		fmt.Println(i, v)
	}
	swap(str1, str)
	strToBinary("poda thendi patti enthuvade ninakku vera paniyille")
	binary := []byte{}
	BinaryToString(binary)
	strk := replaceAlphabets("hello WoRld", 22)
	fmt.Println(strk)

}
func cocat(str1, str string) {
	str = str + str1
	fmt.Println(str)

}
func swap(str1, str string) {
	str1, str = str, str1

	fmt.Print("\n", str+str1, "\n")
}
func strToBinary(str string) {
	b := []byte(str)
	for _, v := range b {
		fmt.Printf("%08b", v)
	}

}
func BinaryToString(binary []byte) {
	str := string(binary)
	fmt.Println(str)

}
func replaceAlphabets(input string, n int) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(alphabets)

	var modified strings.Builder
	fmt.Println(modified, "modified")

	for _, char := range input {
		lowchar := strings.ToLower(string(char))
		fmt.Println(lowchar, "lowchar")
		if strings.Index(alphabets, lowchar) != -1 {
			newIdx := (strings.Index(alphabets, lowchar) + n) % 26
			modified.WriteString(string(alphabets[newIdx]))

		} else {
			modified.WriteString(string(char))
		}
	}
	return modified.String()
}
