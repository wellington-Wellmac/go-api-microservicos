package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("User Service is running on port 8081...")
	handleRequests()
	log.Fatal(http.ListenAndServe(":8081", nil))
}