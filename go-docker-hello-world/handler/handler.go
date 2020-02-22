package handler

import (
	"fmt"
	"net/http"
)

// GetGreetingsHandler sends greetings
func GetGreetingsHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string("{ text: 'Greetings from docker container'}"))
}
