package context

import (
	"context"
	"fmt"
	"time"
)

func WithTimeOutDemo(){
	ctx,cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("done:",ctx.Err())
		}
	}()
	defer cancel()
	time.Sleep(time.Second*10)

}
