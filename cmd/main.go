package main

import (
	"gin-clear-arch/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/anime/", api.GetAnimeList)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
