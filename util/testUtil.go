package util

import (
	"encoding/json"
	"fmt"
)

type DisconnectResponse struct {
	ImpossibleBlockCount int
	PossibleBlockCount struct{
		Tuk int
		Van int
		Car int
		Nano int
		Mini int
	}

}
func Test(){
	resp:=DisconnectResponse{
		ImpossibleBlockCount:3,
		PossibleBlockCount: struct {
			Tuk  int
			Van  int
			Car  int
			Nano int
			Mini int
		}{Tuk:5 , Van: 5, Car:8 , Nano:8 , Mini: 9},
	}
	msg,_:=json.Marshal(resp)

	fmt.Println(string(msg))
}


