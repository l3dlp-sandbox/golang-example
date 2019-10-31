package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

func CustomRequest() {
	moreAdvancedClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
		},
		Timeout: time.Hour * 1,
	}
	body := map[string]string{
		"name": "morpheus",
		"job":  "leader",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		//handle error
		return
	}
	req, err := http.NewRequest("post", "https://e6e15284-6e36-4c81-bde8-0f3415bac9e1.mock.pstmn.io", bytes.NewBuffer(jsonBody))
	if err != nil {
		//handle error
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
