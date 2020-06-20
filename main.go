package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	handler "github.com/high-quality-sausages/msc-backend/handler"
	"github.com/high-quality-sausages/msc-backend/util"
)

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: handler.QueryType,
	})

	if err != nil {
		fmt.Println("error")
	}

	router := gin.Default()
	router.Use(util.Cors())

	router.POST("/gql", handler.Register(schema))
	router.Run("0.0.0.0:8080")
}
