package context

import (
	"context"
	"time"
	"gitlab.mytaxi.lk/pickme/go-util/traceable_context"
	"github.com/google/uuid"
	"gitlab.mytaxi.lk/pickme/go-util/log"
)

func PassLogContext(){
	ctx:=traceable_context.WithUUID(uuid.New())
	go withContext(ctx,"in zero")
	time.Sleep(time.Second*25)
	log.InfoContext(ctx,log.WithPrefix("test","testing"),"f f 5 8")
	time.Sleep(time.Second*50)
}

func withContext(c context.Context, s string){
	for{
		go secondFun(c,"in second")
		select {
		case <- c.Done():
			log.InfoContext(c,log.WithPrefix("test","first one cancelled"),"f f 5 8")
			return
		default:
			log.InfoContext(c,log.WithPrefix("test",s),"f f 5 8")
		}
		time.Sleep(time.Second*20)
	}
}

func secondFun(ctx context.Context,s string){
	for{
		select {
		case <-ctx.Done():
			log.InfoContext(ctx,log.WithPrefix("test","second one cancelled"),"f f 5 8")
			return
		default:
			log.InfoContext(ctx,log.WithPrefix("test",s),"f f 5 8")
		}
		time.Sleep(time.Second*20)
	}
}
