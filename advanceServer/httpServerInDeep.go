package advanceServer

import (
	"net/http"
	"io"
)

func hello(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello World")
}

