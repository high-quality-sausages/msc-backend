package repository

import (
	"encoding/json"
	"testing"
)

func TestGetSongByID(t *testing.T) {
	resp, err := GetSongByID("w001")
	if err != nil {
		t.Error(err)
	}

	t.Log(resp[0].Name)
}

func TestGetSongsByName(t *testing.T) {
	resp, err := GetSongsByName("hello")
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}

func TestGetSongsBySingerName(t *testing.T) {
	resp, err := GetSongsBySinger("adel")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp[0].Name)
}

func TestParse(t *testing.T) {
	info := `
	{
		"data": [
			{"num": "1", "name": "hello", "singer": "Adel"},
			{"num": "2", "name": "上学威龙", "singer": "法老"}
		],
		"err": ""
	}
	`
	resp := RequestResp{}

	temp := []byte(info)
	if err := json.Unmarshal(temp, &resp); err != nil {
		t.Error(err)
	}

	t.Log(resp.Data[0])
}

func TestGetArtist(t *testing.T) {
	artist, err := GetArtist("001")
	if err != nil {
		t.Error(err)
	}
	t.Log(artist.ID)
	// t.Log(artist.Name)
	// t.Log(artist.Nation)
}
