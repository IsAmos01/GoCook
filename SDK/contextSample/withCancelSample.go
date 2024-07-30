package main

import (
	"context"
	"fmt"
	"time"
)

func withCancelSample() {
	// 通过withCancel方法 返回其中cancelCtx和cancel方法体，其中cancelCtx 实现了context接口
	ctx, cancel := context.WithCancel(context.Background())
	// 获取context的过期时间
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, "    ", ok)
	go handelrequest(ctx)
	
	time.Sleep(5 * time.Second)
	fmt.Println("it's time to stop all sub goroutines!")
	cancel()
	
	// just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

func handelrequest(ctx context.Context) {
	go writeredis(ctx)
	go writedatabase(ctx)
	
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handelrequest done.")
			return
		default:
			fmt.Println("handelrequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeredis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeredis done.")
			return
		default:
			fmt.Println("writeredis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writedatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writedatabase done.")
			return
		default:
			fmt.Println("writedatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
