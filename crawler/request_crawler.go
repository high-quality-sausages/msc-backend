package crawler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/high-quality-sausages/msc-backend/conf"
	"github.com/high-quality-sausages/msc-backend/models"
	repo "github.com/high-quality-sausages/msc-backend/repository"
)

var config = conf.GetConfig()

// RequestSongsByName : 请求crawler爬取信息
func RequestSongsByName(name string) ([]*models.Song, error) {
	info := make(map[string]string)
	info["name"] = name

	bytesData, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Faield to parse bytesData")
		return nil, err
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", "http://"+config.CrawlerConf.Host+"/api/name", reader)

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

	jsonResp := repo.RequestResp{}

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
