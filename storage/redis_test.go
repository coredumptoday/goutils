package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/coredumptoday/goutils/storage/redis"
)

func initRedisClient() {
	sc, err := redis.NewConnector([]string{"127.0.0.1:6379"}, redis.Single)
	if err != nil {
		fmt.Println(err)
	}
	redis.RegisterConnector("simpleClint", "xxx_", sc)
	err = redis.Connect()
	if err != nil {
		fmt.Println(err)
	}
}

func TestRedisSimple(t *testing.T) {
	initRedisClient()
	defer redis.Close()

	sc := redis.GetClient("simpleClint")

	err := sc.Set("bbbbb", "1", 3*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	reply, err := sc.Get("bbbbb")
	if err != nil && !redis.IsNilErr(err) {
		fmt.Println(err)
	}
	fmt.Println("GET bbbbb", reply)

	ttl, err := sc.TTL("bbbbb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ttl bbbbb", ttl)

	time.Sleep(4 * time.Second)

	reply, err = sc.Get("bbbbb")
	if err != nil && !redis.IsNilErr(err) {
		fmt.Println(err)
	}
	fmt.Println("GET bbbbb", reply)
}
