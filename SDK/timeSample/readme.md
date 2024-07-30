## time
### 一、定时器
#### 1.1 timer

timer 只能响应一次
```
// 创建定时器
timer := time.NewTimer(2 * time.Second)

// 定时结束后，结束等待
t2 := <-timer.C
fmt.Printf("t2:%v\n", t2)

// 停止定时器
// 正常执行stop方法会返回true；如果到期之后再执行stop方法停止定时器会返回false
b := timer.Stop()

// 重置定时器
timer.Reset(1 * time.Second)

// 定时器延时
<-time.After(2 * time.Second)
```

#### 1.2 ticker
ticker 周期响应

```
ticker := time.NewTicker(1 * time.Second)
i := 0
// 子协程
go func() {
    for {
        i++
        fmt.Println(<-ticker.C)
        if i == 5 {
            // 停止
            ticker.Stop()
        }
    }
}()
```






















