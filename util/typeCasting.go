package util

import "fmt"

type user struct {
	id int
	fName string
	sName string
}

var i interface{}

func check()  {
	testVariable:=user{
		id:123,
		fName:"Damitha",
		sName:"Dayananda",
	}

	i=testVariable
	castedVariable:=i.(user)

	fmt.Println(castedVariable.id)
}

