package Myjson

import (
	"encoding/json"
	"fmt"
)

var data = make(map[string]int)




func JsonMarshal(){
	data["id"]=1456
	data["status"]=0
	data["module"]=3

	s,_:=json.Marshal(data)
	fmt.Println(string(s))
}

func JsonUnMarshal(){
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	value:=json.Unmarshal(byt,&dat)
	fmt.Println(value)
}
