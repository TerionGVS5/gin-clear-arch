package views

import (
	"gin-clear-arch/anime/adapters"
	"gin-clear-arch/anime/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnimeList(c *gin.Context) {
	useCase := services.AnimeListUseCase{
		Repo: adapters.MemoryStorage{},
	}
	c.IndentedJSON(http.StatusOK, useCase.GetAnimeList())
}
