package main

import (
	"gin-clear-arch/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/anime/", api.GetAnimeList)
	err := router.Run("0.0.0.0:8000")
	if err != nil {
		return
	}
}
