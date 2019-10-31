package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/labstack/gommon/log"
	"math/rand"
	"net"
)

func HmacServer() {
	ln, err := net.Listen("tcp", ":9192")
	if err != nil {
		log.Fatal(err)
		fmt.Println("server up and listening on port 6000")

		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go handleConnectionServer(conn)
		}
	}
}

func handleConnectionServer(c net.Conn) {
	fmt.Printf("Client %v connected", c.RemoteAddr())
	serverSideAuthenticate(c, secret)
	fmt.Printf("Connection from %v closed", c.RemoteAddr())
}

func serverSideAuthenticate(clientConn net.Conn, secretKey string) {
	message := randStr(16, "alphanum")
	_, err := clientConn.Write([]byte(message))
	if err != nil {
		clientConn.Close()
	}
	fmt.Println("Data send to client :", message)
	hasher := hmac.New(md5.New, []byte(secret))
	hasher.Write([]byte(message))
	serverHMACdigest := hasher.Sum(nil)
	fmt.Println("Server :", base64.StdEncoding.EncodeToString(serverHMACdigest))

	buffer := make([]byte, 4096)
	n, err := clientConn.Read(buffer)
	if err != nil || n == 0 {
		clientConn.Close()
		return
	}
	clientHMACdigest := buffer[:n]
	fmt.Println("Client : ", base64.StdEncoding.EncodeToString(clientHMACdigest))
	fmt.Println("Connection authenticated:", hmac.Equal(serverHMACdigest, clientHMACdigest))

	fmt.Println("continuation of process")
}
func randStr(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}
	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
