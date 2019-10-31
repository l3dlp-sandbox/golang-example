package http

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"io"
	"bytes"
	"encoding/json"
	"time"
	"crypto/tls"
)

func MakeRequest(){
	resp,err:=http.Get("https://httpbin.org/get")
	if err!=nil{
		log.Fatalln(err)
	}
	body,err:=ioutil.ReadAll(resp.Body)

	if err!=nil{
		log.Fatalln(err)
	}

	log.Println(string(body))
}




type httpCall struct {

}

func(m httpCall) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"go web")
}

func CreateRequest(){
	var htc httpCall
	http.ListenAndServe(":8080",htc)
}

func HelloTest(w http.ResponseWriter,req *http.Request)  {
	io.WriteString(w,"hello world")
}

func CreateRequestAnotherWay(){
	http.HandleFunc("/test",HelloTest)
	http.ListenAndServe(":8080",nil)
}

func NewRequest()  {
	resp,err:=http.NewRequest("GET","http://example.com",bytes.NewReader([]byte("fgfgf")))
	if err!=nil{
		fmt.Println(err)
	}
	b,_:=json.Marshal(resp.Body)
	fmt.Println(string(b))
}

func HttpCallWithClient()  {
	tr:=&http.Transport{
		TLSClientConfig:&tls.Config{
			InsecureSkipVerify:true },
	}
	timeout := time.Duration(5 * time.Second)
	client:=http.Client{
		Transport:tr,
		Timeout:timeout,
	}
	resp,_:=client.Get("http://example.com")
	fmt.Println(resp.Status)
}
