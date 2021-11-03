package main

import (
	"gin-clear-arch/anime/views"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/anime/", views.GetAnimeList)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
