package regx

import (
	"regexp"
	"fmt"
)

func RegExpression(){
	r,_:=regexp.Compile("^p")
	fmt.Println(r.FindAllString("peach punch peach punch",-1))
}