package service

import (
	"awesomeProject3/api"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
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
		redisAddress:  api.RedisAddress,
		redisPassword: api.RedisPassword,
		redisDb:       api.RedisDb,
	}
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     api.RedisAddress,
	Password: api.RedisPassword, // no password set
	DB:       api.RedisDb,       // use default DB
})

func (r *redisService) Set(key string, value string) error {

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error().Msg("redis set失败提示：key is :" + key + "value is :" + value + "   location is :service/redis.go  Set")
		return err
	}

	err = rdb.Expire(ctx, key, 5*time.Minute).Err()
	return nil
}

func (r *redisService) Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func (r *redisService) Del(key string) error {
	return rdb.Del(ctx, key).Err()
}
