package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	// 连接到 Redis 服务器
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // password set
		DB:       0,        // use default DB
	})
	// 每个 Redis 命令都接受一个上下文，您可以使用它来设置超时或传播一些信息，例如跟踪上下文。
	ctx := context.Background()
	// 执行命令 get 命令
	result, _ := rdb.Get(ctx, "inactivity_timeout").Result()
	fmt.Println(result)

	// 或者，您可以保存命令并稍后分别访问值和错误：
	get := rdb.Get(ctx, "inactivity_timeout")
	fmt.Println(get.Val(), get.Err())

}
