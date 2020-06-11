package main

import "fmt"

func main(){
	//数组定义声明
	inits()
}

//数组声明
func inits(){
	var arr1 [2] int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1)

	var arr2 = []int{2,4,5,6,7,8,435,121}
	fmt.Println(arr2)

	fmt.Println(len(arr2))
}

