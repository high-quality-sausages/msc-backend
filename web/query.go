package web

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/high-quality-sausages/msc-backend/store"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"queryByID": &graphql.Field{
				Type: songType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Song ID",
						Type:        graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					idInt, err := strconv.Atoi(id)
					if err != nil {
						fmt.Println("Failed to parse string to int")
					}

					if ok {
						song, err := store.GetSongByID(idInt)
						if err != nil {
							fmt.Println("Failed to get song by id")
							return nil, nil
						}
						return song, nil
					}

				},
			},
		},
	},
)
