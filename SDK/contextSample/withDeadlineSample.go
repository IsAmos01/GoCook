package main

import (
	"context"
	"fmt"
	"time"
)

func withDeadlineSample() {
	// 5秒后会停止子协程
	deadTime := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadTime)
	
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, "   ", ok)
	
	go drink(ctx)
	go eat(ctx)
	
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err(), "     ", "main context deadline exceeded")
				// break只能退出select，不能退出for循环，因此使用return停止或者使用标签
				return
			default:
				fmt.Println("main running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	
	time.Sleep(5 * time.Second)
	// deadtime 时间内想提前停止子协程可以通过执行cancel方法
	cancel()
}

func drink(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), "     ", "drink context deadline exceeded")
			// break只能退出select，不能退出for循环，因此使用return停止或者使用标签
			return
		default:
			fmt.Println("drink  running")
			time.Sleep(1 * time.Second)
		}
	}
}

func eat(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), "     ", "eat context deadline exceeded")
			// break只能退出select，不能退出for循环，因此使用return停止或者使用标签
			return
		default:
			fmt.Println("eat running")
			time.Sleep(1 * time.Second)
		}
	}
}
