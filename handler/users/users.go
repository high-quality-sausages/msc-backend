package handler

import repo "github.com/high-quality-sausages/msc-backend/repository"

type User struct {
	UserName string
	Password string
	Age int
	Gender int
	Location string
	Token string
}

type Query interface {
	QuerySongsByID(ID string) ([]models.Song, error)
}