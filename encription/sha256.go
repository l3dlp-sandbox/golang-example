package encription

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha256(){
	hash:=sha256.Sum256([]byte("damitha"))
	fmt.Println(hex.EncodeToString(hash[:]))
}
