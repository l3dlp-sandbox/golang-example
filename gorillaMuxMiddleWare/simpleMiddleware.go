package gorillaMuxMiddleWare

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/prometheus/common/log"
)

func MiddleWareMux(){
	//s:=mux.NewRouter()
	//s.HandleFunc("/signUp",signUp)


	r:=mux.NewRouter()
	r.HandleFunc("/",simpleMw(handler))
	r.HandleFunc("/signup",signUp)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter,req *http.Request) {
		w.Write([]byte("method not found"))
	})
	//r.Use(simpleMw)

	log.Fatal(http.ListenAndServe("localhost:8080",r))
}
func signUp(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("signed"))
}

func handler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("working"))
}

func simpleMw(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request) {
		if r.Header.Get("Authorization")!="password"{
			log.Info("Authorization failed")
			http.Error(w,"Forbidden",http.StatusForbidden)
			return
		}
		log.Info(r.Header.Get("Authorization"))
		next(w,r)
	}
}