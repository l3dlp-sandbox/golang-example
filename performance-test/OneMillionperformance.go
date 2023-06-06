package OneMillionPerformance

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//mamath or praveen

const OrderRefStart = 100000
var OrderRefEnd int64 = 0
const BrowserCmd = "x-www-browser"
const HtmlLocation = "/home/damitha/go/src/golang-example/performance-test/login-gmo.html"
const PaymentUrl = "https://qjp-payment.circles.life/v3/JP/en-jp/oms/payment/upfront"
const MySqlConnectionString = "qjp_payment:nY434PttCnUaqvTA@tcp(qjp-rds.circles.life:3306)/qjp_payment"
const AccountFileDir = "/home/damitha/go/src/golang-example/performance-test/1m/"
const ResultFileLocation = "/home/damitha/go/src/golang-example/performance-test"
const BillingAccountFileLocation = "/home/damitha/go/src/golang-example/performance-test"


var signal chan bool
var token chan []string
var index int32 = 0

type account struct {
	BillingAccountNumber string
	CustomerAccountNumber string
	ServiceInstanceNumber string
}

var accounts []account

var logFile *os.File

func PerformanceTestMain(){
	start:=time.Now()
	OrderRefEnd = allLineCount()
	var err error
	logFile, err = os.OpenFile(BillingAccountFileLocation+"/result_billing.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	signal = make(chan bool,10)
	token = make(chan []string, 10)
	dbInit()

	go startServer()
	go browserCall()

	s:= findStartFile()
	for i:=s;i<s+fileCount();i++{
		time.Sleep(time.Second*30)
		loadAccounts(i)
		go func() {
			time.Sleep(time.Second * 2)
			signal <- true
			signal <- true
			signal <- true
			signal <- true
		}()
		var wg sync.WaitGroup
		go worker(1, signal, token,&wg)
		wg.Add(1)
		go worker(2, signal, token,&wg)
		wg.Add(1)
		go worker(3, signal, token,&wg)
		wg.Add(1)
		go worker(4, signal, token,&wg)
		wg.Add(1)
		go worker(5, signal, token,&wg)
		wg.Add(1)
		wg.Wait()
		f, err := os.OpenFile(ResultFileLocation+"/result_file.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(fmt.Sprintf("%v done \n",i)); err != nil {
			log.Println(err)
		}
	}
	fmt.Println("Elapsed time:",time.Since(start))
}

func logBillingAccount(ban string,txnId string){
	m:=sync.Mutex{}
	m.Lock()
	if _, err := logFile.WriteString(fmt.Sprintf("%v,%v\n",ban,txnId)); err != nil {
		log.Println(err)
	}
	m.Unlock()
}
func loadAccounts(i int){
	accounts = []account{}
	index = 0
	//r,e:=os.Open(fmt.Sprintf("/home/damitha/Desktop/circle documents/1mFolder/%v",i))
	r,e:=os.Open(AccountFileDir+fmt.Sprintf("%v",i))
	if e !=nil{
		fmt.Println(r)
	}
	scanner:=bufio.NewScanner(r)
	for scanner.Scan(){
		arr:=strings.Split(scanner.Text(),",")
		a:=account{
			arr[2],
			arr[0],
			arr[1],
		}
		accounts=append(accounts,a)
	}
}
func browserCall(){
		for  {
			select {
				case r:= <-signal:
					if r{
						cmd:=exec.Command(BrowserCmd, HtmlLocation)
						err:=cmd.Start()
						fmt.Println(err)
					}
			}
		}
}

var orderRef = int64(OrderRefStart)

func myOrderRef(i int64)int64{
	m:=sync.Mutex{}
	m.Lock()
	o:=orderRef
	orderRef= orderRef +i
	m.Unlock()
	return o
}

func startServer(){
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		b,_:=ioutil.ReadAll(r.Body)
		//fmt.Println("gmo token:",string(b))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var tok []string
		err:=json.Unmarshal(b,&tok)
		if err !=nil{
			fmt.Println("Error decoding token from browser:", err)
			w.Write([]byte("thanks"))
		}
		token <- tok
		w.Write([]byte("thanks"))
		if orderRef >= OrderRefEnd {
			fmt.Println("We are at end, bye...!")
		}
	})
	http.ListenAndServe(":8080",nil)
}

func getAccount() account{
	m:=sync.Mutex{}
	m.Lock()
	if int(index) == len(accounts){
		m.Unlock()
		return account{}
	}
	a:=accounts[index]
	atomic.AddInt32(&index,1)
	m.Unlock()
	return a
}

func worker(id int,s chan bool,t chan []string, wg *sync.WaitGroup){
	for{
		terminate:=false
		select {
		case t:= <-t:
			//fmt.Println(fmt.Sprintf("processing with worker : %v",id))
			txnId:= fmt.Sprintf("0000%v", myOrderRef(1))
			a:=getAccount()
			if a.BillingAccountNumber == ""{
				fmt.Println(fmt.Sprintf("worker %v done",id))
				terminate = true
				break
			}
			//fmt.Println(fmt.Sprintf("ban:%v , token: %v , txnID: %v", a.BillingAccountNumber,t,txnId))
			res,needToHang:=makePayment(t,txnId,txnId)
			if needToHang {
				for{
					time.Sleep(time.Second*30)
					res,needToHang=makePayment(t,txnId,txnId)
					if !needToHang{
						break
					}
					fmt.Println(fmt.Sprintf("worker %v is hagging on makePayment",id))
				}
			}
			if res {
				err:=mySqlInsert(a.BillingAccountNumber, a.CustomerAccountNumber,a.ServiceInstanceNumber,"1", "gmo", txnId)
				if err{
					for{
						time.Sleep(time.Second*5)
						err=mySqlInsert(a.BillingAccountNumber, a.CustomerAccountNumber,a.ServiceInstanceNumber,"1", "gmo", txnId)
						if !err{
							break
						}
						fmt.Println(fmt.Sprintf("worker %v is hagging on mySqlInsert",id))
					}
				}
				logBillingAccount(a.BillingAccountNumber,txnId)
			}
			s <- true
		}
		if terminate {
			myOrderRef(-1)
			wg.Done()
			s <- true
			break
		}
	}
}
func makePayment( tokens []string, uuid string, txnId string)(bool,bool){
	url := PaymentUrl
	method := "POST"
	s:=fmt.Sprintf(`{
    "uuid": "%v",
    "txn_id": "%v",
    "amount": 1000,
    "currency": "JPY",
    "additional_data": {
        "user_name": "",
        "user_email": "avinash140294@gmail.com",
        "description": "Povo Big -  #000001613980821809",
        "shopper_statement": "Povo Big ",
        "delivery_method": "",
        "payment_methods": [
            "gmo"
        ],
        "success_redirect_url": "https://www.circles.life/sg/",
        "failure_redirect_url": "https://www.circles.life/sg/",
        "return_url": "https://www.circles.life/sg/",
        "first_six_digits": "411111",
        "last_four_digits": "1111",
        "card_type": "Visa",
        "token": ["%v", "%v"],
        "secondary_phone_number": null,
        "referral_method": "WEB",
        "valid_duration": 120,
        "custom_payment_method": "card"
    },
    "payment_config_category": "LIVE",
    "notification_channel": {
        "type": "http_webhook",
        "method": "POST",
        "url": "http://sjp-intoms.circles.life/api/v3/jp/payment/orders/status",
        "skip_multiple": true
    }
}`,uuid,txnId,tokens[0],tokens[1])

	payload := strings.NewReader(s)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return false,false
	}
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false,true
	}
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return false
	//}
	//fmt.Println(string(body))
	//fmt.Println("status:",res.StatusCode)
	if res.StatusCode == 200{
		return true,false
	}
	return false,false
}
var Db *sql.DB

func dbInit(){
		db, err := sql.Open("mysql", MySqlConnectionString)
		if err != nil {
			panic(err.Error())
		}
		Db = db
}

func mySqlInsert(bAN,cAN,sIN,isDefault,gateway,orderRef string)bool{
	m:=sync.Mutex{}
	m.Lock()
	stmt,err:=Db.Prepare("UPDATE v3_payment_instruments SET billing_account_number=?,customer_account_number=?, service_instance_number=?," +
		" is_default=?,gateway=?," +
		" updated_at=now() where order_ref = ?")
	if err !=nil{
		fmt.Println("Error in preparing insert statement:",err)
		return true
	}
	r,err:=stmt.Exec(bAN,cAN,sIN,isDefault,gateway,orderRef)
	if err !=nil{
		fmt.Println(err)
		return true
	}
	rowsAffected,_:=r.RowsAffected()
	fmt.Println(fmt.Sprintf("Update payment instrument with order_ref:%v, result:%v",orderRef, rowsAffected))
	m.Unlock()
	return false
}


func fileCount()int{
	f,e:=ioutil.ReadDir(AccountFileDir)
	if e!=nil{
		panic(e)
	}
	return len(f)
}
func allLineCount()int64{
	files,e:=ioutil.ReadDir(AccountFileDir)
	if e!=nil{
		panic(e)
	}
	count:= OrderRefStart
	for _,af := range files{
		f,e:=os.Open(AccountFileDir+af.Name())
		if e !=nil{
			panic(e)
		}
		count = count + lineCount(f)
	}
	return int64(count)
}
func lineCount(r io.Reader) int{
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count+1

		case err != nil:
			panic(err)
			return count
		}
	}
}
func findStartFile()int{
	min := 1000
	files,e:=ioutil.ReadDir(AccountFileDir)
	if e!=nil{
		panic(e)
	}
	for _,ff:= range files{
		i,_:=strconv.Atoi(ff.Name())
		if min > i{
			min = i
		}
	}
	return min
}
