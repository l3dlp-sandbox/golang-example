package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"context"
	"fmt"
	"time"
)

var Db *sql.DB



func Init(){

	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/testdriver")
	//db,err:=sql.Open("mysql","surath:t5uveBaf@(192.168.180.132:3307)/Eztaxi")
	if err!=nil{
		panic(err.Error())

	}

	Db=db
}

func Database()bool{
	ch:=make(chan bool,1)
	ctx,cancel:=context.WithTimeout(context.Background(),100*time.Millisecond)
	defer cancel()
	go func(ch chan bool,ctx context.Context) {
		Db.Exec("SELECT 1")
		err:=Db.PingContext(ctx)
		if err!=nil{
			fmt.Println(err)
		}
		ch<-true
	}(ch,ctx)

	select {
	case <-ctx.Done():return false
	case <-ch:return true
	}



	//select {
	//case <-ctx.Done():
	//	return
	//}


	//err:=Db.Ping()
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//return


	//i := "INSERT INTO driver_trip_transaction (amount,createdTime,description,driverId,transactionCategory,transactionCategoryName,transactionId,transactionType,tripId) VALUES (?,?,?,?,?,?,?,?,?)"
	//
	//ctx,cancel:=context.WithTimeout(context.Background(),100*time.Millisecond)
	//defer cancel()
	//
	//query(ctx,i)


	//_,err:=Db.ExecContext(ctx,i,54.23,45,"khkhk",5,8,"klk",7,"lkk",58)
	//fmt.Println(err)

}

func query(ctx context.Context, q string) bool{
	ch :=make(chan bool)
	go func(ctx context.Context,c chan bool) {
		Db.Exec(q,54.23,45,"khkhk",5,8,"klk",7,"lkk",58)
		Db.ExecContext(ctx,q,54.23,45,"khkhk",5,8,"klk",7,"lkk",58)
		//fmt.Println("stil here")
		c<-true
	}(ctx,ch)
	select {
	case <-ch:
		return true
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return false
	}

}

