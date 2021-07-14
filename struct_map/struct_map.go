package struct_map

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
)
/*
This library can be used to create a map from struct
 */

func StructMapDemo(){
	//s:= `{"token": ["09bd3d41c1072a3e8a015e5a724947e59d1e90e2e556a734a9d2ead0bbf5d37b", "63445e9816f48904c7de085855338dec0d4c949d223caf897872de82d90c7219"], "accessID": "", "customer": {"last_name": "", "first_name": "", "user_email": "", "contact_number": ""}, "card_type": "Visa", "user_name": "", "accessPass": "", "cancel_url": "", "return_url": "", "user_email": "sos4e@help.com", "description": "20GBプラン -  #000001623749448596", "failure_url": "", "order_lines": null, "payment_terms": "", "valid_duration": 120, "billing_address": {"city": "", "name": "", "state": "", "line_1": "", "post_code": "", "country_code": ""}, "delivery_method": "", "payment_methods": ["CARD"], "referral_method": "", "first_six_digits": "411111", "last_four_digits": "1111", "shipping_address": {"city": "", "name": "", "state": "", "line_1": "", "post_code": "", "country_code": ""}, "referrer_user_name": "", "failure_redirect_url": "https://qjp-webfrontek.circles.life/web/payment-status?order_ref=000001623749448596&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmRlcl9yZWYiOiIwMDAwMDE2MjM3NDk0NDg1OTYiLCJ0aW1lc3RhbXAiOiIyMDIxLTA2LTE1VDE4OjMwOjQ5LjUzMSswOTowMCIsIm90cF92ZXJpZmllZCI6dHJ1ZSwic2NvcGUiOlsiYWxsIl19.6kfOEklv4iJlEqoJqpYNnPX-CsVxLmBFBuwOSWgwod4", "success_redirect_url": "https://qjp-webfrontek.circles.life/web/payment-status?order_ref=000001623749448596&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmRlcl9yZWYiOiIwMDAwMDE2MjM3NDk0NDg1OTYiLCJ0aW1lc3RhbXAiOiIyMDIxLTA2LTE1VDE4OjMwOjQ5LjUzMSswOTowMCIsIm90cF92ZXJpZmllZCI6dHJ1ZSwic2NvcGUiOlsiYWxsIl19.6kfOEklv4iJlEqoJqpYNnPX-CsVxLmBFBuwOSWgwod4", "custom_payment_method": "card", "secondary_phone_number": "", "payment_config_category": "TEST", "amazon_order_reference_id": ""}`
	s:=`{}`
	var _map map[string]interface{}
	json.Unmarshal([]byte(s),&_map)
	fmt.Println(_map)
}

func strucToMap(){
	s:= struct {
		Name string
		Age int
	}{Name: "Damitha",Age: 12}
	fmt.Println(structs.Map(s))
}
