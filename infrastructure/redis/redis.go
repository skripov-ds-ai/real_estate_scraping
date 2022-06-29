package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"os"
	"strconv"
)

var Client *redis.Client

func (r *RedisConfig) Load() {
	r.Addr = os.Getenv("Addr")
	r.Password = os.Getenv("Password")
	DB, err := strconv.ParseInt(os.Getenv("DB"), 10, 32)
	if err != nil {
		panic(err)
	}
	r.DB = int(DB)
}

func RedisConnect(ctx context.Context) {
	Config.Load()
	options := redis.Options{
		Addr:     Config.Addr,
		Password: Config.Password,
		DB:       Config.DB,
	}
	Client = redis.NewClient(&options)

	if _, err := Client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
}

func RedisDisconnect() {
	_ = Client.Close()
}
