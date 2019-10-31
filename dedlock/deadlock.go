package deadlock

import "fmt"
//
var myChan chan int
func DeadlockExample(){
	myChan= make(chan int)
	go listner()
	myChan <- 1
}
func listner(){
	fmt.Println(<-myChan)
}

