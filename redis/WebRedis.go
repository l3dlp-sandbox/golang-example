package redis

import (
	"net/http"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/labstack/gommon/log"
	"encoding/json"
	"fmt"
)


var db *pool.Pool

func CreateConnectionPool()  {
	conn,err:=pool.New("tcp","localhost:6379",10)
	if err!=nil{
		log.Panic(err)
	}
	db=conn
}

func WebRedis(){
	http.HandleFunc("/show",showAlbum)
	err:=http.ListenAndServe(":3000",nil)
	if err !=nil{
		fmt.Println(err)
	}
}

func showAlbum(w http.ResponseWriter, r * http.Request){
	if r.Method !="GET"{
		http.Error(w,http.StatusText(405),405)
		return
	}

	id:=r.URL.Query().Get("id")

	conn,err :=db.Get()
	if err!=nil{
		log.Fatal(err)
	}

	defer db.Put(conn)

	reply,err := conn.Cmd("HGETALL","album:"+id).Map()

	x,err:=json.MarshalIndent(reply,"","")

	fmt.Println(string(x))

}
