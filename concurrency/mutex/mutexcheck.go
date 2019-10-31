package mutex

import (
	"sync"
	"fmt"
	"time"
)

var storageReference *storage

type storage struct {
	State map[int]int
	Mutex sync.Mutex
}

func MutexExample1(){
	storageReference = &storage{
		State: make(map[int]int),
		Mutex: sync.Mutex{},
	}
	go func() {
		for  i := 0; i<1000;i++{
			storageReference.Mutex.Lock()
			storageReference.State[i]=i
			storageReference.Mutex.Unlock()
		}
	}()

	go func() {
		for i:=0; i<1000;i++{
			storageReference.Mutex.Lock()
			fmt.Println(storageReference.State[i])
			storageReference.Mutex.Unlock()

		}
	}()
	time.Sleep(time.Second*10)
}
