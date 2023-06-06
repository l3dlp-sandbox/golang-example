package oop

import "fmt"

type animal struct {
	Name string
}

type bird struct {
	Feather int
	animal
}

func CreateBird() {
	b := bird{}
	b.animal.Name = "hern"
	b.Name = "craw"
	fmt.Println(b)
}
