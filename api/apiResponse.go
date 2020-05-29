package api

import (
	"fmt"
	"net/http"
)

func CreateRequestAnotherWay() {
	http.HandleFunc("/_test", HelloTest)
	http.ListenAndServe(":8080", nil)
}
func HelloTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "generate")
}
