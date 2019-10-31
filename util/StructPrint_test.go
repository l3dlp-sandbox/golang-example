package util

import (
	"testing"

	"gitlab.mytaxi.lk/pickme/go-util/log"
	"encoding/json"
)

type person struct{
	Name string
	Id int
	Gender string
}

func TestPrintStruct(t *testing.T) {
	v:=&person{"damithaa",1234,"male"}
	out,_:=json.Marshal(v)


	log.Info(log.WithPrefix("test.util.TestPrintStruct",string(out)))

}
