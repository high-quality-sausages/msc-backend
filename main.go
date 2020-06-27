package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/high-quality-sausages/msc-backend/handler/songs"
	"github.com/high-quality-sausages/msc-backend/util"
)

func main() {

	// schema, err := graphql.NewSchema(graphql.SchemaConfig{
	// 	Query: handler.QueryType,
	// })

	// if err != nil {
	// 	fmt.Println("error")
	// }

	router := gin.Default()
	router.Use(util.Cors())

	// router.POST("/gql", handler.Register(schema))
	router.GET("/songs/query", handler.QuerySongs)
	router.Run(":8080")
}
