package redis

import (
	"errors"
	"sync"

	"github.com/go-redis/redis/v7"
)

var ReInitConnect = errors.New("redis: ConnectManager is already init connect")

type connectManager struct {
	l             sync.Mutex
	clientHandler map[string]*client
	initFlag      bool
	err           error
}

func (cm *connectManager) registerClusterCfg(name, KeyPrefix string, c *config, h redis.Hook) {
	if cm.initFlag {
		cm.err = ReInitConnect
		return
	}
	cm.l.Lock()
	defer cm.l.Unlock()
	cm.clientHandler[name] = &client{
		cliName:   name,
		KeyPrefix: KeyPrefix,
		cfg:       c,
		hook:      h,
	}
}

func (cm *connectManager) getClient(name string) *client {
	return cm.clientHandler[name]
}

func (cm *connectManager) InitConnect() {
	cm.l.Lock()
	defer cm.l.Unlock()

	for _, c := range cm.clientHandler {
		c.cli = c.cfg.initConnectPool()
		c.initSlotsInfo()
		if c.Err() != nil {
			cm.err = c.Err()
			return
		}
		c.cli.AddHook(c.hook)
	}
}
