package redis

import (
	"errors"

	"github.com/go-redis/redis/v7"
)

func IsNil(err error) bool {
	return errors.Is(err, redis.Nil)
}
