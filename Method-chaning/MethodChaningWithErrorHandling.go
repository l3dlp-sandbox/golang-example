package Method_chaning

import (
	"fmt"
	"errors"
)

type Chain struct {
	err error
}

func(v *Chain) funA() *Chain{
	if v.err!=nil{
		return v
	}
	fmt.Println("A")
	return v
}
func (v *Chain) funB() *Chain {
	if v.err != nil {
		return v
	}
	v.err = errors.New("error at funB")
	fmt.Println("B")
	return v
}
func (v *Chain) funC() *Chain {
	if v.err != nil {
		return v
	}
	fmt.Println("C")
	return v
}

func MethodChainWithErrorExample(){
	c:=Chain{}
	d:=c.funA().funB().funC()
	fmt.Println(d.err)
}