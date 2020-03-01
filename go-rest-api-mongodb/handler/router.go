package handler

import (
	mux "net/http"

	repository "../repository"
)

// CreateRoutes creates a new router
func CreateRoutes(repository repository.Repository) *mux.ServeMux {
	r := mux.NewServeMux()
	r.HandleFunc("/add", AddToDoItemHandler(repository))
	r.HandleFunc("/", GetToDoItemsHandler(repository))
	return r
}
