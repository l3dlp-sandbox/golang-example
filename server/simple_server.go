package server

import (
	"fmt"
	"net/http"
)

func StartBasicServer() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe("8080", nil)
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
