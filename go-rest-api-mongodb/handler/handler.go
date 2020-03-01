package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	repo "../repository"
)

// AddToDoItemHandler inserts a new todo item into the mongo db
func AddToDoItemHandler(repository repo.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		s := string(body)
		toToItemData := repo.ToDoItem{}
		err = json.Unmarshal([]byte(s), &toToItemData)
		if err != nil {
			w.WriteHeader(400)
		}

		err = repository.AddItem(toToItemData)

		if err != nil {
			w.WriteHeader(400)
		}

		w.WriteHeader(200)

	}
}

// GetToDoItemsHandler inserts a new todo item into the mongo db
func GetToDoItemsHandler(repository repo.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

		var resultItems, err = repository.GetAllItems()

		if err != nil {
			w.WriteHeader(400)
		}

		b, err := json.Marshal(resultItems)
		if err != nil {
			fmt.Println("error:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	}
}
