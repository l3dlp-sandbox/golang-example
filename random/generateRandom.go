package random

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandom() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(1-1) + 1
	fmt.Println(n)
}
