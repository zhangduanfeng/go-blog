package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/12/18 15:19
 **/
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "124.221.221.82:6379",
		Password: "blog112233", // no password set
		DB:       0,            // use default DB
		PoolSize: 10,
	})
	result := rdb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val() != "PONG" {
		// 连接有问题
		return nil
	}
	return rdb
}
