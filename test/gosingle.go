package main

import "fmt"

//go运算符
func main(){
    //算数运算符 + - * / % ++ --
    //math()
    //关系运算符 == >= <= > < != ==
    //relation()
    //逻辑运算符 && || !
    //logic()
    //位运算符 & | ^ << >>
    //binary()
    //其他运算符 & *
    another()
}

func another(){
    a,b := 3,5
    c := &a
    fmt.Printf("%d & is %d\n",a,c)
    var d = &b
    e := *d
    fmt.Printf("%d * is %d\n",d,e)


}
//位运算符
func binary(){
    a,b := 4,2
    var c = a&b
    fmt.Printf("%d & %d is %d\n",a,b,c)
    c = a|b
    fmt.Printf("%d | %d is %d\n",a,b,c)
    c = a^b
    fmt.Printf("%d ^ %d is %d\n",a,b,c)
    c = a<<b
    fmt.Printf("%d << %d is %d\n",a,b,c)
    c = a>>b
    fmt.Printf("%d >> %d is %d\n",a,b,c)
}

//逻辑运算符
func logic(){
    a,b := 4,-1
    if(a>0 && b>0){
        fmt.Println("a and b is unsign")
    }else if(a>0 || b>0){
        fmt.Println("a or b is unsign")
    }

    if(!(a>0 && b>0)){
        fmt.Println("a or b is unsign")
    }
}

//关系运算符
func relation(){
    a,b := 3,4
    //var c = '3'
    //==  同一类型才可比较
    if(a==b){
        fmt.Println("a=c")
    }else{
        fmt.Println("a!=c")
    }
    //>=
    if(a>=b){
        fmt.Println("a>=b")
    }else{
        fmt.Println("a<b")
    }
}

//算数运算符
func math(){
    a,b := 3,5
    c := a+b
    fmt.Printf("first c=%d\n",c)
    c = b-a
    fmt.Printf("second c=%d\n",c)
    c = b*a
    fmt.Printf("third c=%d\n",c)
    c = b/a
    fmt.Printf("forth c=%d\n",c)
    c = b%a
    fmt.Printf("fifth c=%d\n",c)
    b++
    fmt.Printf("sixth b=%d\n",b)
    b--
    fmt.Printf("seventh b=%d\n",b)
}