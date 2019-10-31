package server

import (
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

func StartAdvanceServer() {
	m := http.NewServeMux()

	m.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      m,
		IdleTimeout:  time.Second * 10,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Fatal(s.ListenAndServe())
}
