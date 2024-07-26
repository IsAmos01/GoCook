package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var wg sync.WaitGroup
var GOTIME = "2006-01-02 15:04:05"

var (
	once      sync.Once
	initValue int
)

func initialize() {
	fmt.Println("Initializing...")
	initValue = 42
}
func doSomething() {
	once.Do(initialize)
	fmt.Println("Value:", initValue)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doSomething()
		}()
	}
	wg.Wait()
	
	// 等待
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	
	// 互斥锁
	var lock sync.Mutex
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go helloLock(&lock, &wg)
	}
	wg.Wait()
	
	// 读写锁
	var rwlock sync.RWMutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go helloRLock(&rwlock, &wg)
		if i%2 == 0 {
			wg.Add(1)
			go helloWLock(&rwlock, &wg)
		}
	}
	wg.Wait()
	
	for i := 0; i < 10; i++ {
		// i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
	
	// sync.map
	dict := sync.Map{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dict.Store(i, i)
		}(i)
	}
	wg.Wait()
	dict.Range(func(key, value interface{}) bool {
		fmt.Println(value.(int))
		return true
	})
	
	// map
	// dict := make(map[int]int)
	// var lock sync.Mutex
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		defer lock.Unlock()
	// 		lock.Lock()
	// 		dict[i] = i
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println(dict)
}

func helloRLock(lock *sync.RWMutex, wg *sync.WaitGroup) {
	defer fmt.Println("RELEASE RRR unlock ", count, "  ", time.Now().Format(GOTIME))
	defer lock.RUnlock()
	defer wg.Done()
	lock.RLock()
	time.Sleep(time.Second)
	fmt.Println("LOCK RRR lock ", count, "  ", time.Now().Format(GOTIME))
}

func helloWLock(lock *sync.RWMutex, wg *sync.WaitGroup) {
	defer fmt.Println("RELEASE WWW unlock ", count, "  ", time.Now().Format(GOTIME))
	defer lock.Unlock()
	defer wg.Done()
	lock.Lock()
	count += 1
	fmt.Println("LOCK WWW lock ", count, "  ", time.Now().Format(GOTIME))
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello world")
}

func helloLock(lock *sync.Mutex, wg *sync.WaitGroup) {
	defer lock.Unlock()
	defer wg.Done()
	lock.Lock()
	count += 1
	fmt.Printf("hello lock %d\n", count)
}
