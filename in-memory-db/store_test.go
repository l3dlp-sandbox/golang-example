package inMemoryDb

import (
	"testing"
	"fmt"
	"encoding/json"
)

func Test_Memory(t *testing.T){
	database=make(map[int]Processed)
	store()
	var old []ClientPro = database[20].ClientPro
	addone:=ClientPro{
		ClientID: "250 - fseopfnskldmvdpi",
		ConnectedAt: "2018-11-20",
		Keepalive: 5,
		Node: "150@fgfg",
	}
	old= append(old, addone)
	//v2,_:=json.Marshal(old)
	//fmt.Println(string(v2))
	v:=database[20]
	v.ClientPro=old
	//v.SubscriptionPro=database[20].SubscriptionPro
	//v.SessionPro=database[20].SessionPro
	database[20]=v
	for key,val:= range database{
		fmt.Println(key)
		v,_:=json.Marshal(val)
		fmt.Println(string(v))
	}
}