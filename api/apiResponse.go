package api

import (
	"net/http"
	"fmt"
)

func CreateRequestAnotherWay(){
	http.HandleFunc("/test",HelloTest)
	http.ListenAndServe(":8080",nil)
}
func HelloTest(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"generate")
}