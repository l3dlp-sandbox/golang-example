package concurrency

import (
	"fmt"
)

func Channel(){
	msg := make(chan string)
	go func() {
		msg <- "hello"
	}()
	message := <-msg

	fmt.Println(message)

}