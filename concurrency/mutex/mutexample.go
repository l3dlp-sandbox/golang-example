package mutex

import (
	"sync"
	"fmt"
	"time"
)

type SafeNumber struct {
	val int
	m sync.Mutex
}
func(i *SafeNumber)Get() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}
func (i *SafeNumber)set(val int)  {
	i.m.Lock()
	i.val=val
	defer i.m.Unlock()
}

func GetNumber(){
	i:= &SafeNumber{}
	go func() {
		i.set(5)
	}()
	time.Sleep(time.Second*5)
	fmt.Println( i.Get())
}
