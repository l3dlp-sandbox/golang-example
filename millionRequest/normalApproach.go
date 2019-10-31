package millionRequestfunc

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func handleNormal(w http.ResponseWriter,r *http.Request){
	var l LoginDetail
	for{
		err:=json.NewDecoder(r.Body).Decode(&l)
		if err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println(l)
			break
		}
	}
	InsertQuery(l.Ids,1)
}