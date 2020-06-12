package main

import "fmt"



func main(){
	//初始化
	//inints()
	//close channel
	closeChannel()
}


func closeChannel(){
	c := make(chan int,10)
	go channelClose(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func channelClose(n int,c chan int){
	x,y := 0,1
	for i:=0;i<n;i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}


func inints(){
	var arr1 = []int{1,4,5,9,-5,2}
	c := make(chan int,2)
	go sum(arr1[:3],c)
	go sum(arr1[3:],c)
	sumb := <-c
	//suma,sumb := <-c,<-c
	fmt.Println(sumb)
	//fmt.Println(suma,sumb,suma+sumb)
}

func sum(a [] int,c chan int){
	fmt.Println(a)
	sum := 0
	for _,v := range a{
		sum +=v
	}
	c <- sum
}