package myRedis

import (
	"commons/tools"

	"github.com/go-redis/redis"
)

type RedisConnect struct {
	redisOptions *redis.Options
	redisClient  *redis.Client
}

func (This *RedisConnect) GetRedisClient() *redis.Client {
	return This.redisClient
}

func (This *RedisConnect) GetRedisOption() *redis.Options {
	return This.redisOptions
}

// NewRedisClient 可变参数顺序：RedisHost, Port, Password, DbIndex string
func NewRedisClient(redisOption ...string) *RedisConnect {
	opt := &redis.Options{
		Addr:     redisOption[0] + ":" + redisOption[1],
		Password: redisOption[2],
		DB:       tools.StringToInt(redisOption[3]),
	}
	return &RedisConnect{redisOptions: opt, redisClient: redis.NewClient(opt)}
}
