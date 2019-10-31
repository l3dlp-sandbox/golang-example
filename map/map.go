package _map

import (
	"time"
)

type request struct{
	driverId int64 `json:"driver_id"`
	vehicleModel int `json:"vehicle_model"`
	amount float32 `json:"amount"`
	timestamp int64 `json:"timestamp"`
}
var myMap map[int64]request

type MapImplementation struct {

}
func (*MapImplementation)InitMap(){
	myMap = make(map[int64]request)
	for i:=0;i<3;i++{
		myMap[int64(i)]=request{
			driverId:int64(i),
			vehicleModel:1,
			amount:50,
			timestamp:time.Now().Unix(),
		}
	}
}
func (*MapImplementation)MaintainMap(){
	for key,value := range  myMap{
		if(time.Now().Unix() - value.timestamp)>5*60{
			delete(myMap,key)
		}
	}
}
