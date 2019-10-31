package middleware

import (
	"net/http"
	"log"
	"fmt"
)

type middleware func(http.HandlerFunc)http.HandlerFunc

func chainMiddleware(mw ...middleware)middleware{
	return func(final http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request){
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}
func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	log.Println("reached home")
	fmt.Fprintf(w, "welcome")
}
func Middleware(){
	mw:=chainMiddleware(withLogging,withTracing)
	http.Handle("/",mw(home))
	http.ListenAndServe(":8081",nil)
}
