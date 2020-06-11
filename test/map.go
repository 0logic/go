package main

import "fmt"


func main(){

	var atype int
	atype = 1
	fmt.Println(float64(atype))
	//var map1 map[string]string
	/*map1 := make(map[string]string)
	map1["france"] = "bary"
	map1["china"] = "beijing"
	map1["austrilina"] = "sini"
	//fmt.Println(map1)
	capital,ok:=map1["American"]
	if(ok){
		fmt.Println("American's captical is",capital)
	}else{
		fmt.Println("American's captical is nil")
	}

	delete(map1,"china")
	fmt.Println(map1)*/
	map1 := make(map[string]string)
	map1["riben"] = "dongjing";
	map1["china"] = "beijing";
	map1["eluosi"] = "mosicao";

	a,ok := map1["amarical"]
	if(ok){
		fmt.Printf(map1["amarical"])
	}else{
		fmt.Printf("amarical is undefined")
	}
	fmt.Println(a)
	fmt.Println(map1["amarical"])

	for x,d :=range map1{
		fmt.Printf("%s captical is %s\n",x,d)
	}
	delete(map1,"riben")

	for x,d :=range map1{
		fmt.Printf("%s captical is %s\n",x,d)
	}
}
