package main

import (
	"fmt"
)

var a []int 
func main(){
	s2 := make([]int, 3, 10)
	fmt.Println(len(s2),cap(s2))	
	a := [10]int{}
	fmt.Println(a)
	s1 := a[:10]
	fmt.Println(s1)
}
