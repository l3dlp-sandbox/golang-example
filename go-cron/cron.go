package go_cron

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
)

func GoCron() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { // at 10-49  "0 49 10 * * *"
		//fmt.Println("every 30 seconds")
		execute()
	})
	c.Start()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	for {
		select {
		case msg := <-signals:
			log.Info("message", msg)
			c.Stop()
		}
	}
}

func execute() {
	fmt.Println("every  seconds")
}
