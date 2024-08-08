package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// 自定义客户端
		ClientName: "goCook",
		Password:   "",
		DB:         0,
	})
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := RDB.Ping(ctx).Result() // 检查连接redis是否成功
	if err != nil {
		fmt.Printf("Connect Failed: %v \n\n", err)
		panic(err)
	}
}
