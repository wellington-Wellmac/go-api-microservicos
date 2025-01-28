package main

import (
	"encoding/json"
	"net/http"
)

func sendEmail(w http.ResponseWriter, r *http.Request) {
	// Lógica para enviar e-mail...
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Email sent successfully"})
}