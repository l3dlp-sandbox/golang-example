package syncPackage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
)

//memory pool - a group of memory block that are freed under programmer
//control.
//with this approach programmer can reduce the pressure on it's garbage
//collector and end of the day performance improvement
//sync.Pool allows Go programmers to allocate and free memory manually

type small struct {
	a int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(small)
	},
}

func inc(s *small) {
	s.a++
}

func BenchmarkWithoutPool(b *testing.B) {
	var s *small
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			s = &small{
				a: 1,
			}
			inc(s)
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			s = pool.Get().(*small)
			s.a = 1
			inc(s)
			pool.Put(s)
		}
	}
}

type requst struct {
	Data string `json:data`
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := requst{}
	buf := bufPool.Get().(*bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&data)
	if err != nil {

	}
	fmt.Println(data)
	bufPool.Put(buf)
}
