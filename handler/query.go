package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/high-quality-sausages/msc-backend/store"
)

var songType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"num": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"singer": &graphql.Field{
				Type: singerType,
			},
			"length": &graphql.Field{
				Type: graphql.Int,
			},
			"releaseDate": &graphql.Field{
				Type: graphql.String,
			},
			"path": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var singerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Singer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"nation": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"songs": &graphql.Field{
				Type: graphql.NewList(songType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Song ID",
						Type:        graphql.ID,
					},
					"name": &graphql.ArgumentConfig{
						Description: "Song Name",
						Type:        graphql.String,
					},
					"singer_name": &graphql.ArgumentConfig{
						Description: "Singer Name",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)

					if ok {
						idInt, err := strconv.Atoi(id)
						if err != nil {
							fmt.Println("Failed to parse string to int")
						}
						song, err := store.GetSongByID(idInt)
						if err != nil {
							return nil, err
						}
						return song, nil
					}

					name, ok := p.Args["name"].(string)
					if ok {
						songs, err := store.GetSongsByName(name)
						if err != nil {
							return nil, err
						}
						return songs, nil
					}

					singerName, ok := p.Args["singer_name"].(string)
					if ok {
						songs, err := store.GetSongsBySingerName(singerName)
						if err != nil {
							return nil, err
						}
						return songs, nil
					}

					err := errors.New("Missing required parameter 'id'")
					return nil, err
				},
			},
		},
	},
)
