package Myjson

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Data map[string]interface{} `json:"Data"`
}

type myData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ComplexJson() {
	mymap := make(map[string]interface{})
	mymap["one"] = myData{
		Name: "Damitha",
		Age:  27,
	}
	mymap["two"] = myData{
		Name: "jaliya",
		Age:  31,
	}
	myData := Data{
		Data: mymap,
	}
	d, err := json.Marshal(myData)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))
}
