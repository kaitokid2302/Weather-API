package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"weather-api/internal/model"
	"weather-api/internal/redis"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	name := c.Param("place")
	fmt.Printf(">>> name: %v\n", name)
	apikey := os.Getenv("apikey")
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", name, apikey)
	client := redis.Client
	val, er := client.Get(c.Request.Context(), name).Result()

	if er != nil {
		// data not exist in redis
		response, er := http.Get(url)
		if er != nil {
			panic(er)
		}
		val, _ := io.ReadAll(response.Body)
		fmt.Printf(">>> val: %v\n", val)
		var data model.Weather
		json.Unmarshal([]byte(val), &data)
		c.JSON(200, data)
		client.SetEx(context.Background(), name, string(val), time.Hour*12)
		return
	} else {
		// data exist in redis
		var data model.Weather
		json.Unmarshal([]byte(val), &data)
		c.JSON(200, data)
		return
	}
}
