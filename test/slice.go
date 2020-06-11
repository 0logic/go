package main

import "fmt"

func main(){

	var slice1 [2] int
	fmt.Println(slice1)

	var slice2 = make([]int, 4,5)
	fmt.Println(slice2,len(slice2),cap(slice2))
	//slice := [] int {1,2,4,5,6}
	//fmt.Println(slice)
	/*var slice = make([]int,3,5)
	fmt.Println(slice)

	var arr1 = [] int {1,2,4,5,6,7,8,60}
	fmt.Println(arr1)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))	
	var arr2 [3] int
	arr2[0] = 3
	fmt.Println(arr2)*/


	/*fmt.Printf("slice length %d\n",len(slice))
	fmt.Printf("arr1 length %d\n",len(arr1))
	fmt.Printf("arr1 cap %d\n",cap(arr1))
	

	var slice1 [] int
	if(slice1==nil){
		fmt.Printf("slice1 is null Slice ,length %d\n",len(slice1))
	}

	var slice2 [] int=make([]int,2,5)
	fmt.Printf("slice2 is not a nil Slice,length  %d\n",len(slice2))
	
	slice3 :=arr1[1:5]
	fmt.Println(slice3)
	

	slice4 := make([]int,len(slice3)*2,cap(slice3)*2)
	copy(slice4,slice3)

	fmt.Println(slice4)
	
	slice4 = append(slice4,2,4,5)


	var slice5[2]int
	//slice5 = append(slice5,2)
	fmt.Println(slice5)
	fmt.Printf("slice5 length %d\n",len(slice5))

	fmt.Println(slice4)*/
}
