package context

import (
	"fmt"
	"context"
	"time"
)

func Print(){

	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Millisecond)
	defer cancel()
	select {
		case <-ctx.Done():
			fmt.Println("ctx done")
		case <-test():
			fmt.Println("function return")
		case <-time.After(time.Millisecond*500):
			fmt.Println("done")
	}
}

func test() <-chan bool{
	c:=make(chan bool)
	time.Sleep(time.Second*10)
	fmt.Println("here")
	close(c)

	return c
	//select {
	//case <-time.After(5*time.Second):return
	//case <-ctc.Done():return
	//}
}