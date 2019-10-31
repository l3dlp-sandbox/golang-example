package context

import (
	"gitlab.mytaxi.lk/pickme/go-util/datasource"
	"context"
	"gitlab.mytaxi.lk/pickme/go-util/mysql"
	"net/http"
	"fmt"
	"io/ioutil"
)

/*
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
*/


func RollbackFunction() {

	ctx,cancel:=context.WithCancel(context.Background())

	defer func() {
		cancel()
	}()

	datasource.RunInTransaction(ctx, func(ct context.Context) error {

		err1:=DatabaseQuery(ct)
		if err1!=nil{
			return err1
		}
		err2 := ApiCall()
		if err2!=nil{
			return err2
		}
		return nil
	},nil)
}

func DatabaseQuery(ctx context.Context)error{
	tx:=datasource.FromAnotherTransaction(ctx)

	q:="INSERT INTO driver_balance_summary(driver_id,balance) VALUES (1,456.23)"
	//database.Connections.Write.Exec(q)

	if tx!=nil{
		tx.ExecContext(ctx,q)
	}else{
		database.Connections.Write.ExecContext(ctx,q)
	}
	return nil
}

func ApiCall() error  {
	resp,err:=http.Get("http://example.com")
	if err!=nil{
		return err
	}
	htmlData,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(htmlData))

	return nil
}