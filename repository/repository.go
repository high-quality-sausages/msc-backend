package repository

import "github.com/high-quality-sausages/msc-backend/models"

type SongHandler interface {
	QueryByID(ID string) (*models.Song, error)
	QueryByName(name string) ([]*models.Song, error)
	QueryBySinger(singerName string) ([]*models.Song, error)
}

// type SongListHandler interface {
// 	SortByName(models.SongList, error)
// 	SortByReleaseDate(models.SongList, error)
// }
