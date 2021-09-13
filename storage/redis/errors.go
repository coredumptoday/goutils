package redis

import (
	"errors"

	"github.com/go-redis/redis/v7"
)

var reInitConnectErr = errors.New("redis: ConnectManager is already init connect")
var clusterTypeErr = errors.New("redis: cluster type set error")

func IsReInitConnectErr(err error) bool {
	return errors.Is(err, reInitConnectErr)
}

func IsClusterTypeErr(err error) bool {
	return errors.Is(err, clusterTypeErr)
}

func IsNilErr(err error) bool {
	return errors.Is(err, redis.Nil)
}
