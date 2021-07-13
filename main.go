package main

import (
	"os"
	"time"

	"al.api/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "3000"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           90 * time.Hour,
	}))
	router.GET("/", service.Base)
	router.POST("/check-in", service.Pouch)

	router.Run("0.0.0.0:" + port)
}
