package dbconnection

import (
	"fmt"
)
//var wg sync.WaitGroup

func RunParalelWrite()  {
	//wg.Add(100)
	for i:=0;i<100;i++{
		 sqlWrite(i)
	}
	//wg.Wait()
}

func sqlWrite(i int){
	//defer wg.Done()
	_,err:=Db.Exec("INSERT INTO _test(data) VALUES(?)",i)
	if err!=nil{
		fmt.Println(err)
	}
}