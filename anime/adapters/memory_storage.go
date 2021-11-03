package adapters

import "gin-clear-arch/anime/services"

type MemoryStorage struct {
}

func (m MemoryStorage) GetAnimeList() []services.Anime {
	return []services.Anime{
		{
			Id:    1,
			Title: "Бездарная Нана",
			Year:  2020,
		},
	}
}
