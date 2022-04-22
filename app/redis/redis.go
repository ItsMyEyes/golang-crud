package redis

import (
	"context"
	"crud_v2/app/enviroment"
	"fmt"
	"time"

	rds "github.com/go-redis/redis/v8"
)

func init() {
	connect()
}

var Client *rds.Client

// create connection to redis
func connect() *rds.Client {
	Client = rds.NewClient(&rds.Options{
		Addr:     enviroment.Get("REDIS_HOST"),
		Password: enviroment.Get("REDIS_PASSWORD"), // no password set
		DB:       0,                                // use default DB
	})

	pong, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, "connected to rds")
	return Client
}

func SetKey(key string, i interface{}, ttl time.Duration) *rds.StringCmd {
	if ttl == 0 {
		Client.Set(context.Background(), key, i, 0)
	} else {
		Client.Set(context.Background(), key, i, ttl)
	}

	return Client.Get(context.Background(), key)
}

func GetKey(key string) *rds.StringCmd {
	return Client.Get(context.Background(), key)
}

func RemoveKey(key string) bool {
	Client.Del(context.Background(), key)
	return true
}
