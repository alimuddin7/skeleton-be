package redis

import (
	"context"
	"fmt"
	"test/configs"
	"test/helpers"
	hModels "test/helpers/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type (
	redisDatabase struct {
		Client *redis.Client
		Logs   zerolog.Logger
	}
	RedisDatabase interface {
		GetClient() *redis.Client
		HealthCheck(ctx context.Context) hModels.DataHealthCheck
	}
)

func InitializeRedisDatabase(conn *redis.Client, log zerolog.Logger) RedisDatabase {
	return &redisDatabase{
		Client: conn,
		Logs:   log,
	}
}

func ConnectRedis(log zerolog.Logger) *redis.Client {
	conf := configs.Cfg.Database.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Address, conf.Port),
		Password: conf.Password,
		DB:       conf.DBType,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Error().Err(err).Msg("Error open redis connection")
		panic("Error open redis connection")
	}

	log.Info().Msg("Redis connected successfully")
	return client
}

func (r *redisDatabase) GetClient() *redis.Client {
	return r.Client
}

func (r *redisDatabase) HealthCheck(ctx context.Context) hModels.DataHealthCheck {
	res := hModels.DataHealthCheck{
		ServiceName: "Redis",
		StatusCode:  200,
	}
	
	if _, err := r.Client.Ping(ctx).Result(); err != nil {
		r.Logs.Error().Ctx(ctx).Err(err).Msg("Redis ping failed")
		res.StatusCode = 500
		res.AdditionalData = err.Error()
		return res
	}
	
	return res
}
