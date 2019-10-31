package grace_full_shutdown

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
	"net/http"
)

var graceFulStop  =  make(chan os.Signal)
func GraceFul()  {
	fmt.Println("entered")
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w,"Server is Running")
	})

	signal.Notify(graceFulStop,syscall.SIGINT)
	signal.Notify(graceFulStop, syscall.SIGINT)

	go func() {
		sig := <-graceFulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 second to finish processing")
		time.Sleep(2*time.Second)
		os.Exit(0)
	}()

	http.ListenAndServe(":1234",nil)
}
