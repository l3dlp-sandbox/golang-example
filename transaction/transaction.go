package transaction

import (
	"database/sql"
	"gitlab.mytaxi.lk/pickme/go-util/log"
	_ "github.com/go-sql-driver/mysql"
)

func Transaction(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/driver_disconnect")
	if err!=nil{
		log.Error(log.WithPrefix("ErrorMessage",err))
	}
	tx,err:=db.Begin()
	q:="INSERT INTO driver_balance_summary (driver_id,balance) VALUES (4568,45.23)"
	tx.Exec(q)
	tx.Commit()
}
