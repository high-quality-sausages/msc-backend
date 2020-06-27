package repository

import (
	"database/sql"
	"fmt"

	"github.com/high-quality-sausages/msc-backend/drivers/postgres"
	"github.com/high-quality-sausages/msc-backend/models"
)

var db *sql.DB

func init() {
	db = postgres.DBConn()
}

type RequestResp struct {
	Data []*models.Song `json:"data"`
	Err  string         `json:"error"`
}

// GetSongByID : 通过ID查询
func GetSongByID(id string) ([]*models.Song, error) {

	song := &models.Song{}
	stmt, err := db.Prepare("SELECT * FROM tbl_songs WHERE id=$1 AND 1=1 limit 1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var singerID string
		var authorID string

		err := rows.Scan(&song.ID, &song.Name, &singerID, &authorID, &song.Picture, &song.Path, &song.Link, &song.ReleaseDate)
		if err != nil {
			return nil, err
		}

		fmt.Println(singerID)
		fmt.Println(authorID)

		singer, err := GetArtist(singerID)
		if err != nil {
			return nil, err
		}

		song.Singer = singer

		author, err := GetArtist(authorID)
		if err != nil {
			return nil, err
		}

		song.Author = author
	}

	songs := []*models.Song{}
	songs = append(songs, song)

	return songs, nil

}

func GetArtist(userID string) (*models.Artist, error) {
	artist := &models.Artist{}
	stmt, err := db.Prepare("SELECT * FROM tbl_artist WHERE id=$1 AND 1=1 limit 30")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&artist.ID, &artist.Name, &artist.Age, &artist.Nation, &artist.Picture)
		if err != nil {
			return nil, err
		}
	}

	return artist, nil
}

// GetSongByName : 通过歌名模糊查询歌曲
func GetSongsByName(name string) ([]*models.Song, error) {
	songSet := []*models.Song{}
	stmt, err := db.Prepare("SELECT * FROM tbl_songs WHERE name=$1 AND 1=1 limit 30")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		song := &models.Song{}
		singerID := ""
		authorID := ""
		if err = rows.Scan(&song.ID, &song.Name, &singerID,
			&authorID, &song.Picture, &song.Path, &song.Link, &song.ReleaseDate); err != nil {
			return nil, err
		}

		// singer, err := GetArtist(singerID)
		// if err != nil {
		// 	return nil, err
		// }
		// song.Singer = singer

		// author, err := GetArtist(authorID)
		// if err != nil {
		// 	return nil, err
		// }

		// song.Author = author

		songSet = append(songSet, song)

	}
	return songSet, nil
}

// GetSongBySingerName : 通过歌手名模糊查询歌曲
func GetSongsBySinger(singerName string) ([]*models.Song, error) {
	// artist := Artist{
	// 	ID:     "1",
	// 	Name:   "Adel",
	// 	Age:    45,
	// 	Nation: "US",
	// }

	result := make([]*models.Song, 0)
	song := &models.Song{
		ID:          "12",
		Name:        "hello",
		ReleaseDate: "2020-01-02",
	}

	result = append(result, song)

	return result, nil
}
