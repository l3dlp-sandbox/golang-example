package memoryUsage

import (
	"time"
	"runtime"
	"fmt"
)

func GoMemoryUsage(){
	for i:=0;i<4;i++{
		time.Sleep(time.Second)
		runtime.GC()
		PrintMemoryUsage()
	}
}
func PrintMemoryUsage(){
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\tAlloc = %v MiB",m.Alloc/(1064*1064))
	fmt.Printf("\t TotalAlloc = %v MiB",m.TotalAlloc)
	fmt.Printf("\t Sys = %v MiB",m.Sys)
	fmt.Println()
}
