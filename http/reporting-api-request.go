package http

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sorter struct {
	Field string `json:"field"`
	DESC  bool   `json:"DESC"`
}
type BaseReportRequest struct {
	FromStaff     bool     `json:"fromStaff"`
	FromPortal    bool     `json:"fromPortal"`
	PagingEnabled bool     `json:"pagingEnabled"`
	PageSize      int      `json:"pageSize"`
	PageIndex     int      `json:"pageIndex"`
	ExportEnabled bool     `json:"exportEnabled"`
	Sorters       []Sorter `json:"sorters"`
}
type FglStakeholderBalanceSummaryReportRequest struct {
	BaseReportRequest
	StakeholderId          interface{} `json:"stakeholderId"`
	CurrencyTypeId         interface{} `json:"currencyTypeId"`
	AccountTypeId          interface{} `json:"accountTypeId"`
	AccountReferenceId     interface{} `json:"accountReferenceId"`
	FromDate               interface{} `json:"fromDate"`
	ToDate                 interface{} `json:"toDate"`
	AccountingRuleIdsIn    []string    `json:"accountingRuleIdsIn"`
	AccountingRuleIdsNotIn []string    `json:"accountingRuleIdsNotIn"`
}
type GatewayRequest struct {
	ApiVersion string `json:"apiVersion"`
	ApiAction  string `json:"apiAction"`
	MessageId  string `json:"messageId"`
	//RequestDate string      `json:"requestDate"`
	ValidateOnly bool        `json:"validateOnly"`
	RequestData  interface{} `json:"requestData"`
}

func FRA_Balance() {
	baseReportRequest := BaseReportRequest{
		ExportEnabled: true,
		FromPortal:    true,
		FromStaff:     true,
		PagingEnabled: false,
		PageSize:      10,
		PageIndex:     0,
	}
	balanceRequest := FglStakeholderBalanceSummaryReportRequest{
		BaseReportRequest:  baseReportRequest,
		StakeholderId:      "MERCHANT",
		AccountReferenceId: 37912,
	}
	gateWayRequest := BuildGateWayRequest(balanceRequest)
	rawRequest, err := json.Marshal(gateWayRequest)
	if err != nil {

	}
	fmt.Println("rawRequest:" + string(rawRequest))
	hmac := GenerateHmac("WkgvOON3FB", string(rawRequest))
	PostHttpRequest("http://146.148.110.253:8080/proxy/finance/reporting", string(rawRequest), hmac)

}
func GenerateHmac(secretKey, rawData string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(rawData))
	hashBytes := mac.Sum(nil)
	return hex.EncodeToString(hashBytes[:])
}
func BuildGateWayRequest(balanceRequest FglStakeholderBalanceSummaryReportRequest) GatewayRequest {
	request := GatewayRequest{
		ApiVersion:  "v3_7_0",
		ApiAction:   "FGL_STAKEHOLDER_BALANCE_SUMMARY",
		RequestData: balanceRequest,
	}
	return request
}
func PostHttpRequest(httpUrl, httpBody, hmac string) {
	httpRequest, err := http.NewRequest("POST", httpUrl, bytes.NewBuffer([]byte(httpBody)))
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("HMAC", hmac)
	httpClient := &http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {

	}
	defer httpResponse.Body.Close()
	rawResponse, _ := ioutil.ReadAll(httpResponse.Body)
	fmt.Println("reporting response:" + string(rawResponse))
}
