package store

type Song struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Singer      Artist `json:"singer"`
	Author      Artist `json:"author"`
	ReleaseDate string `json:"release_date"`
}

type Artist struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Nation string `json:"nation"`
}

func GetSongByID(id int) (*Song, error) {
	artist := Artist{
		ID:     1,
		Name:   "Adel",
		Age:    45,
		Nation: "US",
	}
	return &Song{
		ID:          12,
		Name:        "hello",
		Singer:      artist,
		Author:      artist,
		ReleaseDate: "2020-01-02",
	}, nil
}

func GetSongByName(name string) (*[]Song, error) {}
