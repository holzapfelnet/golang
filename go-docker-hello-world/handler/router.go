package handler

import (
	mux "net/http"
)

// CreateRoutes creates a new router
func CreateRoutes() *mux.ServeMux {
	r := mux.NewServeMux()
	r.HandleFunc("/greet", GetGreetingsHandler)
	return r
}
