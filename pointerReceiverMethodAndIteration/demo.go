package pointerReceiverMethodAndIteration

import (
	"fmt"
	"sync"
)

func Demo() {
	var wg sync.WaitGroup
	for _, v := range []T{{"foo"}, {"bar"}, {"baz"}} {
		wg.Add(1)
		go v.M(&wg)
	}
	wg.Wait()
}

type T struct {
	name string
}

func (t *T) M(wg *sync.WaitGroup) {
	fmt.Println(&t.name)
	fmt.Println(t.name)
	wg.Done()
}

func DeferWithRange() {
	for _, v := range []string{"foo", "bar", "baz"} {
		defer func() {
			fmt.Println(v)
		}()
	}
}
