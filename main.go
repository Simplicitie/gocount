package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx context.Context

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

	// establish connection to the Redis instence
	connection_string := fmt.Sprintf("redis://%s:%s@%s:%s/1", uname, pass, host, port)
	opt, err := redis.ParseURL(connection_string)
	if err != nil {
		panic(err)
	}
	client = redis.NewClient(opt)

	ctx = context.Background()

	// Start the Gin web Server
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to my counter!\n  To check the current count navigate to /count\n  To add one to the number, navigate to /incr\n  To reset the number to 0, navigate to /reset")
	})
	r.GET("/incr", incrCountByOne)
	r.GET("/count", getCount)
	r.GET("/reset", resetCount)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getCount(c *gin.Context) {
	val, err := client.Get(ctx, "count").Result()
	if err != nil {
		c.Error(err)
	}
	if val == "" {
		c.String(200, fmt.Sprintln("Current count is blank/nil!\n  Start counting by navigating to /incr."))
	} else {
		c.String(200, fmt.Sprintln("Current count:", val))
	}
}

func incrCountByOne(c *gin.Context) {
	count, err := client.Incr(ctx, "count").Result()
	if err != nil {
		c.Error(err)
	}
	c.String(200, fmt.Sprintln("Count incremented by one and is now:", count))
}

func resetCount(c *gin.Context) {
	err := client.Set(ctx, "count", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	c.String(200, fmt.Sprintln("Count reset to 0"))
}
