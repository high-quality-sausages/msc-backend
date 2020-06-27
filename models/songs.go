package models

type Song struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Singer      *Artist `json:"singer_id"`
	Author      *Artist `json:"author_id"`
	Picture     string  `json:"picture"`
	Path        string  `json:"path"`
	Link        string  `json:"link"`
	ReleaseDate string  `json:"release_date"`
}

// type SongList struct {
// 	songs []*Song
// }
