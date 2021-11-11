package anime_storage

import "gin-clear-arch/internal/app/anime"

type MemoryStorage struct {
}

func (m MemoryStorage) GetAnimeList() []anime_service.Anime {
	return []anime_service.Anime{
		{
			Id:    1,
			Title: "Бездарная Нана",
			Year:  2020,
		},
	}
}
