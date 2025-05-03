package main

import (
	"log"
	"main/internal/nats"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Debug environment variables
	log.Printf("SMTP_HOST: %s", os.Getenv("SMTP_HOST"))
	log.Printf("SMTP_PORT: %s", os.Getenv("SMTP_PORT"))
	log.Printf("SMTP_FROM: %s", os.Getenv("SMTP_FROM"))
	log.Printf("NATS_URL: %s", os.Getenv("NATS_URL"))

	log.Println("📧 EmailService initalized")
	nats.StartListener()
}
