// Package rdb @author uangi 2023-05
package rdb

import (
	"context"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/config"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var option *redis.Options

var clientLock sync.Mutex

var client *redis.Client

var ctx = context.Background()

var logger = plog.New("redis")

const (
	DefaultTtl = 86400
)

// GetClient 获取已有的连接
func GetClient() *redis.Client {
	if client == nil {
		if clientLock.TryLock() {
			defer clientLock.Unlock()
			if client == nil {
				client = redis.NewClient(option)
			}
		} else {
			for {
				time.Sleep(5 * time.Millisecond)
				return GetClient()
			}
		}
	}
	return client
}

// Init 启动 建议放在main函数内
func Init() {
	c := config.GetPropertiesRO().Server.Redis
	option = &redis.Options{
		Addr:         c.Addr,     //连接地址
		Password:     c.Password, //密码
		DB:           c.Db,       //库
		PoolSize:     c.PoolMax,  //连接池大小8
		MinIdleConns: c.PoolMin,  //最小空闲 2
	}
	GetClient()
	if client != nil {
		logger.Info("Redis client connected to " + c.Addr)
		ping := client.Ping(ctx)
		if ping != nil {
			r, e := ping.Result()
			if e != nil {
				logger.Error(e.Error())
				logger.Error("Redis connection check failed")
			} else {
				logger.Info("Redis 'ping' successfully responded as : " + r)
			}
		}
	}
}

// Set 设置
func Set(key string, value interface{}) {
	SetWithExpire(key, value, DefaultTtl)
}

// SetWithExpire 设置带过期
func SetWithExpire(key string, value interface{}, ttl int) {
	err := GetClient().Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		logger.Error(err.Error())
	}
}

// Get 获取
func Get(key string) string {
	res, err := GetClient().Get(ctx, key).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res
}

// SetExpire 设置过期时间
func SetExpire(key string, ttl int) bool {
	err := GetClient().Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	return true
}

// TryLock 分布式锁
func TryLock(key string, parse string, ttl int) bool {
	b, err := GetClient().SetNX(ctx, key, parse, time.Duration(ttl)*time.Second).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return b
}

// Unlock 解锁
func Unlock(key string, parse string) {
	script := redis.NewScript(`
		if redis.call('get', KEYS[1]) == ARGV[1] 
		then return redis.call('del', KEYS[1])
		else return 0 end;
	`)
	keys := []string{key}
	args := []interface{}{parse}
	err := script.Run(ctx, GetClient(), keys, args).Err()
	if err != nil {
		logger.Error(err.Error())
	}
}

// Incr 自增
func Incr(key string, count int64) int64 {
	res, err := client.IncrBy(ctx, key, count).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res
}
