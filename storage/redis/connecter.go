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

type ClusterType int

const (
	Single ClusterType = 1 << iota
	Sentinel
	Cluster
)

func NewConnector(addrs []string, ct ClusterType) (*config, error) {
	if len(addrs) == 0 {
		return nil, CfgAddrsIsEmpty
	}

	return &config{
		addrs:                 addrs,
		database:              0,
		username:              "",
		password:              "",
		masterName:            "",
		minIdle:               MinIdleConns,
		maxConnPerNode:        MinIdleConns,
		connectTimeout:        100,
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
	connectTimeout        int64  //连接超时时间（单位：毫秒）
	readTimeout           int64
	writeTimeout          int64
	useTLS                bool //使用tls
	tlsInsecureSkipVerify bool
	ct                    ClusterType
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

func (c *config) SetConnectTimeout(co int64) *config {
	c.connectTimeout = co
	return c
}

func (c *config) SetReadTimeout(co int64) *config {
	c.readTimeout = co
	return c
}

func (c *config) SetWriteTimeout(co int64) *config {
	c.writeTimeout = co
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

func (c *config) cfgToClusterOpts() *redis.ClusterOptions {
	//连接数
	poolSize := MinIdleConns
	if c.maxConnPerNode > 0 {
		poolSize = c.maxConnPerNode
	}

	//超时时间
	connectTimeout := TimeoutCFG
	readTimeout := TimeoutCFG
	writeTimeout := TimeoutCFG
	if c.connectTimeout > 0 {
		connectTimeout = time.Duration(c.connectTimeout) * time.Millisecond
	}
	if c.readTimeout > 0 {
		readTimeout = time.Duration(c.readTimeout) * time.Millisecond
	}
	if c.writeTimeout > 0 {
		writeTimeout = time.Duration(c.writeTimeout) * time.Millisecond
	}

	minIdleCnt := MinIdleConns
	if c.minIdle > 0 {
		minIdleCnt = c.minIdle
	}

	opts := &redis.ClusterOptions{
		Addrs:        c.addrs,
		Username:     c.username,
		Password:     c.password,
		MaxRetries:   ConnectMaxRetries,
		DialTimeout:  connectTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		MinIdleConns: minIdleCnt,
		PoolSize:     poolSize,
		IdleTimeout:  3 * connectTimeout,
	}

	if c.useTLS {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.tlsInsecureSkipVerify}
	}

	return opts
}

func (c *config) cfgToFailoverOptions() *redis.FailoverOptions {
	//连接数
	poolSize := MinIdleConns
	if c.maxConnPerNode > 0 {
		poolSize = c.maxConnPerNode
	}

	//超时时间
	connectTimeout := TimeoutCFG
	readTimeout := TimeoutCFG
	writeTimeout := TimeoutCFG
	if c.connectTimeout > 0 {
		connectTimeout = time.Duration(c.connectTimeout) * time.Millisecond
	}
	if c.readTimeout > 0 {
		readTimeout = time.Duration(c.readTimeout) * time.Millisecond
	}
	if c.writeTimeout > 0 {
		writeTimeout = time.Duration(c.writeTimeout) * time.Millisecond
	}

	minIdleCnt := MinIdleConns
	if c.minIdle > 0 {
		minIdleCnt = c.minIdle
	}

	opts := &redis.FailoverOptions{
		SentinelAddrs: c.addrs,
		DB:            c.database,
		Username:      c.username,
		Password:      c.password,
		MaxRetries:    ConnectMaxRetries,
		DialTimeout:   connectTimeout,
		ReadTimeout:   readTimeout,
		WriteTimeout:  writeTimeout,
		MinIdleConns:  minIdleCnt,
		PoolSize:      poolSize,
		IdleTimeout:   3 * connectTimeout,
		MasterName:    c.masterName,
	}

	if c.useTLS {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.tlsInsecureSkipVerify}
	}

	return opts
}

func (c *config) cfgToOptions() *redis.Options {
	//连接数
	poolSize := MinIdleConns
	if c.maxConnPerNode > 0 {
		poolSize = c.maxConnPerNode
	}

	//超时时间
	connectTimeout := TimeoutCFG
	readTimeout := TimeoutCFG
	writeTimeout := TimeoutCFG
	if c.connectTimeout > 0 {
		connectTimeout = time.Duration(c.connectTimeout) * time.Millisecond
	}
	if c.readTimeout > 0 {
		readTimeout = time.Duration(c.readTimeout) * time.Millisecond
	}
	if c.writeTimeout > 0 {
		writeTimeout = time.Duration(c.writeTimeout) * time.Millisecond
	}

	minIdleCnt := MinIdleConns
	if c.minIdle > 0 {
		minIdleCnt = c.minIdle
	}

	opts := &redis.Options{
		Addr:         c.addrs[0],
		DB:           c.database,
		Username:     c.username,
		Password:     c.password,
		MaxRetries:   ConnectMaxRetries,
		DialTimeout:  connectTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		MinIdleConns: minIdleCnt,
		PoolSize:     poolSize,
		IdleTimeout:  3 * connectTimeout,
	}

	if c.useTLS {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.tlsInsecureSkipVerify}
	}

	return opts
}
