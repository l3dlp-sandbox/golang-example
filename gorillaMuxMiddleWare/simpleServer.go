package gorillaMuxMiddleWare

import "net/http"


type TestHandler struct {

}
func(T TestHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
}

func(T TestHandler) TestHandlerFunc(path string, f func(w http.ResponseWriter, r *http.Request)){
	http.HandleFunc(path,f)
}

func SimpleServer(){
	s:=TestHandler{}
	http.ListenAndServe(":8080", s)
	s.TestHandlerFunc("/",handlerSimple)
}
func handlerSimple(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("custom handler working"))
}
