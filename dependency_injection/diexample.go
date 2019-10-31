package dependency_injection

import (
	"net/http"
	"fmt"
	"flag"
	"errors"
	"encoding/json"
	"io/ioutil"
)

var(
	url string
)

type HttpClient interface {
	Get(string) (*http.Response,error)//HttpClient has this function any type implement this method
									  //also became type of HttpClient
}

func init()  {
	flag.StringVar(&url,"url","http://google.com","which url u wanna parse")
	flag.Parse()
}

func Dependency() error {
	client:=&http.Client{}
	return send(client,url)
}

func send(client HttpClient,link string)error{
	//client:=&http.Client{}
	response,err:=client.Get(link)
	if err!=nil{
		msg,_:=json.Marshal(customError{
			msg:"Couldn't read body",
			id:100,
		})
		return errors.New(string(msg))
	}
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println(err.Error())

	}else {
		fmt.Println(string(body))

	}
	return nil
}

type customError struct {
	msg string
	id int
}


