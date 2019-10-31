package database

import (
	"context"
	"fmt"
	"database/sql"
)

type SqlError struct {
	 Status string
	 DriverId int

}

func (s *SqlError)RunQuery()(){

	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/driver_disconnect")
	if err!=nil{
		fmt.Println("error connection")
		panic(err.Error())

	}

	Db=db
	ctx:=context.Background()
	tx,_:=Db.Begin()

	query:="UPDATE driver SET login_status=?  WHERE id=? "

	if tx!=nil{
		r,err1 :=tx.ExecContext(ctx,query,s.Status,s.DriverId)
		if err1!=nil{
			fmt.Println("errrrrrrrrrrrrrrr")
		}else {
			fmt.Println(r)
		}

	}



	err3:=tx.Commit()

	if err3!=nil{
		fmt.Println("err when commit")
	}


}
