package util

import "fmt"

func chain(i ...int)func(j int){
	//fmt.Println(i)
	return func(j int) {
		fmt.Println(i)
		fmt.Println(j)
	}
}

func ChainCall(){
	x:=chain(4,5)
	x(3)
}
