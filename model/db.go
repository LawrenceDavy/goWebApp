package model

import (
	"github.com/go-redis/redis"
)

//database
var client *redis.Client

// init the databse
func Init() {
	//connect to redis server
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}