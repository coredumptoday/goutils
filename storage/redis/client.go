package redis

import (
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

type Client struct {
	cliName   string
	KeyPrefix string
	cfg       *config
	cli       redis.UniversalClient
	err       error
}

func (c *Client) initConnectPool() {
	if c.cfg.ct == Cluster {
		c.cli = redis.NewClusterClient(c.cfg.cfgToClusterOpts())
	} else if c.cfg.ct == Sentinel {
		c.cli = redis.NewFailoverClient(c.cfg.cfgToFailoverOptions())
	} else if c.cfg.ct == Single {
		c.cli = redis.NewClient(c.cfg.cfgToOptions())
	} else {
		c.err = clusterTypeErr
	}
}

func (c *Client) clusterConnectionIsOpen() bool {
	if c.cli == nil {
		return false
	}

	if err := c.cli.Ping().Err(); err != nil {
		c.err = err
		return false
	}
	return true
}

func (c *Client) closeConnectPool() error {
	if c.cli == nil {
		return nil
	}
	return c.cli.Close()
}

func (c *Client) addHook(hook redis.Hook) {
	c.cli.AddHook(hook)
}

func (c *Client) SetKeyPrefix(k string) {
	c.KeyPrefix = k
}

func (c *Client) GetKeyPrefix() string {
	return c.KeyPrefix
}

func (c *Client) Err() error {
	return c.err
}

func (c *Client) fixKey(keyName string) string {
	if c.KeyPrefix != "" {
		return c.KeyPrefix + keyName
	}
	return keyName
}

func (c *Client) cleanKey(keyName string) string {
	if c.KeyPrefix != "" {
		return strings.Replace(keyName, c.KeyPrefix, "", 1)
	}
	return keyName
}

func (c *Client) Get(keyName string) (string, error) {
	if c.err != nil {
		return "", c.err
	}
	return c.cli.Get(c.fixKey(keyName)).Result()
}

func (c *Client) Set(keyName, value string, timeout time.Duration) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Set(c.fixKey(keyName), value, timeout).Err()
}

func (c *Client) SetNX(keyName, value string, timeout time.Duration) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.SetNX(c.fixKey(keyName), value, timeout).Err()
}

func (c *Client) TTL(keyName string) (ttl int64, err error) {
	if c.err != nil {
		return 0, c.err
	}

	duration, err := c.cli.TTL(c.fixKey(keyName)).Result()
	return int64(duration.Seconds()), err
}

func (c *Client) Expire(keyName string, timeout time.Duration) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Expire(c.fixKey(keyName), timeout).Err()
}

func (c *Client) Decr(keyName string) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Decr(c.fixKey(keyName)).Err()
}

func (c *Client) Incr(keyName string) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Incr(c.fixKey(keyName)).Err()
}

func (c *Client) Del(keyName string) (int64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.cli.Del(c.fixKey(keyName)).Result()
}

func (c *Client) Exists(keyName string) (int64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.cli.Exists(c.fixKey(keyName)).Result()
}

func (c *Client) Pipeline(cmds [][]interface{}) ([]string, error) {
	if len(cmds) <= 0 {
		return nil, nil
	}

	pCmds := make([]*redis.Cmd, 0, len(cmds))
	p := c.cli.Pipeline()
	for _, cmd := range cmds {
		pCmds = append(pCmds, p.Do(cmd...))
	}

	_, err := p.Exec()
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(pCmds))
	for _, pcmd := range pCmds {
		tmp, _ := pcmd.Text()
		res = append(res, tmp)
	}

	return res, nil
}

func (c *Client) Do(args ...interface{}) (interface{}, error) {
	if len(args) <= 0 {
		return nil, nil
	}
	res, err := c.cli.Do(args...).Result()
	return res, err
}
