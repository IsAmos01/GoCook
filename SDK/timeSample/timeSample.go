package main

import (
	"fmt"
	"time"
)

func timeSample() {
	// 1.timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
	
	// 2.验证timer只能响应1次
	// timer2 := time.NewTimer(1 * time.Second)
	// for {
	// 	fmt.Println(<-timer2.C)
	// 	fmt.Println("时间到")
	// }
	
	// 2.timer被关闭后再读取会报错,到期之后stop方法会返回false
	// timer2 := time.NewTimer(1 * time.Second)
	// time.Sleep(2 * time.Second)
	// b := timer2.Stop()
	// if !b {
	// 	fmt.Println(<-timer2.C)
	// }
	// for {
	// 	fmt.Println(<-timer2.C)
	// 	fmt.Println("running")
	// }
	
	// 3.timer实现延时的功能
	// // (1)
	// time.Sleep(time.Second)
	// // (2)
	// timer3 := time.NewTimer(2 * time.Second)
	// <-timer3.C
	// fmt.Println("2秒到")
	// // (3)
	// <-time.After(2 * time.Second)
	// fmt.Println("2秒到")
	
	// 4.停止定时器
	// timer4 := time.NewTimer(2 * time.Second)
	// go func() {
	// 	<-timer4.C
	// 	fmt.Println("定时器执行了")
	// }()
	// b := timer4.Stop()
	// if b {
	// 	fmt.Println("timer4已经关闭")
	// }
	
	// 5.重置定时器
	// timer5 := time.NewTimer(3 * time.Second)
	// timer5.Reset(1 * time.Second)
	// fmt.Println(time.Now())
	// fmt.Println(<-timer5.C)
	
	// ticker := time.NewTicker(1 * time.Second)
	// i := 0
	// // 子协程
	// go func() {
	// 	for {
	// 		// <-ticker.C
	// 		i++
	// 		fmt.Println(<-ticker.C)
	// 		if i == 5 {
	// 			// 停止
	// 			ticker.Stop()
	// 		}
	// 	}
	// }()
	//
	// time.Sleep(10 * time.Second)
}