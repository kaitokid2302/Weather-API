package main

import (
	"weather-api/internal/handler"
	"weather-api/internal/middleware"

	_ "weather-api/internal/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	server.GET("/:place", middleware.RateLimit(), handler.GetWeather)

	server.Run()
}
