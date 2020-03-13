package service

import (
	"log"
	"time"

	"gamenews.niracler.com/monitor/setting"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var redisPool *pool.Pool

func df(network, addr string) (*redis.Client, error) {
	client, err := redis.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	if err = client.Cmd("AUTH", setting.RedisPass).Err; err != nil {
		client.Close()
		return nil, err
	}
	return client, nil
}

// 创建redis连接池
func newRedisPool() *pool.Pool {
	redisPool, err := pool.NewCustom("tcp", setting.RedisHost, 2*5, df)
	if err != nil {
		log.Fatal("Redis pool created failed.")
	} else {
		go func() {
			for {
				redisPool.Cmd("PING")
				time.Sleep(3 * time.Second)
			}
		}()
	}

	return redisPool
}

// 初始化Redis连接池
func init() {
	redisPool = newRedisPool()
}

// 返回Redis连接池
func RedisPool() *pool.Pool {
	return redisPool
}
