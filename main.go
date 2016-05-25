package main

import (
	"flag"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	RedisSchemeURL := flag.String("redis-url", "redis://@localhost:6379/0", "Url to connect to redis")

	flag.Parse()

	c, err := redis.DialURL(*RedisSchemeURL)
	if err != nil {
		// handle connection error
		fmt.Println("Could not connect to REDIS: ", err)
	}
	defer c.Close()
}
