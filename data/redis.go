package data

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func SetupRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-master.cloud-ml-pipeline.svc.cluster.local:6379",
		Password: "zWpuJLoN6d", // no password set
		DB:       0,            // use default DB
	})
	n, err := rdb.Exists(ctx, "id").Result()
	if err != nil {
		panic(err)
	}

	if n < 1 {
		err = rdb.Set(ctx, "id", 0, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func GetRedisData() (data []string) {
	keys := rdb.Keys(ctx, "*").Val()
	for _, key := range keys {
		if key == "id" {
			continue
		}
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		data = append(data, val)
	}
	return
}

func CreateRedisData() string {
	id, err := rdb.Get(ctx, "id").Result()
	if err != nil {
		panic(err)
	}

	svcName := "studio-" + id
	err = rdb.Set(ctx, id, svcName, 0).Err()
	if err != nil {
		panic(err)
	}

	rdb.Incr(ctx, "id")

	return svcName
}

func printRedisData() {
	keys := rdb.Keys(ctx, "*").Val()
	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", key, "val:", val)
	}
}
