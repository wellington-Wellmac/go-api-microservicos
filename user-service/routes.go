package main

import "net/http"

func handleRequests() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/", getUser)
}
