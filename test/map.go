package main 
import "fmt"

func main(){
	//var map1 map[string]string
	map1 := make(map[string]string)
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
	fmt.Println(map1)

}
