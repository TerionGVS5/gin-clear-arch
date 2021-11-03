package services

type AnimeStorage interface {
	GetAnimeList() []Anime
}

type AnimeListUseCase struct{
	Repo AnimeStorage
}

func (useCase AnimeListUseCase) GetAnimeList() []Anime {
	return useCase.Repo.GetAnimeList()
}

