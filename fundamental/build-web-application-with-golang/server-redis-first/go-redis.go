package main

import (
	"fmt"
	"strconv"
	"context"
	"github.com/go-redis/redis/v8"
)

// redis-server --port 7079
const PORT_GO_REDIS = 7079
var ctx = context.Background()

func ExampleNewClient() {
	fmt.Println("FUNC: ExampleNewClient()")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:" + strconv.Itoa(PORT_GO_REDIS),
		Password: "",
		DB: 0,
	})
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func ExampleClient() {
	fmt.Println("FUNC: ExampleClient()")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:" + strconv.Itoa(PORT_GO_REDIS),
		Password: "",
		DB: 0,
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func execGoRedis() {
	fmt.Println("START: use go-redis/redis")
	ExampleNewClient()
	ExampleClient()
	fmt.Println("FINISH: use go-redis/redis")
}
