package mapJsonToStruct

import (
	"encoding/json"
	"fmt"
)

var jsonString = "{\"age\":15,\"reportIndex\":{\"name\":\"finance_report_fgl_reporting_log*\",\"type\":\"fgl_reporting_logs\"}}"

type reportIndex struct {
	Name string
	Type string
}

type fDoc struct {
	Age         int
	ReportIndex reportIndex
}

func MapJsonToStruct() {
	fdoc := fDoc{}
	err := json.Unmarshal([]byte(jsonString), &fdoc)
	if err != nil {
		fmt.Println("Error" + err.Error())
	}
	fmt.Println(fdoc.Age)
}
