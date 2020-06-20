package store

import (
	"encoding/json"
	"testing"
)

func TestGetSongByID(t *testing.T) {
	resp, err := GetSongByID(10)
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}

func TestGetSongsByName(t *testing.T) {
	resp, err := GetSongsByName("hello")
	if err != nil {
		t.Error(err)
	}

	t.Log(resp[0].Name)
}

func TestGetSongsBySingerName(t *testing.T) {
	resp, err := GetSongsBySingerName("adel")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp[0].Name)
}

func TestRequestSongsByName(t *testing.T) {
	resp, err := RequestSongsByName("hello")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp[0].Singer)
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
