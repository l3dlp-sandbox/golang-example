package dependency_injection

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"testing"
)

type MockHttpClient struct {
}
func(m *MockHttpClient)Get(url string)(*http.Response,error){
	response := &http.Response{
		Body : ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
	}
	return response, nil
}

func TestSendWithValidResponse(t *testing.T){
	httpClient := &MockHttpClient{}
	err:=send(httpClient, "IT_JUST_WORKS")
	if err!=nil{
		t.Errorf("received error-%s",err)
	}
}