package dbconnection

import (
	"database/sql"
	"fmt"
)

type Data struct {
	ID int `json:"passenger_log_id"`
	Booking_Key string `json:"booking_key"`
}

func GoSelect(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/testdriver")
	if err!=nil{
		fmt.Println("error connection")
		panic(err.Error())

	}
	defer db.Close()
	//results,err1:= db.Query("SELECT passengers_log_id,booking_key FROM passengers_log WHERE passengers_log_id=?",17973918)
	results,err1:= db.Query("SELECT passengers_log_id,booking_key FROM passengers_log WHERE passengers_log_id = ?",382792492)

	if err1 != nil{
			fmt.Println(err1)
	}
	for results.Next(){
		var data Data
		//err2:=results.Scan(&data.ID,&data.Booking_Key)
		err2:=results.Scan(&data.ID,&data.Booking_Key)
		if err2 != nil{
			fmt.Println(err2)
		}
		fmt.Println(data.ID)
	}
	defer results.Close()
}
func SelectFromPassengers(){
	var passengerIdFromDb int64
	row,err1:=Db.Query("SELECT id FROM passengers WHERE id=?",3)
	defer row.Close()
	for row.Next(){
		row.Scan(&passengerIdFromDb)
	}
	if err1!=nil{
		fmt.Println(err1.Error())
	}
	err:=row.Scan(&passengerIdFromDb)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(passengerIdFromDb)
}
