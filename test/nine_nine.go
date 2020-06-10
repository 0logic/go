package main

import "fmt"

func main(){
	var a,b int
	for a=1;a<=9;a++ {
		for b=1;b<=9;b++ {
			if(b>a){
				continue;
			}else{
				if(a==b){
					fmt.Printf("%d*%d=%d;\n",a,b,a*b)		
				}else{
					fmt.Printf("%d*%d=%d;",a,b,a*b)
				}
			}
		}
	}
}
