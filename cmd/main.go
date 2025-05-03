package main

import (
	"log"
	"main/internal/nats"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if exists, but don't fail if it doesn't
	_ = godotenv.Load()

	// Debug environment variables
	log.Printf("SMTP_HOST: %s", os.Getenv("SMTP_HOST"))
	log.Printf("SMTP_PORT: %s", os.Getenv("SMTP_PORT"))
	log.Printf("SMTP_FROM: %s", os.Getenv("SMTP_FROM"))
	log.Printf("NATS_URL: %s", os.Getenv("NATS_URL"))

	log.Println("📧 EmailService initalized")
	nats.StartListener()
}
