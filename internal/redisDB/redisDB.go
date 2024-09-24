package redisDB

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ConnectRedisDB() *redis.Client {
	options := &redis.Options{
		Addr: viper.GetString("redis.uri"),
	}

	client := redis.NewClient(options)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("Connected to Redis!")
	return client
}
