package main

type User struct {
	ID       int    json:"id"
	Username string json:"username"
	Email    string json:"email"
}

var users = []User{
	{ID: 1, Username: "tiago", Email: "tiago@example.com"},
	{ID: 2, Username: "ana", Email: "ana@example.com"},
}

handlers.go:

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
