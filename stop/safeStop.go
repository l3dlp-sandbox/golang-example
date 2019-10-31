package stop

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

var gracefulstop = make(chan os.Signal)

func TestSafeExit()  {
	signal.Notify(gracefulstop,syscall.SIGTERM)
	signal.Notify(gracefulstop,syscall.SIGINT)
	go func() {
		<-gracefulstop
		time.Sleep(time.Second*5)
		fmt.Println("got signal to exit")
		os.Exit(0)
	}()

	for{
		fmt.Println("testing........")
	}
}
