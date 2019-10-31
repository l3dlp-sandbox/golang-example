package context

import (
	"context"
	"fmt"
	"time"
)

func Handler(){
	ctx,cancel:=context.WithCancel(context.Background())

	go doStuff(ctx)
	for i:=0;i<10;i++{
		fmt.Println(i)
		time.Sleep(time.Second*5)
		if i==5{
			cancel()
		}
	}

}
func doStuff(ctx context.Context)  {

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		return
	}
}
