package users

import (
	"fmt"

	"github.com/high-quality-sausages/msc-backend/crawler"
	repo "github.com/high-quality-sausages/msc-backend/repository"
	"github.com/high-quality-sausages/msc-backend/songs"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   int    `json:"gender"`
	Role     int    `json:"role"` // 0: Admin; 1: Normal; 2: Visitor
}

type Handler interface {
	Update() (*User, error)
	Delete() error

	QuerySongByID(ID string) ([]songs.Song, error)              // 通过id查找歌曲
	QuerySongsByName(name string) ([]songs.Song, error)         // 通过歌名查找歌曲
	QuerySongsBySinger(singerName string) ([]songs.Song, error) // 通过歌手查找歌曲

	QuerySingerByID(singerID string)
	QuerySingerByName(singerName string)
	QuerySingerByNation(nation string)
}

func (user *User) QuerySongByID(ID string) ([]songs.Song, error) {
	songSet, err := repo.GetSongByID(ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return songSet, nil
}

func (user *User) QuerySongsByName(name string) ([]songs.Song, error) {
	songSet, err := repo.GetSongsByName(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(songSet) == 0 {
		if songSet, err = crawler.RequestSongsByName(name); err != nil {
			fmt.Println(err)
			return nil, err
		}
		return songSet, nil
	}
	return songSet, nil
}

func (user *User) QuerySongsBySinger(singerName string) ([]songs.Song, error) {
	songSet, err := repo.GetSongsBySinger(singerName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return songSet, nil
}
