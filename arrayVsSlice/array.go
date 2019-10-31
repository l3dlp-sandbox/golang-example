package arrayVsSlice

import "fmt"

func ArrayDemo() {
	intArray1 := [...]int{1, 2, 3}
	fmt.Println(intArray1)
}
func SliceDemo() {
	intSlice := make([]int, 10)
	x := append(intSlice, 1, 2, 2)
	for i, num := range x {
		fmt.Printf("index:", i)
		fmt.Println("number:", num)
	}
}
