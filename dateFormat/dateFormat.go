package dateFormat

import (
	"fmt"
	"time"
)

func DateFormat() {
	var s1 string = "2019-12-09"
	var s2 string = "09:45"
	fmt.Println(s1 + ":" + s2)
	d1, err := time.Parse("2006-01-02:15:04", s1+":"+s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d1)
}
