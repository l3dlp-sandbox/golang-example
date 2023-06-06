package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//_,err:=NtbToken.GetNtbToken()
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//server_mux.ServerMuxDemo()
	//p:=pointers.PointerDemo{}
	//p.PointerDemo()
	//p.PointerDemoNew()
	//p.PointerToPoinDemo()
	//advanceServer.Init()
	//DesignPattern.Builder()
	//profilling.Server()
	//p:=make(map[string]float64)
	//p["method"]=1
	//p["amount"]=123.12
	//q:=make(map[string]float64)
	//q["method"]=3
	//q["amount"]=150.52
	//pa:=[]map[string]float64{}
	//pa=append(pa, p,q)
	//r,_:=json.MarshalIndent(pa,"","")
	//fmt.Println(string(r))

	//config_load.ConfigLoad()
	//ba:=encription.GetNewCiperBlock("")
	//hba:=encription.HexEncription(ba)
	//fmt.Println(string(hba))
	//dba,_:=encription.DeHexString("67656e65726174656369706589a176a970ec80321e202")

	//970ea89371c2c36b272548246e6d57a3b640d80e2589a
	//dba,_:=encription.DeHexString("67656e65726174656369706589a17fa970e9833319264e120c8c9b41bbc187e2bd80bf761fe6")
	////dba,_:=encription.DeHexString("67656e65726174656369706589a176a970ea89371c2c36b272548246e6d57a3b640d80e2589a")
	//ds:=encription.Decrypt(dba,"password")
	//fmt.Println(string(ds))
	//memoryUsage.GoMemoryUsage()
	//log.Println("log _test")
	//clog.Info("log _test")
	////encription.EncodeStringInHexaDecimal()
	//encription.CreateHash()
	//encription.GetNewCiperBlock(encription.CreateHash())
	//deadlock.DeadlockExample()
	// millionRequestfunc.MyImplementationMain()
	//database.InsertQuery()
	//v:=strconv.FormatInt(111112222,10)
	//fmt.Println)
	//Method_chaning.MethodChainWithErrorExample()
	//millionRequestfunc.StartServer()
	//database.Init()
	//database.SelectFromPassengers()
	//start:=time.Now()
	//database.RunParalelWrite()
	//fmt.Println(time.Since(start))
	//file.ReadFile()
	//gorillaMuxMiddleWare.SimpleServer()
	//gorillaMuxMiddleWare.MiddleWareMux()
	//jwt.JwtMain()
	//uuid.GetUuidNewV4()
	//regx.RegExpression()
	//grace_full_shutdown.GraceFul()
	//util.ChainCall()
	//database.GoSelect()
	//middleware.Middleware()
	//go_cron.GoCron()
	//string_to_struct.StringToStruct()
	//time.Sleep(time.Second*80)
	//mutex.MutexExample2()
	//pass_by_value_reference.StartPass()
	//x:=make(map[int]string)
	//x[10]="dd"
	//x[20]="aa"
	////for k:= range x {
	////	fmt.Println(k)
	////}
	//v,ok:=x[10]
	//if !ok{
	//	fmt.Println("err")
	//}else {
	//	fmt.Println(v)
	//}
	//mutex.MutexExample1()
	//context.PassLogContext()
	//uuid.GetUuid()
	//middleware.AuthenticateMiddleware()
	//database.Init()
	//result:=database.Db.QueryRow("SELECT taxi_modelid FROM passengers_log_archive WHERE passengers_log_id=?",33062)
	//
	//var num int
	//err := result.Scan(&num)
	//if err != nil{
	//	resultSecond:=database.Db.QueryRow("SELECT taxi_modelid FROM passengers_log WHERE passengers_log_id=?",33062)
	//	fmt.Println("passenger log archive"+err.Error())
	//	err2:= resultSecond.Scan(&num)
	//	if err2 != nil{
	//		fmt.Println("passenger logs"+err2.Error())
	//	}
	//}
	//fmt.Println(num)
	//k_stream.KStreamMain()
	//dependency_injection.Dependency()
	//x:=util.Round{
	//
	//}
	//x.Value=123.3246
	//x.RoundFunc()
	//api.CreateRequestAnotherWay()
	//fmt.Println((time.Now().UnixNano()/100000))
	//
	//database.Init(
	//	//database.WithReadConfig(&mysql.Config{
	//	//	ReadTimeout:  100 * time.Millisecond,
	//	//	Timeout:      100 * time.Millisecond,
	//	//	WriteTimeout: 100 * time.Millisecond,
	//	//}),
	//	database.WithWriteConfig(&mysql.Config{
	//		ReadTimeout:  boostrap.Utils.ReadTimeOut * time.Millisecond,
	//		Timeout:      30 * time.Millisecond,
	//		WriteTimeout: boostrap.Utils.WriteTimeOut * time.Millisecond,
	//	}),
	//)
	//
	//defer database.Close(database.Connections.Read)
	//defer database.Close(database.Connections.Write)
	//
	//for  {
	//	fmt.Println(boostrap.Utils.ConnectionTimeOut)
	//	//start:=time.Now()
	//	//database.Connections.Write.Ping()
	//	//fmt.Println(time.Since(start))
	//}

	//stop.TestSafeExit()
	//context.RollbackFunction()
	//x:=56.32
	//s:=fmt.Sprintf("%.2f",x)
	//v,_:=strconv.ParseFloat(s,64)
	//fmt.Println(v)
	//
	//addRounded := strconv.FormatFloat(x,'f',2,64)
	//fmt.Println(addRounded)

	//c:=math.Round(45.12)
	//x:=int64(c)
	//y:=time.Unix(1533168020 ,0)
	///fmt.Println(y.Format("2006-01-02"))
	//fmt.Println(c)
	//time.Parse("2006-01-02","2006-01-02")
	//	//transaction.Transaction()//for testing transaction
	//	//context.Handler()
	//
	//	producer.NewProducer()//Test produce produce message to driver ledger
	//
	//	//api.ApiCall()
	//
	//	//Myjson.JsonMarshal()
	//	//Myjson.JsonUnMarshal()
	//
	//	//log.Info(log.WithPrefix("_test","working"))
	//
	//	//v:="12345"
	//	//i,err:=strconv.Atoi(v)
	//	//if err!=nil{
	//	//
	//	//}
	//	//fmt.Println(i)
	//	//api.MqttApi() //api testing mqtt
	//
	//t := time.Unix(1494505756, 0)
	//	//y:=t.Format("2006-01-02 15:04:05")
	//	//fmt.Println(y)
	//
	//	//context.Start()
	//	//log.Error(log.WithPrefix("","Testing"))
	//
	//
	//	//database.TestConnectionOff()
	//	//time.Now().Format( "15:04:05")
	//
	//	//server_mux.ServerMux()
	//
	//	////****************promitheus*****************

	//
	//database.Init()
	//defer database.Close(database.Connections.Read)

	//	context.RollbackFunction()
	////
	//for {
	//	promitheus.ServerInit()
	//	time.Sleep(time.Millisecond*250)
	//}
	//	fmt.Println("here")
	//	for{
	//		go func() {
	//			promitheus.Example()
	//		}()
	//
	//		time.Sleep(time.Millisecond*50)
	//
	//	}
	//	////*******************************************
	//

	//x:=consumer.KconsumerLedger{}
	//x.InitializerLedger()
	//x.ConsumerLedger()

	//}

	//func main()  {
	//	v:=SqlError{
	//	}
	//	v.DriverId=9644
	//	v.Status="Y"
	//	v.RunQuery()
	//}
	//
	//type SqlError struct {
	//	Status string
	//	DriverId int
	//
	//}
	//
	//func (s *SqlError)RunQuery()(){
	//
	//	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/driver_disconnect")
	//	if err!=nil{
	//		fmt.Println("error connection")
	//		panic(err.Error())
	//
	//	}
	//
	//	db.Exec("UPDATE driver SET login_status=?  WHERE id=?")
	//
	//
	//database.UpdateQuery()
	//gracefull_shutdown_server.MainStartHttpServer()

	//go map
	//m:=_map.MapImplementation{}
	//m.InitMap()
	//m.MaintainMap()

	//random.GenerateRandom()
	//Myjson.ComplexJson()
	//http.CustomRequest()
	//protobuff.ProtoDemo()
	//mapJsonToStruct.MapJsonToStruct()
	//dateFormat.DateFormat()
	//deffer.DefferDemo()
	//xmlToCsv.XmlToCsv()

	//typeConversion.TypeConversionDemo()
	//urlParsing.UrlParsingDemo()
	//syncPackage.BenchMarkWithoutPool()
	//database.AccountIdTest()
	//time.TimeDemo()
	//pointerReceiverMethodAndIteration.Demo()
	//pointerReceiverMethodAndIteration.DeferWithRange()
	//jwt.JwtMain()
	//context.WithTimeOutDemo()
	//enterprices_customer.MainEnterpricesCustomer()
	//sm.StructMapDemo()
	//type gmoStatus struct {
	//	ShopId []string `json:"ShopID"`
	//	ShopPass	[]string `json:"ShopPass"`
	//	AccessID	string `json:"AccessID"`
	//	AccessPass string `json:"AccessPass"`
	//	OrderID string `json:"OrderID"`
	//	Status string `json:"Status"`
	//	JobCd	string `json:"JobCd"`
	//	Amount  string `json:"Amount"`
	//	Tax     float32 `json:"Tax"`
	//	Currency string `json:"Currency"`
	//	Forward string `json:"Forward"`
	//	PayTimes string `json:"PayTimes"`
	//	TranID  int64 `json:"TranID"`
	//	Approve int64 `json:"Approve"`
	//	TranDate string `json:"TranDate"`
	//	ErrCode string `json:"ErrCode"`
	//	ErrInfo string `json:"ErrInfo"`
	//	PayType string `json:"PayType"`
	//}
	//http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
	//	err:=request.ParseForm()
	//	if err !=nil{
	//		fmt.Println(err)
	//	}
	//	fmt.Println(request.Form)
	//	var gmo gmoStatus
	//	b,_:=json.Marshal(request.Form)
	//	err=json.Unmarshal(b,&gmo)
	//	if err !=nil{
	//		fmt.Println(err)
	//	}
	//	fmt.Println("marshaled:",gmo)
	//})
	//http.ListenAndServe(":8080",nil)
	//s:=fmt.Sprintf()
	//var  x *int= nil
	//fmt.Println(*x)

	//for i:=0;i<10;i++{
	//	fmt.Println(rand.Intn(2))
	//}
	//OneMillionPerformance.SplitFiles()
	//for i:=0;i<100;i++{
	//	load()
	//}
	//load()
	//loadPostForm()
	//encription.Sha256()
	//encription.Md5()
	//web.EmbeddedFileServerMain()wai
	//req, err := http.NewRequest(http.MethodGet, "http://scl-mobkd.cxos.tech/v4/cl/en/mobile/settings/option/list/get?app_type=ecosystem&option_id=autoboost_limit%2Csettings_show_autoboost_limit%2Cidentity_info&uuid=<no value>", nil)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//headers := map[string][]string{
	//	"Child-sin":{""},
	//	"Content-Type":{"application/json"},
	//	"Service-Id":{"Quilt"},
	//	"User-Agent":{"selfcare/1.0.0-B2B-RC2-qa_7d6c5bb6 Android/11 Nokia 3.4/en_GB"},
	//	"User-Id":{"JP-CGJV9GDCVEMY"},
	//	"X-AUTH":{"eyJhbGciOiJSUzI1NiIsImtleV9pZCI6IjYzYTRmNWQ2YTYyZjZhOTZlMjlmYTE3N2VkZDJmMWM1IiwidHlwIjoiSldUIn0.eyJkZXZpY2VfdHlwZSI6Ik1vYmlsZSIsImV4cGlyeV90aW1lIjoxNjQ3OTQxMDc4LCJleHRlcm5hbF9pZCI6IkpQLUNHSlY5R0RDVkVNWSIsImlzcyI6ImNpcmNsZXMiLCJsb2dpbl90aW1lIjoxNjQ3OTQwMTc4LCJ0eXAiOjV9.O2GF5WXH_HAzINc_OTOHyFoVReRBuwtkp7MXJj54EeqFMcZA6-YPeTI_CRPUPkpGjvFD-wk3fv_cgBlss6Ba18tCY5oGkie3lc9ld1AyUd3B_0Rl7zc0mMgx9cIXWkIh9X3Z-THl0y91hmEnIUydxddW3mDK0B0q7BBXo8H2ss-l1kya-Pej1gvhfD9G1UvmICDXoec8nd0JoY0oN0pCr2dypGwBaGzrKYLk6CXzTrOPcl2_4507PVWcInulU361HoObn8AGD20TTrMbGYl1Y7YUlSWCnWofA4yGL_TuZFBldwzjYeI1VAE6r6dqr-d5QeKAsxxJVanDFmE-a225rw"},
	//	"X-Deviceid":{"ffffffff-bbc7-d7ca-0000-017fb070e7bacom.circles.selfcare.xx.qa"},
	//	"X-Newrelic-Id":{"VQ4BUlNSARABVFVVDwQBV1IB"},
	//	"X-Newrelic-Transaction": {"PxQCBAJWAANSU1hVU1NVBVcFFB8EBw8RVU4aV11cAQoDBw5YBgVRAAcFA0NKQQ1SBwJZBFEBFTs="},
	//	"X-Request-Id":{"b7b98e56-1d9b-4f55-81a3-2d653d38ef90"},
	//	"X-Source":{"QuiLt"},
	//	"x-app-version":{""},
	//}
	//req.Header = headers
	//
	//res, reqErr := http.DefaultClient.Do(req)
	//if reqErr != nil{
	//	fmt.Println(reqErr)
	//}
	//b,_:=ioutil.ReadAll(res.Body)
	//fmt.Println(string(b))

	//url := "http://scl-mobkd.cxos.tech/v4/cl/en/mobile/settings/option/list/get?app_type=ecosystem&option_id=autoboost_limit%252Csettings_show_autoboost_limit%252Cidentity_info&uuid=%3Cno%20value%3E"
	//method := "GET"
	//
	//client := &http.Client {
	//}
	//req, err := http.NewRequest(method, url, nil)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("Child-sin", "")
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Service-Id", "Quilt")
	//req.Header.Add("User-Agent", "selfcare/1.0.0-B2B-RC2-qa_7d6c5bb6 Android/11 Nokia 3.4/en_GB")
	//req.Header.Add("User-Id", "JP-CGJV9GDCVEMY")
	//req.Header.Add("X-AUTH", "eyJhbGciOiJSUzI1NiIsImtleV9pZCI6IjYzYTRmNWQ2YTYyZjZhOTZlMjlmYTE3N2VkZDJmMWM1IiwidHlwIjoiSldUIn0.eyJkZXZpY2VfdHlwZSI6Ik1vYmlsZSIsImV4cGlyeV90aW1lIjoxNjQ3OTQxMDc4LCJleHRlcm5hbF9pZCI6IkpQLUNHSlY5R0RDVkVNWSIsImlzcyI6ImNpcmNsZXMiLCJsb2dpbl90aW1lIjoxNjQ3OTQwMTc4LCJ0eXAiOjV9.O2GF5WXH_HAzINc_OTOHyFoVReRBuwtkp7MXJj54EeqFMcZA6-YPeTI_CRPUPkpGjvFD-wk3fv_cgBlss6Ba18tCY5oGkie3lc9ld1AyUd3B_0Rl7zc0mMgx9cIXWkIh9X3Z-THl0y91hmEnIUydxddW3mDK0B0q7BBXo8H2ss-l1kya-Pej1gvhfD9G1UvmICDXoec8nd0JoY0oN0pCr2dypGwBaGzrKYLk6CXzTrOPcl2_4507PVWcInulU361HoObn8AGD20TTrMbGYl1Y7YUlSWCnWofA4yGL_TuZFBldwzjYeI1VAE6r6dqr-d5QeKAsxxJVanDFmE-a225rw")
	//req.Header.Add("X-Deviceid", "ffffffff-bbc7-d7ca-0000-017fb070e7bacom.circles.selfcare.xx.qa")
	//req.Header.Add("X-Newrelic-Id", "VQ4BUlNSARABVFVVDwQBV1IB")
	//req.Header.Add("X-Newrelic-Transaction", "PxQCBAJWAANSU1hVU1NVBVcFFB8EBw8RVU4aV11cAQoDBw5YBgVRAAcFA0NKQQ1SBwJZBFEBFTs=")
	//req.Header.Add("X-Request-Id", "b7b98e56-1d9b-4f55-81a3-2d653d38ef90")
	//req.Header.Add("X-Source", "QuiLt")
	//req.Header.Add("x-app-version", "")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
	// You can edit this code!
	// Click here and start typing.
	//gracefull_shutdown_server.MainStartHttpServer()
	//oop.CreateBird()
	//cgo.Cmain()
	//str := "{\"alphabet\":{\"id\":\"A\"\"name\":\"dam\"}}"
	//c := ABC{}
	//err := json.Unmarshal([]byte(str), &c)
	//fmt.Println(err)
	//fmt.Println(c)

	//validateHash("55b62c36b409b0b7348730a0af8721aa2dd8fc7b0e304d7ae97036f6987ae036", "CXT-LHI9PVYAAN", "YF_XKbGlXC5GKShg2WW8LLmg", "17297", "000")
	http.HandleFunc("/v4/cl/en-US/dev_test/internal/call-back/payment", handler)
	log.Fatal(http.ListenAndServe(":3002", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
	return // Send "Hello, World!" as the response
}

func validateHash(validationHash, basketId, merchantSecretKey, merchantId, errCode string) bool {
	h := sha256.New()
	composedString := fmt.Sprintf("%s|%s|%s|%s", basketId, merchantSecretKey, merchantId, errCode)
	h.Write([]byte(composedString))
	hash := h.Sum(nil)
	encoded := fmt.Sprintf("%x", hash)
	fmt.Println(encoded)
	return encoded == validationHash
}
