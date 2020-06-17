package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/high-quality-sausages/msc-backend/web"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: web.QueryType,
	})

	if err != nil {
		fmt.Println("error")
	}

	http.Handle("/graphql", handler(schema))
	http.ListenAndServe("0.0.0.0:8080", nil)
}
