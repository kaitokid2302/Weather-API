package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	godotenv.Load()
	redisAdress := os.Getenv("redisAdress")
	redisPort := os.Getenv("redisPort")
	fmt.Printf("redisAdress: %v\n", redisAdress)
	fmt.Printf("redisPort: %v\n", redisPort)

	address := fmt.Sprintf("%v:%v", redisAdress, redisPort)
	Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		fmt.Errorf("không thể kết nối đến Redis: %v", err)
	}
}
