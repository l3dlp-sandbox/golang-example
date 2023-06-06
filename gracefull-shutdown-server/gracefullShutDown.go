package gracefull_shutdown_server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func startHttpServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware(http.HandlerFunc(defaultRoute)))
	mux.HandleFunc("/second", secondRoute)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}

	}()
	return srv
}
func middleware(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1")
		next.ServeHTTP(w, r)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered")
			}
		}()
	}
}
func defaultRoute(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	fmt.Println("2")
	panic("panicking")
	w.Write([]byte("it's working"))
}
func secondRoute(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 35)
	fmt.Println("response written from second route")
	w.Write([]byte("second is working"))
}
func MainStartHttpServer() {
	srv := startHttpServer()
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	select {
	case <-stop:
		fmt.Println("server going to shut down")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}
}
