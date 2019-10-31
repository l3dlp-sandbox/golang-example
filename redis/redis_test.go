package redis

import "testing"

func Test_redis(t *testing.T){
	//ConnectRedis()
	//GoPutHashRedis()
	//GoGetFromHashRedisOne()
	//GoGetFromHashAll()
	//defer func() {
	//	ClientRedis.Close()
	//}()

	CreateConnectionPool()
	WebRedis()
}
