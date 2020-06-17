package web

import "github.com/graphql-go/graphql"

type Song struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Singer      Singer `json:"singer"`
	Author      string
	Length      int
	ReleaseDate string
	Path        string `json:"path"`
}

type Singer struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Nation string `json:"nation"`
}

var songType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
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

// fields := graphql.Fields{
// 	"song": &graphql.Field{
// 		Type: songType,
// 		Description: "Get Song By ID",
// 		Args: graphql.FieldConfigArgument{
// 			"id": &graphql.ArgumentConfig{
// 				Type: graphql.Int,
// 			},
// 		},
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 			id, ok := p.Args["id"].(int)
// 			if ok {
// 				for _, song := range songs {
// 					if int(song.ID) == id {
// 						return song, nil
// 					}
// 				}
// 			}
// 			return nil, nil
// 		},
// 	},
// }
