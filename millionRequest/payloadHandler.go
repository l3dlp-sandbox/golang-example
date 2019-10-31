package millionRequestfunc

import (
	"net/http"
	"encoding/json"
	"io"
	"os"
	"github.com/labstack/gommon/log"
)

type PayLoadCollection struct {
	WindowsVersion  string    `json:"version"`
	Token           string    `json:"token"`
	Payloads        []Payload `json:"data"`
}
type Payload struct {

}
func(p *Payload)UploadToS3()error{
	return nil
}
//*********************************************************
func payloadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var content = &PayLoadCollection{}
	err:=json.NewDecoder(io.LimitReader(r.Body,1024)).Decode(&content)
	if err !=nil{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _,payload:=range content.Payloads{
		work:=Job{Payload:payload}
		JobQueue<-work
	}
	w.WriteHeader(http.StatusOK)
}
//*********************************************************
var(
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue = os.Getenv("MAX_QUEUE")
)

type Job struct {
	Payload Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit chan bool
}

func NewWorker(workerPool chan chan Job)Worker  {
	return Worker{
		WorkerPool: workerPool,
		JobChannel:make(chan Job),
		quit:make(chan bool),
	}
}

func(w Worker)Start(){
	go func() {
		for{
			w.WorkerPool<-w.JobChannel
			select {
				case job:= <-w.JobChannel:
					if err:=job.Payload.UploadToS3();err!=nil{
						log.Errorf("Error uploading to S3: %s", err.Error())
					}
			case <-w.quit:
				return
			}
		}
	}()
}
func (w Worker)Stop(){
	go func() {
		w.quit <-true
	}()
}
