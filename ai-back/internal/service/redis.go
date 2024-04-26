package service

import (
	"awesomeProject3/config"
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
		redisAddress:  config.GetConfig().Redis.RedisAddress,
		redisPassword: config.GetConfig().Redis.RedisPassword,
		redisDb:       config.GetConfig().Redis.RedisDb,
	}
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     config.GetConfig().Redis.RedisAddress,
	Password: config.GetConfig().Redis.RedisPassword, // no password set
	DB:       config.GetConfig().Redis.RedisDb,       // use default DB
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
