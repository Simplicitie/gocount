package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Example useage: REDIS_HOST=localhost REDIS_PORT=6379 go run main.go
	// Example useage2: REDIS_HOST=123.123.456.3 REDIS_USER=bob REDIS_PASSWORD=secret_password REDIS_PORT=6379 go run main.go
	uname := os.Getenv("REDIS_USER") // blank username is valid in some redis cases
	pass := os.Getenv("REDIS_PASS")  // blank password is valid in some redis instances
	host := os.Getenv("REDIS_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("REDIS_PORT")
	if len(port) == 0 {
		port = "6379"
	}

	connection_string := fmt.Sprintf("redis://%s:%s@%s:%s/1", uname, pass, host, port)
	opt, err := redis.ParseURL(connection_string)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	ctx := context.Background()

	err = client.Set(ctx, "count", 0, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "count").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	count, err := client.Incr(ctx, "count").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", count)

}
