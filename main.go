package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "net/http/pprof"
	"test/test/mapJsonToStruct"
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
	///////dba,_:=encription.DeHexString("67656e65726174656369706589a176a970ec80321e202")

	//970ea89371c2c36b272548246e6d57a3b640d80e2589a
	//dba,_:=encription.DeHexString("970ea89371c2c36b272548246e6d57a3b640d80e2589a")
	//dba,_:=encription.DeHexString("67656e65726174656369706589a176a970ea89371c2c36b272548246e6d57a3b640d80e2589a")
	//ds:=encription.Decrypt(dba,"password")
	//fmt.Println(string(ds))
	//memoryUsage.GoMemoryUsage()
	//log.Println("log test")
	//clog.Info("log test")
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
	//	//log.Info(log.WithPrefix("test","working"))
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
	mapJsonToStruct.MapJsonToStruct()
}
