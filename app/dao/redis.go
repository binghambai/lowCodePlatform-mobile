package dao

import (
	"binghambai.com/lowCodePlatform-mobile/app/conn"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

var redisConn = conn.Rd

func RedisGet(key string) *string {
	val, err := redisConn.Get(ctx, key).Result()
	if err != nil {
		panic("connect to redis has error")
	} else if err == redis.Nil {
		log.Println("当前key不存在")
		return nil
	}
	return &val
}

func RedisExist(key string) bool {
	_, err := redisConn.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	return true
}

func RedisSet(key, val string, expire time.Duration) {
	if _, err := redisConn.Set(ctx, key, val, expire).Result(); err != nil {
		panic("can not set key-value to redis")
	}
}
