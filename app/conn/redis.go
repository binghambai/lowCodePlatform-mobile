package conn

import "github.com/redis/go-redis/v9"

var Rd *redis.Client

func initRedis() {
	Rd = redis.NewClient(&redis.Options{
		Addr:     "localhost:56379",
		Password: "redis",
	})
}
