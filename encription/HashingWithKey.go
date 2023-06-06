package encription

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"github.com/prometheus/common/log"
	"strconv"
	"strings"
	"io"
)

func CreateHash()string{
	hasher:=md5.New()
	hasher.Write([]byte("password"))
	hashedString:= hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(hashedString)
	return hashedString
	//5f4dcc3b5aa765d61d8327deb882cf99
}
func GetNewCiperBlock(key string)[]byte{
	data:=[]byte("dayananda")
	block,_:=aes.NewCipher([]byte(CreateHash()))
	gcm,_:=cipher.NewGCM(block)
	nonceSize:=gcm.NonceSize()
	fmt.Println("nonceSize : "+strconv.FormatInt(int64(nonceSize),10))
	nonce:=make([]byte,nonceSize)
	newReader:=strings.NewReader("generateciper")
	io.ReadFull(newReader,nonce)
	ciperText:=gcm.Seal(nonce,nonce,data,nil)
	fmt.Println(ciperText)
	return ciperText
	//[103 101 110 101 114 97 116 101 99 0 0 0 23 3 105 78 145 96 243 58 183 137 177 97 27 125 39 188 145 221 123 69 220 220 139 150 95]
	//[0 0 0 0 0 0 0 0 0 0 0 0 15 138 40 205 41 20 30 223 181 198 111 207 66 71 224 17 23 99 87 179 32 204 224 238 126]
}
func HexEncription(enc []byte) string{
	return hex.EncodeToString(enc)
}
func DeHexString(str string) ([]byte, error) {
	return hex.DecodeString(str)
}
func Decrypt(data []byte, password string) []byte {
	key := []byte(CreateHash())
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR getting cipher: ", err)
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Info("ERROR getting gcm: ", err)
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Info("ERROR getting plaintext: ", err)
		panic(err.Error())
	}
	return plaintext
}