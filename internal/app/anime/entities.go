package anime_service

type Anime struct {
	Id    int    `json:"id,int"`
	Title string `json:"title,string"`
	Year  int    `json:"year,int"`
}
