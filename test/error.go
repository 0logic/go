package main

import "fmt"

//定义结构体
type DevideError struct{
	devidee int
	devider int
}

//定义错误接口
func (de *DevideError) Error() string{
	strFmt := `
	Cannot precess the devider is zero
	devidee %d
	devider 0
`
	return fmt.Sprintf(strFmt,de.devidee)
}
//定义除法运算
func Devide(varDevidee int,varDevider int)(result int,err string){
	if(varDevider==0){
		dData := DevideError{
			devidee: varDevidee,
			devider: varDevider,
		}
		err = dData.Error()
		return
	}else{
		return varDevidee/varDevider,""
	}
}


func main(){
	//正常情况
	na,nb := 100,20
	resn,ok := Devide(na,nb)
	if (len(ok)==0) {
		fmt.Printf("%d divide %d is %d\n",na,nb,resn)
	}
	//异常情况
	sa,sb := 10,0
	_,oks := Devide(sa,sb)
	if(len(oks)>0){
		fmt.Printf("%d divide %d error is %s",sa,sb,oks)
	}
}