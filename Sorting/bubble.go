package main
import "fmt"
func main() {
	data := [] int{23, 45, 12, 23, 22, 67, 56}
	Bubble(data, len(data))
	//gh:=*&data[1]
	fmt.Println(&data)
}
func Bubble(data []int, l int) {
	if l == 1 {
		return
	}
	for i := 0; i < l-1; i++ {
		if data[i] > data[i+1] {
			data[i+1], data[i] = data[i], data[i+1]
		}
		Bubble(data, l-1)
	}
}
