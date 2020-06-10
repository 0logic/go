package main
import "fmt"
func main(){
    var a int = 20
    fmt.Printf("const palce :%x",&a)
    println()
    var ip *int
    
    ip = &a
    println(ip)
    fp := *ip
    println(fp)
}
