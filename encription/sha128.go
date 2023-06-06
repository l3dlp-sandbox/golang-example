package encription

import (
	"crypto/sha1"
	"fmt"
)

func Sha128()  {
	h:=sha1.New()
	h.Write([]byte("abc"))
	bs:=h.Sum(nil)
	fmt.Println(bs)
}

func Sha1MyCode(){
	d:= new(digest)
	d.h[0] = 0x67452301
	d.h[1] = 0xEFCDAB89
	d.h[2] = 0x98BADCFE
	d.h[3] = 0x10325476
	d.h[4] = 0xC3D2E1F0
	d.nx   = 0
	d.len  = 0

	d.write([]byte("abc"))
}
func (d *digest)write(p []byte){


}

type digest struct {
	h   [5]uint32
	x   [64]byte
	nx  int
	len uint64
}
