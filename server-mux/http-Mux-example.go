package server_mux

import (
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func(th *timeHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	tm:=time.Now().Format(th.format)
	w.Write([]byte("The time is"+tm))
}

func ServerMuxDemo(){
	mux:=http.NewServeMux()
	th:=&timeHandler{
		format:time.RFC1123,
	}
	mux.Handle("/time",th)
	http.ListenAndServe(":3000",mux)
}
