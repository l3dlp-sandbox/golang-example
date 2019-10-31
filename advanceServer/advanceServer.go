package advanceServer

import (
	"net/http"
	"fmt"
)

func InitAdvanceServer(){
	h:=&http.Server{
		Addr:":8080",
		ReadTimeout:10,
		WriteTimeout:10,
	}
	err:=h.ListenAndServe()
	if err!=nil{
		fmt.Println(err)
	}
}