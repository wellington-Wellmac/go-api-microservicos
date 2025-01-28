package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Email Service is running on port 8083...")
	handleRequests()
	log.Fatal(http.ListenAndServe(":8083", nil))
}