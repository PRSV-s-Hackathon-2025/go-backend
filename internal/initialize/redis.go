package initialize

import (
	"context"
	"fmt"

	"github.com/PRSV-Hackathon-2025/go-backend/global"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	redisConfig := global.Setting.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,  
		PoolSize: redisConfig.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Errorf("Redis connection failed", err)
		panic(err)
	}

	fmt.Print("Init Redis success")
	global.Redis = rdb
}
