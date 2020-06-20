package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func Register(schema graphql.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		query, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data": "",
				"err":  err.Error(),
			})
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})

		c.JSON(http.StatusOK, gin.H{
			"data": result.Data,
			"err":  "",
		})
	}
}

func RequestSongByName(name string) {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/json; charset=utf-8", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)
}
