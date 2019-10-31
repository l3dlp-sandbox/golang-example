package database

import "fmt"

func InsertQuery(){
	Init()
	_,err:=Db.Query(`INSERT INTO payment_transactions(trip_id,card_id)VALUES(2444444444444555444,69)`)
	if err!=nil{
		fmt.Println(err.Error())
	}
}