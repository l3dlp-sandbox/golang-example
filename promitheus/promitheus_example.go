package promitheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"time"
	"gitlab.mytaxi.lk/pickme/go-util/mysql"
	"runtime"
)

var(
	QueryTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:"Db_query_time",
			Help:"count_db_query_time",
			Namespace:"driver_ledger",
			Subsystem:"driver_ledger_database",
		},
		[]string{"method","error"})

	NumGoRoutine = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "no_of_go_routine",
			Help:      "no of go routine",
			Namespace: "driver_finance",
			Subsystem: "driver_ledger_transaction",
		},
		[]string{"method", "error"})

	pushCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
		Name:"increment_count",
		Help:"incomming request count",
		Namespace: "driver_finance",
		Subsystem: "driver_ledger_transaction_test_count",
	},)


	SqlConnection = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:"sql_connection_count",
			Namespace:"driver_finance",
			Help:"Sql connection count",
			Subsystem: "driver_ledger_transaction_sqlConnection",
		},
		[]string{"method", "error"})
	SuccessRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "current_connections",
			Help:      "incomming request count",
			Namespace: "driver_finance",
			Subsystem: "driver_ledger_transaction",
		},
		[]string{"method", "error"})
)

func init(){
	prometheus.MustRegister(NumGoRoutine)
	prometheus.MustRegister(SqlConnection)
	prometheus.MustRegister(SuccessRequests)
	prometheus.MustRegister(QueryTime)
	prometheus.MustRegister(pushCounter)
	fmt.Println("here i'm")
}

func Example()  {
	start:=time.Now()
	var count=5.0
	fmt.Println("jfsjbfjsb")

	pushCounter.Inc()

	NumGoRoutine.With(prometheus.Labels{
		"method":"/publish",
		"error" : "",
	}).Set(float64(runtime.NumGoroutine()))

	time.Sleep(time.Second*3)
	database.Connections.Read.Ping()
	//conn:=database.Connections.Read.Stats().OpenConnections
	SqlConnection.With(prometheus.Labels{
		"method":"/publish",
		"error" : "",
	}).Set(float64(database.Connections.Read.Stats().OpenConnections))

	SuccessRequests.With(prometheus.Labels{
		"method":"/publish",
		"error" : "",
	}).Set(count)

	elapsed:=time.Since(start)

	QueryTime.With(
		prometheus.Labels{
			"method" : "/database",
			"error" : "",
		}).Observe(float64(elapsed.Nanoseconds()/1000))

}
func check(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"urllllllllll")
}
func ServerInit(){

	router:=mux.NewRouter()
	router.Handle("/metrics",promhttp.Handler())
	router.HandleFunc("/check",check)
	go http.ListenAndServe(":8080",router)


}
