package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedisConnection() (*redis.Client, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(pong)

	err = client.Set(ctx, "token", "mYincrEdableToKen123", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	r, err := client.Get(ctx, "token").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Key: %s\n", r)

	return client, nil
}
