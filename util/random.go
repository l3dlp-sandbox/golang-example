package util

import (
	"fmt"
	"math/rand"
)

type random struct {

}

func (*random)randomGenerator()  {

	x:=rand.Perm(3)
	for y:=range x{
		fmt.Println(y)
	}

	z:=rand.NewSource(1)
	fmt.Println(z)

}
