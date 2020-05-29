package context

import (
	"context"
	"github.com/google/uuid"
	"gitlab.mytaxi.lk/pickme/go-util/log"
	"gitlab.mytaxi.lk/pickme/go-util/traceable_context"
	"time"
)

func PassLogContext() {
	ctx := traceable_context.WithUUID(uuid.New())
	go withContext(ctx, "in zero")
	time.Sleep(time.Second * 25)
	log.InfoContext(ctx, log.WithPrefix("_test", "testing"), "f f 5 8")
	time.Sleep(time.Second * 50)
}

func withContext(c context.Context, s string) {
	for {
		go secondFun(c, "in second")
		select {
		case <-c.Done():
			log.InfoContext(c, log.WithPrefix("_test", "first one cancelled"), "f f 5 8")
			return
		default:
			log.InfoContext(c, log.WithPrefix("_test", s), "f f 5 8")
		}
		time.Sleep(time.Second * 20)
	}
}

func secondFun(ctx context.Context, s string) {
	for {
		select {
		case <-ctx.Done():
			log.InfoContext(ctx, log.WithPrefix("_test", "second one cancelled"), "f f 5 8")
			return
		default:
			log.InfoContext(ctx, log.WithPrefix("_test", s), "f f 5 8")
		}
		time.Sleep(time.Second * 20)
	}
}
