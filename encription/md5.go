package encription

import (
	"crypto/md5"
	"fmt"
)

func Md5(){
	byteArray:=md5.Sum([]byte("damitha"))
	fmt.Println(string(byteArray[:]))
}