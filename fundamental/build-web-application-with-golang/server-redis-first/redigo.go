package main

import (
	"fmt"
	"strconv"
	"github.com/gomodule/redigo/redis"
)

var PORT_REDIGO = 7079

func execRedigo() {
	fmt.Println("START: use redigo/redis")
	conn, err := redis.Dial("tcp", "localhost:" + strconv.Itoa(PORT_REDIGO))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r, err := conn.Do("SET", "temperature", "25")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	s, err := redis.String(conn.Do("GET", "temperature"))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println("START: use redigo/redis")
}
