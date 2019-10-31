package main

import (
	"time"
	"fmt"
)

func main(){
	c1:=make(chan int)
	c2:=make(chan bool)
	go add(c1)
	go hold(c2)

	<-c2


	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)




}
func add(c1 chan int){
	for i:=1;i<=5 ;i++{
		c1<-i
	}
}

func hold(c2 chan bool){
	time.Sleep(time.Second*2)
	c2<-true
}