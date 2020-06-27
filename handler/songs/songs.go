package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/high-quality-sausages/msc-backend/crawler"
	"github.com/high-quality-sausages/msc-backend/models"
	repo "github.com/high-quality-sausages/msc-backend/repository"
)

func QuerySongs(c *gin.Context) {
	latitude := c.Query("latitude")
	switch latitude {
	case "0":
		ID := c.Query("id")
		songs, err := QuerySongByID(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": "",
				"msg":  "Failed to query song by id",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": songs,
			"msg":  "",
		})

		return
	case "1":
		name := c.Query("name")
		songs, err := QuerySongsByName(name)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": "",
				"msg":  "Failed to query song by name",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": songs,
			"msg":  "",
		})
		return
	}
}

func QuerySongByID(ID string) ([]*models.Song, error) {
	song, err := repo.GetSongByID(ID)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func QuerySongsByName(name string) ([]*models.Song, error) {
	songs, err := repo.GetSongsByName(name)
	if err != nil {
		return nil, err
	}

	if len(songs) == 0 {
		songs, err := crawler.RequestSongsByName(name)
		if err != nil {
			return nil, err
		}

		return songs, nil
	}

	return songs, nil
}

func QuerySongsBySinger(singerName string) ([]*models.Song, error) {
	songs, err := repo.GetSongsBySinger(singerName)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
