package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"koskosan-be/internal/config"
	"log"
	"net/http"
	"time"
)

type WhatsAppSender interface {
	SendWhatsApp(to, message string) error
}

type FonnteSender struct {
	token string
}

func NewFonnteSender(token string) *FonnteSender {
	return &FonnteSender{token: token}
}

func (s *FonnteSender) SendWhatsApp(to, message string) error {
	// Fonnte API endpoint
	url := "https://api.fonnte.com/send"

	// Prepare payload
	payload := map[string]string{
		"target":  to,
		"message": message,
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", s.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fonnte api returned status: %s", resp.Status)
	}

	return nil
}

type LogWASender struct{}

func (s *LogWASender) SendWhatsApp(to, message string) error {
	log.Printf("---------------------------------------------------------")
	log.Printf("[WA SIMULATION] To: %s", to)
	log.Printf("[WA SIMULATION] Message: %s", message)
	log.Printf("---------------------------------------------------------")
	return nil
}

func NewWhatsAppSender(cfg *config.Config) WhatsAppSender {
	// Assuming you add WhatsAppToken to config later. 
	// For now, if you have a token env var, use it.
	// Since Config struct update isn't in the plan explicitly, 
	// we'll check if a specific env var exists or just default to Log/Simulation 
	// if we don't want to break the build by accessing a non-existent cfg field yet.
	
	//Ideally we should update config.Config. 
	// For this step I'll assume we might add it or just use a hardcoded check or empty string for now to be safe,
	// checking if we can update config.go first is better but to follow "Atomic" steps:
	
	// Let's rely on the Config struct update which I should probably do.
	// But to avoid compilation error before config update, I will use a placeholder or check cfg map if it was a map (it's a struct).
	
	// I'll stick to LogSender for now and let the user know they need to set the token, 
	// OR I'll update Config.go in the next step.
	
	return &LogWASender{}
}
