package main

import "fmt"

type person struct{
	love int
	money float64
	house int
	feel string
}

func main(){
	var dxh person
	//var wuker person
	dxh.love = 2
	dxh.money = 8
	dxh.house = 1
	dxh.feel = "sad"
	wuker := &dxh
	fmt.Println(dxh)
	fmt.Println(dxh.love)
	fmt.Printf("dxh place is:%x",wuker)
}





/*type student struct{
	name string
	score float32
	age int
	sex string
}

func main(){
	fmt.Println(student{"wuker",80.0,26,"male"})
	fmt.Println(student{name:"dxh",score:29.5,age:12,sex:"female"})
	fmt.Println(student{name:"wife",score:68.4})
}*/
