package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func testStrings() {
	// 1.set 设置值，并添加过期时间
	// set k v ex 5  设置kv值并添加5秒超时时间
	ctx := context.Background()
	result1, err := RDB.Set(ctx, "stringT1", "2024080701", 5*time.Second).Result()
	_, err = RDB.Set(ctx, "stringT2", "2024080708", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result1)
	
	// get 获取数据
	time.Sleep(2 * time.Second)
	result1, err = RDB.Get(ctx, "stringT1").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("stringT1 does not exist")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(result1)
	
	// get 过期时间到了之后再获取数据
	time.Sleep(5 * time.Second)
	result1, err = RDB.Get(ctx, "stringT1").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("stringT1 does not exist")
	} else if err != nil {
		panic(err)
	}
	
	// 2.获取某kv是否存在
	// exists k
	i, err := RDB.Exists(ctx, "stringT2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
	
	// 3.删除
	// del k
	del := RDB.Del(ctx, "stringT2")
	fmt.Println(del.String(), del.Val())
	
	// 4.SetNX， 设置并指定过期时间，仅当 key 不存在时候才设置有效
	// setnx k v ex 5
	err = RDB.SetNX(ctx, "setnx-key", "setnx-val", 0).Err()
	if err != nil {
		fmt.Println("setnx value failed: ", err)
		panic(err)
	}
	
	// 5.MSet 设置值
	// mset k1 v1 k2 v2
	err = RDB.MSet(ctx, "mset-key1", "mset-val1", "mset-key2", "mset-val2", "mset-key3", "mset-val3").Err()
	if err != nil {
		fmt.Println("MSet ERROR : ", err)
	}
	// MGet 获取值
	// mget k1 k2
	vals, err := RDB.MGet(ctx, "mset-key1", "mset-key2", "mset-key3").Result()
	if err != nil {
		fmt.Println("MGet ERROR: ", err)
		panic(err)
	}
	fmt.Println("vals: ", vals)
	
	// 6.原子操作 加 减
	err = RDB.SetNX(ctx, "nums", 2, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("set nums : ", 2)
	
	// Incr
	// incr k1
	val, err := RDB.Incr(ctx, "nums").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("incr: ", val)
	
	// IncrBy
	// incrby k1 5
	val, err = RDB.IncrBy(ctx, "nums", 10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("incrby: ", val)
	
	// Decr
	val, _ = RDB.Decr(ctx, "nums").Result()
	fmt.Println("desc: ", val)
	
	// DecrBy
	val, _ = RDB.DecrBy(ctx, "nums", 5).Result()
	fmt.Println("decrby: ", val)
	
	// 7.设置过期时间
	RDB.Set(ctx, "setkey-expire-1", "value-expire-1", 0).Err()
	RDB.Set(ctx, "setkey-expire-2", "value-expire-2", time.Second*40).Err()
	
	// Expire, 设置key在某个时间段后过期
	// expire k1 5
	val1, _ := RDB.Expire(ctx, "setkey-expire-1", time.Second*20).Result()
	fmt.Println("expire: ", val1)
	
	// ExpireAt，设置key在某个时间点后过期
	val2, _ := RDB.ExpireAt(ctx, "setkey-expire-2", time.Now().Add(time.Second*50)).Result()
	fmt.Println("expire at: ", val2)
	
	// TTL
	expire, err := RDB.TTL(ctx, "setkey-expire-1").Result()
	fmt.Println(expire, err)
	
}
