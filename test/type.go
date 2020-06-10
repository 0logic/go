package main
import "fmt"
//全局变量
var (
    a int
    b bool
)
func main(){
    /******三种变量申明方式******/
    //申明变量及类型
    var inta int
    inta = 1
    //申明变量，不指定类型，会自动获取类型
    var intb=3
    //通过:=赋值  其中:=是声明变量语句如果已经声明会报错，只要含有未声明的就会通过
    var intd int
    intc,intd := 4,5
    //空白变量
    _,inte := 7,9

    fmt.Printf("%s %s %s %s",inta,intb,intc,intd)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Print(inte)
}