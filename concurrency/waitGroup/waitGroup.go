package waitGroup

import (
	"sync"
	"time"
	"fmt"
)

func WaitGroup(){
	mess := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		mess <- 1
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		mess <- 2
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		mess <- 3
	}()

	go func() {
		for i:= range mess {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}