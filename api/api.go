package api

import (
	"gin-clear-arch/internal/app/anime"
	"gin-clear-arch/internal/repositories/anime"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnimeList(c *gin.Context) {
	useCase := anime_service.AnimeListUseCase{
		Repo: anime_storage.MemoryStorage{},
	}
	c.IndentedJSON(http.StatusOK, useCase.GetAnimeList())
}
