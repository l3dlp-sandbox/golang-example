package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net"
)

var secret = "mysecret"

func HmacClient() {
	for {
		dialConn, err := net.Dial("tcp", "localhost:9192")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("connection established local and server")
		fmt.Printf("Remote Address : %s \n", dialConn.RemoteAddr().String())
		fmt.Printf("Local Address : %s \n", dialConn.LocalAddr().String())

		go handleConnection(dialConn)
	}
}
func handleConnection(c net.Conn) {
	buffer := make([]byte, 4096)
	for {
		n, err := c.Read(buffer)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		msg := string(buffer[:n])
		fmt.Println("\n Data received from server : ", msg)
		clientSideAuthenticate(c, secret, msg)
	}
}
func clientSideAuthenticate(serverConn net.Conn, secretKey string, message string) {
	hasher := hmac.New(md5.New, []byte(secretKey))
	hasher.Write([]byte(message))
	clientHMACdigest := hasher.Sum(nil)
	fmt.Println("Digest send to server:", base64.StdEncoding.EncodeToString(clientHMACdigest))
	n, err := serverConn.Write(clientHMACdigest)
	if err != nil || n == 0 {
		serverConn.Close()
		return
	}
}
