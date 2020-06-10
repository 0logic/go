package main

import "fmt"

func main(){
	var n [10] int 
	var i,j int
	for i=0;i<10;i++ {
		n[i] = i+100
	}
	for j=0;j<10;j++ {
		fmt.Printf("Element[%d] = %d\n",j,n[j])
	}
	
	var avg float32
	var blance = [] int {100,23,444,23,1}
        avg = getAverage(blance,5)
	println()
	fmt.Printf("blance average is:%f",avg)

}

func getAverage(arr [] int,size int) float32 {
	var avf float32
	var i,sum int
	for i=0;i<size;i++{
		sum+=arr[i]
	}
	//avf = float32(sum/size)
	avf = float32(sum)/float32(size)
	return avf
}

