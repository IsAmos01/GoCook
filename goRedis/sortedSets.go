package main

/*
ZADD：ZADD key score member [[score member] [score member] …]，将一个或多个member元素以及score加入到有序集合key中
ZSCORE：ZSCORE key member，返回集合 key 中 member 成员的分数
ZRANGE：ZRANGE key start stop [WITHSCORES]，返回集合 key 中指定区间的元素，score 从小到大排序，start 和 stop 都是 0 开始。
ZREVRANGE：ZREVRANGE key start stop [WITHSCORES]，与 zrange 相反，返回集合 key 中指定区间元素，score 从大到小排序。
ZRANGEBYSCORE：ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]，返回结果的数量区间，score 从小到大排序，LIMIT 参数指定返回结果集的数量和区间，后面可选的 [limit offset count] 像 SQL 中的 select ... limit offset,count。
ZREVRANGEBYSCORE：与上面 ZRANGEBYSCORE 几乎相同，不同是 score 是从大到小排序
ZREVRANGEBYSCOREWITHSCORES：和 ZRANGEBYSCORE 一样，区别是它不仅返回集合元素，也返回元素对应分数
ZREM：删除元素
ZREMRANGEBYRank：根据索引范围删除
ZREMRANGEBYSCORE：根据分数区间删除
*/
