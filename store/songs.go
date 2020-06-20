package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RequestResp struct {
	Data []Song `json:"data"`
	Err  string `json:"error"`
}

type RespSong struct {
	Num    string `json:"num"`
	Name   string `json:"name"`
	Singer string `json:"singer"`
}

type Song struct {
	ID          string `json:"id"`
	Num         string `json:"num"`
	Name        string `json:"name"`
	Singer      Artist `json:"singer"`
	Author      Artist `json:"author"`
	ReleaseDate string `json:"release_date"`
}

// Artist : 歌手/作曲家
type Artist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Nation string `json:"nation"`
}

// GetSongByID : 通过ID查询
func GetSongByID(id int) (*Song, error) {
	artist := Artist{
		ID:     "1",
		Name:   "Adel",
		Age:    45,
		Nation: "US",
	}
	return &Song{
		ID:          "12",
		Name:        "hello",
		Singer:      artist,
		Author:      artist,
		ReleaseDate: "2020-01-02",
	}, nil
}

func RequestSongsByName(name string) ([]Song, error) {
	info := make(map[string]string)
	info["name"] = name

	bytesData, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Faield to parse bytesData")
		return nil, err
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", "http://localhost:8000/api/name", reader)

	if err != nil {
		fmt.Println("Failed to carete request")
		return nil, err
	}

	defer request.Body.Close()

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Failed to do request")
		return nil, err
	}

	jsonResp := RequestResp{}

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	if jsonResp.Err != "" {
		fmt.Println("Error:", jsonResp.Err)
		return nil, errors.New(jsonResp.Err)
	}

	return jsonResp.Data, nil
}

// GetSongByName : 通过歌名模糊查询歌曲
func GetSongsByName(name string) ([]Song, error) {
	songList, err := RequestSongsByName(name)
	if err != nil {
		fmt.Println(err)
	}
	return songList, nil
}

// GetSongBySingerName : 通过歌手名模糊查询歌曲
func GetSongsBySingerName(singerName string) ([]*Song, error) {
	artist := Artist{
		ID:     "1",
		Name:   "Adel",
		Age:    45,
		Nation: "US",
	}

	result := make([]*Song, 0)
	song := &Song{
		ID:          "12",
		Name:        "hello",
		Singer:      artist,
		Author:      artist,
		ReleaseDate: "2020-01-02",
	}

	result = append(result, song)

	return result, nil
}

func parseCrawlerResp(rawInfo string) {}
