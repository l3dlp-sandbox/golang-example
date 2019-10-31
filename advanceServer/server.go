package advanceServer

import (
	"net/http"
	"fmt"
	_"test/test/gracefull-shutdown-server"
)

func Init()  {
	http.HandleFunc("/default",handler)
	http.ListenAndServe(":8080",nil)
}
func handler(w http.ResponseWriter,r *http.Request )  {
fmt.Println("/ route")
}
