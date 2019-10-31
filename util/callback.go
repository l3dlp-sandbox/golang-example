package util

import (
	"fmt"
	"time"
)



func testCallBack (){
	y:=func(s string){
		time.Sleep(time.Second*15)
		fmt.Println(s)
	}
	CallBack(y,"tsting")
}

func CallBack(f func(s string),s string)  {
	f(s)
	fmt.Println("call-back return")
}
