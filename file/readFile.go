package file

import (
	"io/ioutil"
	"fmt"
)

func ReadFile(){
	configContent,err:= ioutil.ReadFile(`settings.json`)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(configContent))
}
