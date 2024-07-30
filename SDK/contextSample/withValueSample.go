package main

import (
	"context"
	"fmt"
	"time"
)

func withValueSample() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	
	// 可以配合其他方法一起使用，传参需要其他方法返回的context对象
	ctx = context.WithValue(ctx, "user", "XiaoMing")
	
	go pushData(ctx)
	time.Sleep(5 * time.Second)
}

func pushData(ctx context.Context) {
	value := ctx.Value("user")
	fmt.Println(value)
	
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), "  ", "push context deadline exceeded")
			return
		default:
			fmt.Println("push context running")
			time.Sleep(1 * time.Second)
		}
	}
}
