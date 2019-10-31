package DesignPattern

import "fmt"

type user struct{
	name string
	age int
}

func(u user)itself(){
	fmt.Println(u.name)
}

func Builder(){
	u:=user{}
	u.name="damitha"
	u.age=27
	u.itself()
	newBuilder()
}
func newBuilder()  {
	u:=user{}
	fmt.Println(u.age)
}