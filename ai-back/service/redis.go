package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

type RedisService interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type redisService struct {
	redisAddress  string
	redisPassword string
	redisDb       int
}

func NewRedisService() RedisService {
	return &redisService{
		redisAddress:  "localhost:6379",
		redisPassword: "",
		redisDb:       0,
	}
}

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "gpt-redis:6379",
	Password: "", // 没有密码，默认值
	DB:       0,  // 默认DB 0
})

func (r redisService) Set(key string, value string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println("redis set失败提示：key is :", key, "value is :", value, "   location is :service/redis.go  Set")
		fmt.Println("redis set error", err.Error(), "   location is :service/redis.go  Set")
		return err
	}

	err = rdb.Expire(ctx, key, 5*time.Minute).Err()
	return nil
}

func (r redisService) Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func (r redisService) Del(key string) error {
	return rdb.Del(ctx, key).Err()
}
