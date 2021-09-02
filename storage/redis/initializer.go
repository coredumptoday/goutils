package redis

import (
	"sync"

	"github.com/go-redis/redis/v7"
)

var cm *connectManager

func init() {
	cm = &connectManager{
		l:             sync.Mutex{},
		clientHandler: make(map[string]*client),
		initFlag:      false,
	}
}

func RegisterServerCfg(name, KeyPrefix string, c *config, h redis.Hook) {
	cm.registerClusterCfg(name, KeyPrefix, c, h)
}

func Connect() error {
	cm.InitConnect()
	return cm.err
}

func GetClient(name string) *client {
	return cm.getClient(name)
}
