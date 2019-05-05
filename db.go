package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, _ := redis.Dial("tcp", ":6379")
	defer func() { _ = conn.Close() }()
	conn.Do("set", "c1", "hello")
}
