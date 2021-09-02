package redis

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

type client struct {
	cliName   string
	KeyPrefix string
	cfg       *config
	cli       redis.UniversalClient
	err       error
	hook      redis.Hook
	si        []struct {
		low, height int
		env         string
	}
}

func (c *client) initSlotsInfo() {
	if Cluster == c.cfg.ct {
		s, err := c.cli.ClusterSlots().Result()
		fmt.Printf("%#v, %#v", s, err)
	} else {
		c.err = c.cli.Ping().Err()
	}
}

func (c *client) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	if c.hook != nil {
		return c.hook.BeforeProcess(ctx, cmd)
	}
	return ctx, nil
}

func (c *client) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	if c.hook != nil {
		return c.hook.AfterProcess(ctx, cmd)
	}
	return nil
}

func (c *client) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	if c.hook != nil {
		return c.hook.BeforeProcessPipeline(ctx, cmds)
	}
	return ctx, nil
}

func (c *client) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	if c.hook != nil {
		return c.hook.AfterProcessPipeline(ctx, cmds)
	}
	return nil
}

func (c *client) SetKeyPrefix(k string) {
	c.KeyPrefix = k
}

func (c *client) GetKeyPrefix() string {
	return c.KeyPrefix
}

func (c *client) Err() error {
	return c.err
}

func (c *client) fixKey(keyName string) string {
	return c.KeyPrefix + keyName
}

func (c *client) cleanKey(keyName string) string {
	return strings.Replace(keyName, c.KeyPrefix, "", 1)
}

func (c *client) clusterConnectionIsOpen() bool {
	if c.cli == nil {
		return false
	}

	if err := c.cli.Ping().Err(); err != nil {
		c.err = err
		return false
	}
	return true
}

func (c *client) Get(keyName string) (string, error) {
	if c.err != nil {
		return "", c.err
	}
	return c.cli.Get(c.fixKey(keyName)).Result()
}

func (c *client) Set(keyName, value string, timeout time.Duration) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Set(c.fixKey(keyName), value, timeout).Err()
}

func (c *client) SetNX(keyName, value string, timeout time.Duration) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.SetNX(c.fixKey(keyName), value, timeout).Err()
}

func (c *client) TTL(keyName string) (ttl int64, err error) {
	if c.err != nil {
		return 0, c.err
	}

	duration, err := c.cli.TTL(c.fixKey(keyName)).Result()
	return int64(duration.Seconds()), err
}

func (c *client) Expire(keyName string, timeout time.Duration) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Expire(c.fixKey(keyName), timeout).Err()
}

func (c *client) Decr(keyName string) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Decr(c.fixKey(keyName)).Err()
}

func (c *client) Incr(keyName string) error {
	if c.err != nil {
		return c.err
	}
	return c.cli.Incr(c.fixKey(keyName)).Err()
}

func (c *client) Del(keyName string) (int64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.cli.Del(c.fixKey(keyName)).Result()
}

func (c *client) Exists(keyName string) (int64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.cli.Exists(c.fixKey(keyName)).Result()
}

func (c *client) MultiGet(keys []string) (map[string]string, error) {
	if c.err != nil {
		return nil, c.err
	}

	kCnt := len(keys)

	if kCnt <= 0 {
		return nil, nil
	}

	keyNames := make([]string, kCnt)
	copy(keyNames, keys)
	for index, val := range keyNames {
		keyNames[index] = c.fixKey(val)
	}

	result := make(map[string]string, kCnt)

	if c.cfg.ct == Cluster {
		getCmds := make([]*redis.StringCmd, 0, kCnt)
		pipe := c.cli.Pipeline()
		for _, key := range keyNames {
			getCmds = append(getCmds, pipe.Get(key))
		}
		_, err := pipe.Exec()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
		for i, cmd := range getCmds {
			if cmd.Err() != nil {
				if errors.Is(cmd.Err(), redis.Nil) {
					continue
				} else {
					return nil, cmd.Err()
				}
			} else {
				result[keys[i]] = cmd.Val()
			}
		}
	} else {
		values, err := c.cli.MGet(keyNames...).Result()
		if err != nil {
			return nil, err
		}
		for i, val := range values {
			if nil == val {
				continue
			}
			result[keys[i]] = val.(string)
		}
	}

	return result, nil
}
