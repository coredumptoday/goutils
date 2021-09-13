package redis

import (
	"sync"

	"github.com/go-redis/redis/v7"
)

var cm *connectManager

type connectManager struct {
	l             sync.Mutex
	clientHandler map[string]*Client
	isAlreadyInit bool
	err           error
}

func init() {
	cm = &connectManager{
		l:             sync.Mutex{},
		clientHandler: make(map[string]*Client),
		isAlreadyInit: false,
	}
}

func Err() error {
	return cm.err
}

func RegisterConnector(name, KeyPrefix string, c *config) {
	if cm.isAlreadyInit {
		cm.err = reInitConnectErr
		return
	}

	cm.l.Lock()
	defer cm.l.Unlock()

	cm.clientHandler[name] = &Client{
		cliName:   name,
		KeyPrefix: KeyPrefix,
		cfg:       c,
	}
}

func Connect() error {
	if cm.err != nil {
		return cm.err
	}
	if cm.isAlreadyInit {
		cm.err = reInitConnectErr
		return cm.err
	}

	hasErr := false

	cm.l.Lock()
	defer func() {
		if hasErr {
			//关闭链接
			Close()
		}
		cm.l.Unlock()
	}()

	for _, c := range cm.clientHandler {
		c.initConnectPool()
		c.clusterConnectionIsOpen()
		if c.Err() != nil {
			cm.err = c.Err()
			hasErr = true
			break
		}
	}

	return cm.err
}

func Close() {
	if cm.err != nil {
		return
	}

	for _, c := range cm.clientHandler {
		err := c.closeConnectPool()
		if err != nil {
			cm.err = err
			continue
		}
	}
}

func GetClient(name string) *Client {
	if cm.err != nil {
		return nil
	}
	return cm.clientHandler[name]
}

func RegisterHook(name string, hook redis.Hook) {
	if cm.err != nil {
		return
	}
	cm.clientHandler[name].addHook(hook)
}
