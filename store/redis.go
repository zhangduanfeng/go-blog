package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

var Ctx = context.Background()

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/12/18 15:19
 **/
func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "124.221.221.82:6379",
		Password: "blog112233", // no password set
		DB:       0,            // use default DB
		PoolSize: 10,
	})
	result := RedisClient.Ping(Ctx)
	if result.Val() != "PONG" {
		logrus.Error("redis连接失败!")
	} else {
		logrus.Info("redis连接成功!")
	}
}
