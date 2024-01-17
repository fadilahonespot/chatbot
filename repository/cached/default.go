package cached

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func NewWrapper() CacheWrapper {
	fmt.Println("Connect Redis Client.....")

	addr, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(addr)
	err = client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return &cache{client: client}
}

func (w *cache) Set(ctx context.Context, key, value string, duration time.Duration) (err error) {
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", key, value)
	err = w.client.Set(ctx, key, value, duration).Err()
	return
}

func (w *cache) Get(ctx context.Context, key string) (value string, err error) {
	fmt.Printf("[CACHED GET] key: %v \n", key)
	err = w.client.Get(ctx, key).Scan(&value)
	return
}

func (w *cache) Delete(ctx context.Context, key string) (err error) {
	fmt.Printf("[CACHED DEL] key: %v \n", key)
	err = w.client.Del(ctx, key).Err()
	return
}
