package config_load

import (
	"github.com/olebedev/config"
	"log"
	"fmt"
)

func ConfigLoad(){
	Cfg,err:=config.ParseJsonFile("settings.json")
	if err !=nil{
		log.Println("in finding file"+err.Error())
	}
	location,err:=Cfg.String("database.mysql.url")
	if err!=nil{
		log.Println(err.Error())
	}
	fmt.Println(location)
}
