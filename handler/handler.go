package handler

import "net/http"

type GraphQLHandler struct {
	schema *graphql.Schema
}

func (h *GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
