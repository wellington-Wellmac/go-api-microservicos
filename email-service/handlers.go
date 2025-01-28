package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

// Estrutura para representar um e-mail
type EmailRequest struct {
	To      string json:"to"
	Subject string json:"subject"
	Body    string json:"body"
}

// Função para lidar com o envio de e-mails
func sendEmail(w http.ResponseWriter, r *http.Request) {
	var emailReq EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sendEmailSMTP(emailReq.To, emailReq.Subject, emailReq.Body)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email sent successfully")
}

// Função para enviar e-mails via SMTP usando Gomail
func sendEmailSMTP(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "seu-email@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "seu-email@gmail.com", "sua-senha")

	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("Erro ao enviar e-mail: %v", err)
		return err
	}
	return nil
}