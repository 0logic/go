package main

import "fmt"

func main(){
	/********const常量*********/
	const ha int = 4
	const wa =  5

	const aa = ha*wa
	fmt.Println(aa)
	/*********iota特殊常量
		可以理解被编辑器修改的特殊常量
		在const出现时被重置0，const每增加一行常量iota计数一次
	*********/
	const (
		ia = iota
		ib = iota
		ic = iota
	)
	fmt.Println(ia)
	fmt.Println(ib)
	fmt.Println(ic)
}

