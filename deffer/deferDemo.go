package deffer

import "fmt"

func DefferDemo() {
	fmt.Println(foo())
	x := "mex"
	if errorCall(); x == "me" {
		fmt.Println("working")
	}
}
func foo() (result string) {
	defer func() {
		result = "Change World" // change value at the very last moment
	}()
	return "Hello World"
}
func errorCall() bool {
	return true
}
