package time

import (
	"fmt"
	"time"
)

func TimeDemo() {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t := time.Date(2013, 3, 31, 2, 30, 0, 0, loc)
	fmt.Println(t)
}
