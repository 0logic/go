package main

import "fmt"
//定义结构体
type Books struct{
	title string
	author string
	desc string
	page int
}

type Rect struct{   //定义矩形类
	x,y float64       //类型只包含属性，并没有方法
	width,height float64
}

func main(){
	//结构图实例化
	//initc()
	var Rect1 Rect
	Rect1.width = 5
	Rect1.height = 7
	//Area(Rect1)
}


//结构体实例化
func initc(){
	var book1 Books
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	book1.title = "Go 语言"
	//fmt.Println(Books{title: "logic",author: "wuker",desc: "logic with math"})
	fmt.Println(book1.title)


}


func (r *Rect) Area() float64{    //为Rect类型绑定Area的方法，
	return r.width*r.height         //方法归属于类型
}