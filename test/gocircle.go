package main

import "fmt"

func main(){
	//for 循环1  init; condition; post
	//for1()
	//for 循环2 只有 condition
	//for2()
	//for 循环3 空循环
	//for3()
	//乘法表
	for4()
}
//常见的带有;的循环
func for1(){
	_,b := 1,100
	for a:=1;a<b;a+=20{
		fmt.Println(a)
	}
}

//for循环只有循环条件
func for2(){
	sum :=1
	for ;sum <= 10; {
		sum += sum
	}
	fmt.Println(sum)
	sum =1
	for sum<10{
		sum +=sum
	}
	fmt.Println(sum)
}

//for循环
func for3(){
	sum :=1
	for {
		sum++
	}
	fmt.Println(sum)
}

func for4(){
	for a:=1;a<10;a++{
		for b:=1;b<=a;b++{
			if(b<a){
				fmt.Printf("%dx%d=%d ",b,a,a*b)
			}else if(a==b){
				fmt.Printf("%dx%d=%d\n",b,a,a*b)
			}
		}
	}
}