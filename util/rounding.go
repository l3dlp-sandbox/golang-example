package util

import (
	"fmt"
	"strconv"
)

type Round struct{
	Value float64
}
func (r* Round) RoundFunc() {
	x:=fmt.Sprintf("%.2f",r.Value)
	value,err:=strconv.ParseFloat(x,64)
	if err!=nil{
		fmt.Println("unable to convert")
		return
	}
	fmt.Println(value)
}
