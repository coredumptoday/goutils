package storage

import (
	"context"
	"fmt"
	"testing"
	"time"

	redis2 "github.com/coredumptoday/goutils/storage/redis"

	"github.com/go-redis/redis/v7"
)

func TestClusterMGet(t *testing.T) {
	c := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"redis-stg.myhll.cn:30000"}})
	p := c.Pipeline()
	p.Get("bbbbbbbbbb")
	p.Get("ccc")
	p.Get("aaa")
	fmt.Println(p.Exec())
}

func TestInit(t *testing.T) {
	stgCfg, err := redis2.NewConnectCfg([]string{"redis-stg.myhll.cn:30000"}, redis2.Cluster)
	if err != nil {
		fmt.Println(err)
	}
	stgCfg.SetTimeout(500)
	localCfg, err := redis2.NewConnectCfg([]string{"127.0.0.1:6379"}, redis2.Single)
	if err != nil {
		fmt.Println(err)
	}
	localCfg.SetTimeout(500)

	redis2.RegisterServerCfg("local", "", localCfg, &hook{})
	redis2.RegisterServerCfg("stg", "", stgCfg, &hook{})

	if err := redis2.Connect(); err != nil {
		fmt.Println(err)
	}

	mres, err := redis2.GetClient("stg").MultiGet([]string{"bbbbbbbbbb", "ccc", "aaa"})
	if err != nil && !redis2.IsNil(err) {
		fmt.Println(err)
	}
	fmt.Println("=====", mres)

	//fmt.Println("local slots", m.getClient("local").cli.ClusterSlots())
	//fmt.Println("stg slots", m.getClient("stg").cli.ClusterSlots())

	/*fmt.Println(m.getClient("aaa").cli.ClusterSlots())
	fmt.Println(m.getClient("aaa").cli.ClusterKeySlot("aaaaaaaaaa"))
	res, err := m.getClient("aaa").Get("bbbbbbbbbb")
	if err != nil && !IsNil(err) {
		fmt.Println(err)
	}
	fmt.Println("----", res, IsNil(err))
	*/
}

type hook struct {
}

func (h *hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	fmt.Println("BeforeProcess")
	fmt.Println("key", cmd.Args()[1])
	ctx = context.WithValue(ctx, "st", time.Now())
	fmt.Println()
	return ctx, nil
}

func (h *hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	fmt.Println("AfterProcess")
	fmt.Println(cmd)
	tmp := ctx.Value("st")
	st := tmp.(time.Time)
	fmt.Println(time.Since(st).Seconds())
	fmt.Println()
	return nil
}

func (h *hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	fmt.Println("BeforeProcessPipeline")
	fmt.Println(cmds)
	ctx = context.WithValue(ctx, "st", time.Now())
	return ctx, nil
}

func (h *hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	fmt.Println("AfterProcessPipeline")
	fmt.Println(cmds)
	tmp := ctx.Value("st")
	st := tmp.(time.Time)
	fmt.Println(time.Since(st).Nanoseconds())
	return nil
}
