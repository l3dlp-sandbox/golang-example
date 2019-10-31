package api

import (
	"time"
	"encoding/json"
	"gitlab.mytaxi.lk/pickme/go-util/log"

	"fmt"
	"net/http"
	"bytes"
)

func MqttApi(){
	pm:=make(map[string]interface{})
	pm["reason"]="testing"
	pm["block_status"]=true
	m:=make(map[string]interface{})
	m["driver_id"]=9647
	m["type"]="driver_blocked"
	m["created_at"]=time.Now().Unix()
	m["publish_message"]=pm

	dataJson,err:=json.Marshal(m)

	url:="http://driver.mytaxi.lk:8013/v1/driver_notification"

	if err!=nil{
		log.Error(log.WithPrefix("driver-disconnect-2.api.MqttDriverPushNotification","Fail to json encoding"),err)
		return
	}

	resp ,err:=http.Post(url,"application/json",bytes.NewBuffer(dataJson))

	if err!=nil{
		log.Error(log.WithPrefix("driver-disconnect-2.api.DriverAPIPublishUnpublish",fmt.Sprintf("Fail to send request-driverID-%vError-%v",9647,err)))
		return
	}

	//fmt.Println(resp.Status)

	if resp.Status=="200 OK"{
		fmt.Println("done")
	}


	//url:="http://driver.mytaxi.lk:8013/v1/driver_notification"
	//
	//resp ,err:=http.Post(url,"application/json",bytes.NewBuffer(dataJson))
	//
	//if err!=nil{
	//	log.Error(log.WithPrefix("driver-disconnect-2.api.DriverAPIPublishUnpublish",fmt.Sprintf("Fail to send request-driverID-%vError-%v",9647,err)))
	//	return
	//}
	//
	//if resp.Status!="200 OK"{
	//	//TODO have check real response status
	//}

}
