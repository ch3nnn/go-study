package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // password set
		DB:       0,        // use default DB
	})
	ctx := context.Background()
	// 获取 keys *
	iterator := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iterator.Next(ctx) {
		fmt.Println("keys", iterator.Val())
	}

	if err := iterator.Err(); err != nil {
		println(err)
	}

	//集合和哈希
	//迭代集合元素：
	//iter := rdb.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
	//哈希：
	//iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
	//iter := rdb.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator()

}
