package millionRequestfunc

import (
	"net/http"
	"encoding/json"
	"fmt"
	"database/sql"
)

func StartServer()  {
	http.HandleFunc("/",handler)
	http.HandleFunc("/normal",handleNormal)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Println(err)
	}

}
var dbInstance *sql.DB
func InitSql(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/testdriver")
	if err!=nil{
		panic(err.Error())

	}
	dbInstance=db
}

var jobQueue chan int

type LoginDetail struct {
	Name string `json:"name"`
	Ids  int   `json:"id"`
}
func handler(w http.ResponseWriter,r *http.Request){
	var l LoginDetail
	for{
		err:=json.NewDecoder(r.Body).Decode(&l)
		if err!=nil{
			fmt.Println(err)
		}else {
			break
			//fmt.Println(l)
		}
	}

	jobQueue<-l.Ids

	w.WriteHeader(http.StatusOK)
	return
}

type MyWorker struct {
	Id int
	MyWorkerPool chan chan int
	MyJobChannel chan int
	quit chan bool
}
func MyNewWorker(workerPool chan chan int,name int)MyWorker{
   return MyWorker{
   		Id: name,
   		MyWorkerPool:workerPool,
   		MyJobChannel:make(chan int),
   		quit:make(chan bool),
   }
}
func (w MyWorker) start() {
	go func() {
		for{
			//worker channel is added workerQueue by itself
			w.MyWorkerPool <-w.MyJobChannel
			select {
				case number:= <-w.MyJobChannel:
					InsertQuery(number,w.Id)
				case<-w.quit:
					return
			}
		}
	}()
}
func InsertQuery(number int,id int){
	//fmt.Printf("worker : %d,value is : %d",id,number)

	for{
		_,err:=dbInstance.Exec(`INSERT INTO test(data)VALUES(?)`,number)
		if err!=nil{
			fmt.Println(err)
		}else {
			break
		}
	}
}
func (w MyWorker)MyWorkerStop(){
	go func() {
		w.quit <-true
	}()
}
type MyDispatcher struct {
	MaxWorkers int
	MyWorkerPool chan chan int
}
func MyNewDispatcher(maxWorker int) *MyDispatcher{
	pool:=make(chan  chan int,maxWorker)
	return &MyDispatcher{
		MaxWorkers:10,//number of workers we would like to start same size of buffered workerQueue
		MyWorkerPool:pool,
	}
}
func(d *MyDispatcher)RunMyWorker(){
	for i:=0;i<d.MaxWorkers;i++{
		//fmt.Println("start worker : "+strconv.FormatInt(int64(i),10))
		worker:=MyNewWorker(d.MyWorkerPool,i+1)
		worker.start()
	}
}
func(d *MyDispatcher)dispatch(){
	for{
		select {
		case number:= <-jobQueue:
			//fmt.Println("received work request")
			go func(number int) {
				jobChannel:= <-d.MyWorkerPool
				jobChannel<-number
			}(number)

		}
	}
}

func MyImplementationMain()  {
	InitSql()
	dispatcher:=MyNewDispatcher(10)
	dispatcher.RunMyWorker()
	jobQueue=make(chan int,10)
	go dispatcher.dispatch()
	StartServer()
}

