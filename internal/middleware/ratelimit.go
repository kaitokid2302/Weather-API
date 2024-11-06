package middleware

import (
	"time"
	lamredis "weather-api/internal/redis"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := lamredis.Client
		ctx := c.Request.Context()

		_, er := client.Get(ctx, "ratelimit").Result()
		if er == redis.Nil {
			c.Next()
			// set with expire 10 second
			client.SetEx(ctx, "ratelimit", true, time.Second*10)
			return
		} else if er != nil {
			panic(er)
		} else {
			c.AbortWithStatusJSON(429, gin.H{"error": "Rate limit, access later"})
		}
	}
}
