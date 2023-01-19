package main

import "fmt"


type Student struct{
	name string
	grade []int
	age   uint
}

func (s *Student)getAge()([]int,uint){
s.age=12

	return s.grade,s.age
}

func (s Student)getname()string{
	df:=fmt.Sprint( s.age)
	return  df
}
func (s Student)average()float32{
	count:=0
	 var sum float32 
	for _,v:=range s.grade{
		count++
		sum:=+v
	}
	average:=sum/float32(count)
	return average

}


func main() {

	s1:=Student{"prasanth",[]int{67,56,45},34}
	fmt.Println(s1)
fmt.Println(s1.getAge())
fmt.Println(s1.age)
	
}