package string_to_struct

import (
	"fmt"
	"encoding/json"
)

type driver struct {
	key interface{}
	value interface{}
}

var str = `[[{"driver_id":123123956,"transaction_category":4,"transaction_type":"CREDIT","description":"Kkh",
			"amount":"100.00","created_date":"2018-10-08","created_by":1471}]]`
var stri = `["trip_created_store","trip_completed_store"]`
var sq =`[dam],[day],[mal]`
 var forMap =`{"key":"382791652","value":{"id":"3dc2d113-c74e-45c1-9ec9-70e605e503c7","type":"trip_created","body":{"module":5,"booked_by":0,"trip_id":382791652,"vehicle_type":44,"pre_booking":false,"passenger":{"id":1098084},"driver":{"id":0},"corporate":{"id":0,"dep_id":0},"pickup":{"time":1541739088,"location":[{"address":"No. 150, Piyadasa Hewavitharana Building, High Level Road, Nugegoda","lat":6.86496,"lng":79.8969}]},"drop":{"location":[{"address":"71, Sunethradevi Road, Nugegoda, Sri Lanka","lat":6.863959,"lng":79.88895}]},"promotion":{"code":""},"region":{"ids":[111]},"payment":{"primary_method":1,"secondary_method":0},"comments":{"remark":"","driver_notes":""},"filters":{"driver":{"language_id":0},"vehicle":{"company_id":0,"brand_id":0,"color_id":0}},"surge":{"region_id":0,"value":0},"fare_details":{"fare_type":"","min_km":0,"min_fare":0,"additional_km_fare":0,"waiting_time_fare":0,"free_waiting_time":0,"night_fare":0,"ride_hours":0,"extra_ride_fare":0,"driver_bata":0,"trip_type":0}},"created_at":1541739088853,"expiry":0,"version":1,"trace_info":{"trace_id":{"high":0,"low":0},"span_id":0,"parent_id":0,"sampled":false}}}
`
//var forMap = `{"num":6.13,"strs":["a","b"]}`
//var forMap =`{
//	"key": "382791652",
//	"value": "damitha"
//}`
func StringToStruct(){
	//r,_:=regexp.Compile("p([a-z]+)ch")
	//r,_:=regexp.Compile("p([a-z]+)ch")
	//r,_:=regexp.Compile(`"(.*?)"`)
	//r,_:=regexp.Compile(`"(.*?)"`)
	//result:=r.FindAllString(stri ,-1)
	//var re driver
	 dat :=make(map[string]interface{})
	err:=json.Unmarshal([]byte(forMap),&dat)
	if err!=nil{
		fmt.Println(err)
	}
	v,_:=json.MarshalIndent(dat,"","")
	fmt.Println(string(v))
	//fmt.Println(re.value)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(strings.Split(`"gdgdggd"`,`"`))
}
