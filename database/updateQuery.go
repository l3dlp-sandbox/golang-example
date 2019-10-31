package database

import (
	"database/sql"
	"fmt"
)

func UpdateQuery(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/testdriver")
	if err != nil{
		fmt.Println(err)
	}

	defer db.Close()

	db.Exec("UPDATE payment_cards SET is_active=0,date_added=date_added  WHERE id=?",5)

}

