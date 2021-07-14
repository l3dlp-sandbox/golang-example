package enterprices_customer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func MainEnterpricesCustomer() {
	 values :=[]string{
		 "LW101975572",
		 "LW101975575",
		 "LW101975578",
		 "LW101975581",
		 "LW101975584",
		 "LW101975587",
		 "LW101975590",
		 "LW101975593",
		 "LW101975596",
		 "LW101975599",
		 "LW101975602",
		 "LW101975605",
		 "LW101975608",
		 "LW101975611",
		 "LW101975614",
		 "LW101975617",
		 "LW101975620",
		 "LW101975623",
		 "LW101975626",
		 "LW101975629",
		 "LW101975632",
		 "LW101975635",
		 "LW101975638",
		 "LW101975641",
		 "LW101975644",
		 "LW101975647",
		 "LW101975650",
		 "LW101975653",
		 "LW101975656",
		 "LW101975659",
		 "LW101975662",
		 "LW101975665",
		 "LW101975668",
		 "LW101975671",
		 "LW101975674",
		 "LW101975677",
		 "LW101975680",
		 "LW101975683",
		 "LW101975686",
		 "LW101975689",
		 "LW101975692",
		 "LW101975695",
		 "LW101975698",
		 "LW101975701",
		 "LW101975704",
		 "LW101975707",
		 "LW101975710",
		 "LW101975713",
		 "LW101975716",
		 "LW101975719",
		 "LW101975722",
		 "LW101975725",
		 "LW101975728",
		 "LW101975731",
		 "LW101975734",
		 "LW101975737",
		 "LW101975740",
		 "LW101975743",
		 "LW101975746",
		 "LW101975749",
	 }
	 for _,v:= range values{
	 	private(v)
	 }
}
func private(v string){

	url := "http://psg-cmsapi.circles.life:80/api/1/account/customer/update/Am8zaj8wzk"
	method := "POST"
	s:=fmt.Sprintf(`{
 "accountNumber": "%s",
 "accountProfiles":{
   "strcustom1":"2815885696544929",
   "strcustom2":"-",
   "strcustom3":"dbs",
   "strcustom4":"master", 
   "strcustom5":"3683",
   "strcustom11":"-"
 }
}`,v)

	payload := strings.NewReader(s)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}