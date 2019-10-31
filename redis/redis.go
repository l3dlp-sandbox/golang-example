package redis

import (
	"github.com/mediocregopher/radix.v2/redis"
	"gitlab.mytaxi.lk/pickme/go-util/log"
	"fmt"
	"encoding/json"
)

var ClientRedis *redis.Client

type album struct {
	title string
	author string
	price float64
	likes int64
}


func ConnectRedis()  {
	conn,err := redis.Dial("tcp","localhost:6379")
	ClientRedis =conn
	if err != nil{
		log.Fatal(err)
	}

}

func GoPutHashRedis(){

	resp :=ClientRedis.Cmd("HMSET","album:5","title","2012","author","ACC","price", 250.00,"likes",10000)

	if resp.Err !=nil{
		log.Fatal(resp.Err)
	}

	fmt.Println(resp.String())
}


func GoGetFromHashRedisOne()  {

	title, err := ClientRedis.Cmd("HGET", "album:11", "title").Str()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(title)

}

func GoGetFromHashAll(){
	allRecord,err := ClientRedis.Cmd("HGETALL","album:1").Map()
	if err !=nil{
		log.Fatal(err)
	}

	x,err:=json.MarshalIndent(allRecord,"","")

	fmt.Println(string(x))
}


