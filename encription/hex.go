package encription

import (
	"encoding/hex"
	"fmt"
)

func EncodeStringInHexaDecimal(){
	src:=[]byte("damitha")
	encString:=hex.EncodeToString(src)
	fmt.Println(encString)
}