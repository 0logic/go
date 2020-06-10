package main

import "fmt"

func main(){
	var a = 3
	fmt.Printf("a`s memory address is %x\n",&a)
	var ip *int
	if(ip==nil){
		fmt.Println("this is a null pointer\n")
	}
	fmt.Printf("export ip pointer:%p\n",ip)

	ip = &a
	fmt.Printf("ip pointer address is:%p\n",ip)
	fmt.Printf("ip pointer value is:%v",*ip)

}
