package util

import (
	"testing"

	"encoding/json"
	"gitlab.mytaxi.lk/pickme/go-util/log"
)

type person struct {
	Name   string
	Id     int
	Gender string
}

func TestPrintStruct(t *testing.T) {
	v := &person{"damithaa", 1234, "male"}
	out, _ := json.Marshal(v)

	log.Info(log.WithPrefix("_test.util.TestPrintStruct", string(out)))

}
