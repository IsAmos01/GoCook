package main

import (
	"context"
	"fmt"
)

/*
1.LPUSH：list头部(左边)插入值，最后的值在最前面。LPUSHX 仅当列表值存在时才插入值

2.LPOP：移除列表的头部值并返回这个值

3.RPUSH：list尾部(右边)插入值。RPUSHX 仅当列表值存在才插入值

4.RPOP：移除列表的尾部值并返回这个值

5.LRANGE：返回key列表指定区间的值

6.BLPOP: 语法 BLPOP key [key ...] timeout，从 key 列表头部弹出一个值，没有就阻塞 timeout 秒，如果 timeout=0 则一直阻塞

7.BRPOP：与上面 BLPOP 用法相似，只不过 BRPOP 是从尾部弹出一个值

8.LLEN：返回列表的长度

9.LINSERT：在指定位置插入数据

10.LREM：删除列表中的数据

11.LINDEX：根据索引查询列表中的值

12.LSET：根据索引设置列表中的某个值
*/

func testLists() {
	ctx := context.Background()
	// 1.插值  LPush
	result, err := RDB.LPush(ctx, "list1", 1, 2, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// 插值 RPush
	result, err = RDB.RPush(ctx, "list1", 7, 8, 9).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	
	// 2.取值 RPop LPop
	popL := RDB.LPop(ctx, "list1")
	fmt.Println(popL)
	popR := RDB.RPop(ctx, "list1")
	fmt.Println(popR)
	
	// 3.返回指定区间（闭区间）的值. 其中 0 -1 返回所有的值
	value := RDB.LRange(ctx, "list1", 0, -1).Val()
	fmt.Println(value)
	
	// 4.LInsert 在指定位置插入数据,从左往右第一个匹配的value根据要求插入
	err = RDB.LInsert(ctx, "list1", "before", "1", 10).Err()
	err = RDB.LInsert(ctx, "list1", "after", "1", 10).Err()
	if err != nil {
		panic(err)
	}
	
	// 5.LREM 删除列表中的数据
	RDB.LRem(ctx, "list1", 3, 10) // 从列表左边开始删除值 10，出现重复元素只删除前3个
	RDB.LRem(ctx, "list1", -3, 6) // 从列表尾部(右边)开始删除值 6，出现重复元素只删除前3个
	
	// 6.LIndex 根据索引查询值，索引是从0开始
	val, _ := RDB.LIndex(ctx, "list1", 3).Result()
	fmt.Println("LIndex val: ", val)
	
	// LSET 根据索引设置某个值，索引从0开始
	val, _ = RDB.LSet(ctx, "list1", 3, 10).Result()
	fmt.Println("lset: ", val)
}
