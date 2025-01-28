package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Auth Service is running on port 8082...")
	handleRequests()
	log.Fatal(http.ListenAndServe(":8082", nil))
}