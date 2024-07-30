### Context
#### 一、context的概念及作用
```text
在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。
```

#### 二、context接口
```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

其中:

+ **Deadline方法** 需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）;
+ **Done方法** 需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
+ **Err方法** 会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
  + 如果当前Context被取消就会返回Canceled错误；
  + 如果当前Context超时就会返回DeadlineExceeded错误；
+ **Value方法** 会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；


emptyCtx 实现了context接口

```
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (*emptyCtx) Done() <-chan struct{} { return nil }
func (*emptyCtx) Err() error { return nil }
func (*emptyCtx) Value(key any) any { return nil }
```

通过Background()和TODO()方法获取空的context

```
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context { return background }

func TODO() Context { return todo }
```


方法

```
WithCancel()
WithDeadline()
WithTimeout()  // 内部调用WithDeadline()方法实现其功能
WithValue()
```


cancelCtx 实现了context接口，当使用WithCancel()方法时会返回该结构体并调用其方法
```
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     atomic.Value          // of chan struct{}, created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}

func (c *cancelCtx) Value(key any) any 
func (c *cancelCtx) Done() <-chan struct{} 
func (c *cancelCtx) Err()
func (c *cancelCtx) String() string
func (c *cancelCtx) cancel

```

timerCtx实现了context接口，WithDeadline()方法时会返回该结构体并调用其方法
```
type timerCtx struct {
	cancelCtx
	timer *time.Timer // Under cancelCtx.mu.

	deadline time.Time
}

func (c *timerCtx) Deadline()
func (c *timerCtx) String() string
func (c *timerCtx) cancel
```

valueCtx实现了context接口，WithValue()方法时会返回该结构体并调用其方法
```
type valueCtx struct {
	Context
	key, val any
}

func (c *valueCtx) String() string
func (c *valueCtx) Value(key any) any
```




















