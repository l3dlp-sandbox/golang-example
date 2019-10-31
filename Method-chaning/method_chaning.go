package Method_chaning

import (
	"fmt"
	"github.com/pkg/errors"
)

type human interface {
	wakeup() human
	doWork() human
	gotoBed()human
}
type person struct {
	firstName string
	Error error
}

func(p person) wakeup() human{
	fmt.Printf("%s  got up, ",p.firstName)
	return p
}
func(p person) doWork() human{
	abError:= errors.New("new error")
	return person{
		"damitha",
		abError,
	}
	fmt.Printf("%s done work, ",p.firstName)
	return p
}
func(p person) gotoBed() human{
	fmt.Printf("%s went to bed, ",p.firstName)
	return p
}

func Method_Chaning()  {
	p:=person{
		"damitha",
		nil,
	}
	p.wakeup().doWork().gotoBed()
}