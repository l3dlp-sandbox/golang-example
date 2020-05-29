package profilling

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/*
go tool pprof -http localhost:6061 http://localhost:6060/debug/pprof/heap
go tool pprof --seconds 5 http://localhost:8040/debug/pprof/profile



*/
var numArray []int

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Profilling(num int) {
	fmt.Println(Fib(num))
}

//
func Server() {
	http.HandleFunc("/_test", testHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	go increasingMemory()
	w.Write([]byte("adding success"))
}
func increasingMemory() {
	for {
		numArray = append(numArray, rand.Intn(100))
		for i := range numArray {
			fmt.Println(i)
		}
		fmt.Println("****************")
		time.Sleep(time.Millisecond * 100)
	}
}
