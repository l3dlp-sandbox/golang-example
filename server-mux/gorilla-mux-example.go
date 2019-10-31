package server_mux

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"

	"gitlab.mytaxi.lk/pickme/go-util/log"
	"encoding/json"
)

func ServerMux(){
	r:=mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}",router)
	r.HandleFunc("/check",nopageFound)
	r.HandleFunc("/scheck",pageFound)
	r.HandleFunc("/card",ComplexResponse)

	err:=http.ListenAndServe(":8080",r)
	if err!=nil{
		log.Fatal("error happened",err)
	}

}

func nopageFound(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprint(w,"where the shit url")
}

func pageFound(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprint(w,"here i'm")
}

func router(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	title:=vars["title"]
	page:=vars["page"]
	fmt.Fprintf(w,"title-%s,page-%s",title,page)
}



type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    string `json:"data"`
}

type Card struct {
	id int
	bank string
	city string
}

func ComplexResponse(rw http.ResponseWriter, re *http.Request)  {
	c:=Card{
		id:123,
		bank:"commercial",
		city:"Nugegoda",
	}
	data,_:=json.MarshalIndent(c,"","")
	D:=string(data)
	r:=Response{
		Error:true,
		Message:"ok",
		Data:D,
	}


	resp,_:=json.Marshal(r)

	rw.Header().Add("Content-Type","application/json")
	rw.WriteHeader(200)
	fmt.Fprint(rw,string(resp))
}
