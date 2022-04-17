package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var Pool *redis.Pool

func InitPool(address string,maxIdle,MaxActive int,IdleTimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,   //最大空闲数
		MaxActive:  MaxActive,   //表示和数据库的最大连接数，0为不限制
		IdleTimeout: IdleTimeout, //最大空闲时间
		Dial:func() (redis.Conn, error) { //初始化链接，链接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
}
