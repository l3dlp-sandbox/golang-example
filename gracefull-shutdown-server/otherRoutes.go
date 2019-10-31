package gracefull_shutdown_server

import (
	"net/http"
	"fmt"
)

func init(){
	http.HandleFunc("/myProfile",myProfileHandler)
	fmt.Println("default profile called")
}
func myProfileHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("my profile route")
}
