package main

/*
	every routine neglect when
	main routine finished
*/

import "fmt"

func main(){
	go print()
	var input string
	fmt.Scanln(&input)

}

func print(){
	for i:=0;;i++{
		fmt.Println("print",i)
	}
}