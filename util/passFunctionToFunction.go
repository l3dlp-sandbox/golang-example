package util

import (
	"strings"
	"fmt"
)

func CallFunction(){
	toUpperSync("Hello Callbacks", func(v string) {
		fmt.Printf("Callback: %s",v)
	})
}

func toUpperSync(iw string,f func(v string)) {
	f(strings.ToUpper(iw))
}
