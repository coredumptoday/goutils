package redis

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/go-redis/redis/v7"
)

const TimeoutCFG = time.Duration(100) * time.Millisecond
const MinIdleConns = 10
const ConnectMaxRetries = 5

var CfgAddrsIsEmpty = errors.New("redis: config addrs is empty")

type clusterType int

const (
	Single clusterType = 1 << iota
	Sentinel
	Cluster
)

func NewConnectCfg(addrs []string, ct clusterType) (*config, error) {
	if len(addrs) == 0 {
		return nil, CfgAddrsIsEmpty
	}

	return &config{
		addrs:                 addrs,
		database:              0,
		username:              "",
		password:              "",
		masterName:            "",
		minIdle:               0,
		maxConnPerNode:        0,
		timeout:               0,
		useTLS:                false,
		tlsInsecureSkipVerify: false,
		ct:                    ct,
	}, nil
}

type config struct {
	addrs                 []string //集群地址[]string{"127.0.0.1:6379"}
	database              int      //数据库编号，只适用于单节点和sentinel集群模式
	username              string
	password              string
	masterName            string //sentinel集群模式
	minIdle               int    //空闲连接数
	maxConnPerNode        int    //最大连接数
	timeout               int64  //超时时间（单位：毫秒）
	useTLS                bool   //使用tls
	tlsInsecureSkipVerify bool
	ct                    clusterType
}

func (c *config) SetDB(no int) *config {
	c.database = no
	return c
}

func (c *config) SetUser(user string) *config {
	c.username = user
	return c
}

func (c *config) SetPwd(pwd string) *config {
	c.password = pwd
	return c
}

func (c *config) SetMasterName(name string) *config {
	c.masterName = name
	return c
}

func (c *config) SetMinIdle(co int) *config {
	c.minIdle = co
	return c
}

func (c *config) SetMaxConnPerNode(co int) *config {
	c.maxConnPerNode = co
	return c
}

func (c *config) SetTimeout(co int64) *config {
	c.timeout = co
	return c
}

func (c *config) EnableSSL() *config {
	c.useTLS = true
	return c
}

func (c *config) UnableSSL() *config {
	c.useTLS = false
	return c
}

func (c *config) SSLVerify() *config {
	c.tlsInsecureSkipVerify = true
	return c
}

func (c *config) SkipSSLVerify() *config {
	c.tlsInsecureSkipVerify = false
	return c
}

func (c *config) initConnectPool() redis.UniversalClient {
	//连接数
	poolSize := MinIdleConns
	if c.maxConnPerNode > 0 {
		poolSize = c.maxConnPerNode
	}

	//超时时间
	timeout := TimeoutCFG
	if c.timeout > 0 {
		timeout = time.Duration(c.timeout) * time.Millisecond
	}

	if c.ct == Cluster {
		opts := &redis.ClusterOptions{
			Addrs:        c.addrs,
			Username:     c.username,
			Password:     c.password,
			MaxRetries:   ConnectMaxRetries,
			DialTimeout:  timeout,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			MinIdleConns: MinIdleConns,
			PoolSize:     poolSize,
			IdleTimeout:  3 * timeout,
		}
		if c.useTLS {
			opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.tlsInsecureSkipVerify}
		}
		return redis.NewClusterClient(opts)
	} else if c.ct == Sentinel {
		opts := &redis.FailoverOptions{
			SentinelAddrs: c.addrs,
			DB:            c.database,
			Username:      c.username,
			Password:      c.password,
			MaxRetries:    ConnectMaxRetries,
			DialTimeout:   timeout,
			ReadTimeout:   timeout,
			WriteTimeout:  timeout,
			MinIdleConns:  MinIdleConns,
			PoolSize:      poolSize,
			IdleTimeout:   3 * timeout,
			MasterName:    c.masterName,
		}
		if c.useTLS {
			opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.tlsInsecureSkipVerify}
		}
		return redis.NewFailoverClient(opts)
	} else {
		opts := &redis.Options{
			Addr:         c.addrs[0],
			DB:           c.database,
			Username:     c.username,
			Password:     c.password,
			MaxRetries:   ConnectMaxRetries,
			DialTimeout:  timeout,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			MinIdleConns: MinIdleConns,
			PoolSize:     poolSize,
			IdleTimeout:  3 * timeout,
		}
		if c.useTLS {
			opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.tlsInsecureSkipVerify}
		}
		return redis.NewClient(opts)
	}
}
