package main

import (
	"context"
	"time"
)

func withTimeoutSample() {
	// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	//	return WithDeadline(parent, time.Now().Add(timeout))
	// }
	// 底层使用WithDeadline 方法，
	// WithDeadline 接收一个具体的时间点作为截止时间
	// WithTimeout 接收一个相对时间段作为超时时间
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	_, _ = ctx.Deadline()
}
