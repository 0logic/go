package main

import "fmt"

type Phone interface{
	call()
}

type NokiaPhone struct{

}

func (ssss NokiaPhone) call() {
	fmt.Println("this is nokia phone")
}

type ApplePhone struct{

}

func (tttt ApplePhone) call(){
	fmt.Println("this is apple phone")
}


func main(){
	var phone Phone
	phone = new(ApplePhone)
	phone.call()

	phone = new(NokiaPhone)
	phone.call()
}

















/*type Interface1 interface{
	call()
}

type Struct1 struct{
	
}

func (nokia Struct1) call(){
	fmt.Println("I am Nokia,I'll call you")
}

type Struct2 struct{

}

func (iphone  Struct2) call(){
	fmt.Println("I am iphone,I'll call you")
}

func main(){
	var phone Interface1
	phone = new(Struct1)
	phone.call()
	phone = new(Struct2)
	phone.call()
	
}*/

