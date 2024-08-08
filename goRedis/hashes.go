package main

import (
	"context"
	"fmt"
)

/*
	1.HSET 单个设置值。

	2.HGET 单个获取值。

	3.HMSET 批量设置。

	4.HMGET 批量获取值。

	5.HGETALL 获取所有值。

	6.HDEL 删除字段，支持删除多个字段。

	7.HLEN 获取hash表中key的值数量。

	8.HEXISTS 判断元素是否存在。

	9.HINCRBY 根据key的field字段的整数值加减一个数值。

	10.HSETNX 如果某个字段不存在则设置该字段值。
*/

func testHash() {
	
	ctx := context.Background()
	// 1.HSET，根据key设置field字段值,可设置多个值,版本3以上可以用
	setResult, _ := RDB.HSet(ctx, "hash1", "stu1", "xiaoming").Result()
	fmt.Println(setResult)
	setResult, _ = RDB.HSet(ctx, "hash1", []string{"stu2", "zhangsan", "stu3", "erchui"}).Result()
	fmt.Println(setResult)
	setResult, _ = RDB.HSet(ctx, "hash1", map[string]interface{}{"stu4": "lisi", "stu5": "wangwu", "stu6": 666}).Result()
	fmt.Println(setResult)
	
	// 同hset 使用方式一致，版本3后不再使用，但未弃用
	b2, _ := RDB.HMSet(ctx, "hash1", "stu6", "xiaoming", "stu7", "xiongda").Result()
	fmt.Println(b2)
	
	// 某字段不存在时设置该字段
	result, _ := RDB.HSetNX(ctx, "hash1", "stu1", "xiaoming2hao").Result()
	fmt.Println(result)
	
	// 2.HGET 获取值
	getResult, _ := RDB.HGet(ctx, "hash1", "stu6").Result()
	fmt.Println(getResult)
	
	getResults, _ := RDB.HMGet(ctx, "hash1", "stu6", "stu7", "stu8").Result()
	fmt.Println(getResults)
	
	getAllResult, err := RDB.HGetAll(ctx, "hash1").Result()
	fmt.Println(getAllResult, err)
	
	// 3.HDEL 删除字段
	b, _ := RDB.HDel(ctx, "hash1", "stu1", "stu2").Result()
	fmt.Println(b)
	
	// 4.HLen
	i, _ := RDB.HLen(ctx, "hash1").Result()
	fmt.Println(i)
	
	// HEXISTX，某个hashkey中字段field否存在
	ok, _ := RDB.HExists(ctx, "hash1", "stu8").Result()
	fmt.Println("HExists: ", ok) // HExists: true
	
	// HIncrBy，根据key的field字段的整数值加减一个数值
	stu6, _ := RDB.HIncrBy(ctx, "hash1", "stu6", -3).Result()
	fmt.Println("HIncrBy : ", stu6) // HIncrBy :  20
}
