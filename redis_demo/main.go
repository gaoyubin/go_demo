package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
		PoolSize: 50,
	})
	fmt.Println(client)
	result, err := client.Ping().Result()
	if err != nil {
		fmt.Println(result, err)
		panic(err)
	}
	fmt.Println(result, err)
	//client.SAdd()
	//fmt.Println(result)
	//val, err := client.Get("id_2").Result()
	// 检测，查询是否出错
	//if err != redis.Nil {
	//
	//}
	//fmt.Println("key", val)

	str_list := []interface{}{"1", "3", "hello", "world"}
	//str_list := []int{1, 3, 4, 5}
	err = client.SAdd("setkey:1", str_list...).Err()
	if err != nil {
		panic(err)
	}

	// SMEMBERS，获取集合中的所有元素数据
	smembers, err := client.SMembers("setkey:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("SMembers，setkey:1: ", smembers)

}

//func main() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:        "localhost:6379",
//		Password:    "",
//		DB:          0,
//		IdleTimeout: 350,
//		PoolSize:    50, // 连接池连接数量
//	})
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
//	defer cancel()
//	_, err := rdb.Ping(ctx).Result() // 检查连接redis是否成功
//	if err != nil {
//		fmt.Println("Connect Failed: %v \n", err)
//		panic(err)
//	}
//
//	ctx = context.Background()
//	// 设置 key 的值，0 表示永不过期
//	err = rdb.Set(ctx, "setkey-1", "value-1", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 设置 key 的值的过期时间为 30 秒
//	err = rdb.Set(ctx, "setkey-2", "value-2", time.Second*30).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 获取key的值
//	val, err := rdb.Get(ctx, "setkey-1").Result()
//	if err == redis.Nil { // 如果返回 redis.Nil 说明key不存在
//		fmt.Println("key not exixt")
//	} else if err != nil {
//		fmt.Println("Get Val error: ", err)
//		panic(err)
//	}
//	fmt.Println("Get Val: ", val)
//
//	val, _ = rdb.Get(ctx, "setkey-2").Result()
//	fmt.Println("Get Val setkey-2: ", val)
//}
