package context

import (
	"context"
	"time"
	"fmt"
)

func Start(){
	ch:=make(chan bool,1)
	ctx,cancel:=context.WithTimeout(context.Background(),time.Duration(time.Second*2))
	defer cancel()
	go infinite(ctx,ch)
	select {
	case <-ch:
		fmt.Println("returned from channel")
	case <-ctx.Done():
		fmt.Println("context done")

	}
}

func infinite(ctx context.Context,ch chan bool){
	for{
		fmt.Println("working")
		time.Sleep(time.Millisecond*100)
	}
	ch<-true
}
