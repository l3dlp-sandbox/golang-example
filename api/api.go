package api

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"gitlab.mytaxi.lk/pickme/go-util/log"
)

func ApiCall(){
	response,err:=http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))



}
