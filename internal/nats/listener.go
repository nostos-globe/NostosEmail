package nats

import (
	"encoding/json"
	"log"
	"os"

	"main/internal/mailer"

	"github.com/nats-io/nats.go"
)

type EmailEvent struct {
	Event     string `json:"event"`
	Email     string `json:"email"`
	Name      string `json:"name,omitempty"`
	Link      string `json:"confirmation_link,omitempty"`
	ResetLink string `json:"reset_link,omitempty"`
}

func StartListener() {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal("❌ Failed to connect to NATS:", err)
	}
	defer nc.Close()

	sub, err := nc.Subscribe("user.*", func(msg *nats.Msg) {
		var evt EmailEvent
		if err := json.Unmarshal(msg.Data, &evt); err != nil {
			log.Println("❌ Invalid event:", err)
			return
		}

		switch evt.Event {
		case "user.registered":
			if err := mailer.SendConfirmationEmail(evt.Email, evt.Name, evt.Link); err != nil {
				log.Printf("❌ Failed to send confirmation email: %v", err)
			}
		case "user.password_reset_requested":
			if err := mailer.SendPasswordResetEmail(evt.Email, evt.ResetLink); err != nil {
				log.Printf("❌ Failed to send password reset email: %v", err)
			}
		}
	})
	if err != nil {
		log.Fatal("❌ Failed to subscribe:", err)
	}
	defer sub.Unsubscribe()

	log.Println("✅ Listening for NATS events on 'user.*'")
	select {}
}
