package album

type Album struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Artist string  `db:"artist"`
	Price  float64 `db:"price"`
}

type AlbumDTO struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type NewAlbumDTO struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
