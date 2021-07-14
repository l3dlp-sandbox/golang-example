package dbconnection

import (
	"gitlab.mytaxi.lk/pickme/go-util/mysql"
	mq "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
	"context"
)

func TestConnectionOff(){
	Init(
		database.WithReadConfig(&mq.Config{
			//Timeout: 100 * time.Millisecond,
			//ReadTimeout: 100 * time.Millisecond,

		}),
		database.WithWriteConfig(&mq.Config{
			//Timeout: 100 * time.Millisecond,
			//ReadTimeout: 100 * time.Millisecond,

		}),
	)

	for{
		start:=time.Now()
		dbCall()
		fmt.Println(time.Since(start))
		//return
	}
}

func dbCall(){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Millisecond)
	q:="SELECT * FROM driver_balance_summary WHERE driver_id=188"
	database.Connections.Read.ExecContext(ctx,q)
}
